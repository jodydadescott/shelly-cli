package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"sync"
	"time"

	gorilla "github.com/gorilla/websocket"
	"go.uber.org/zap"

	"github.com/jodydadescott/shelly-manager/shelly/plus/types"
)

const (
	WsScheme         = "ws"
	WsPath           = "/rpc"
	FailWaitDuration = time.Duration(3) * time.Second
)

var defaultSendTimeout = time.Duration(time.Second * 15)

type MessageHandlerFactory = types.MessageHandlerFactory
type MessageHandler = types.MessageHandler
type Request = types.Request
type AuthResponse = types.AuthResponse
type AuthRequest = types.AuthRequest
type Response = types.Response

type Config interface {
	GetHostname() string
	GetPassword() string
	GetUsername() string
	GetSendTimeout() time.Duration
	IsDebugEnabled() bool
}

type Client struct {
	hostname       string
	username       string
	password       string
	mutex          sync.RWMutex
	handleMap      map[int]*Handle
	egressMessages chan []byte
	uniqID         int
	wg             sync.WaitGroup
	cancel         context.CancelFunc
	sendTimeout    time.Duration
	debugEnabled   bool
}

func New(config Config) (MessageHandlerFactory, error) {
	zap.L().Debug("New")

	t := &Client{
		hostname:       config.GetHostname(),
		password:       config.GetPassword(),
		username:       config.GetUsername(),
		sendTimeout:    config.GetSendTimeout(),
		handleMap:      make(map[int]*Handle),
		egressMessages: make(chan []byte, 50),
		debugEnabled:   config.IsDebugEnabled(),
	}

	if t.hostname == "" {
		return nil, fmt.Errorf("hostname is required")
	}

	if t.sendTimeout <= 0 {
		t.sendTimeout = defaultSendTimeout
		zap.L().Debug("sendTimeout set to default")
	}

	if t.password == "" {
		zap.L().Debug("password is not set")
	}

	t.run()
	return t, nil
}

func (t *Client) Close() {
	zap.L().Debug("(*Client) Close()")
	t.cancel()
	t.wg.Wait()
	close(t.egressMessages)
}

func (t *Client) run() {

	if t.sendTimeout <= 0 {
		t.sendTimeout = defaultSendTimeout
		zap.L().Debug(fmt.Sprintf("sendTimeout set to default %d", t.sendTimeout))
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.cancel = cancel

	routeMessage := func(b []byte) {
		msg := &Response{}
		err := json.Unmarshal(b, msg)
		if err != nil {
			zap.L().Error(fmt.Sprintf("routeMessage error %v", err))
			return
		}

		t.mutex.RLock()
		defer t.mutex.RUnlock()

		handle := t.handleMap[msg.ID]
		if handle == nil {
			zap.L().Error(fmt.Sprintf("handle lookup ID %d failure", msg.ID))
			return
		}

		handle.receive <- &responseWrapper{
			response: msg,
			rawBytes: b,
		}
	}

	handleEgress := func(conn *gorilla.Conn, errs chan error) {

		go func() {
			for {
				select {
				case <-ctx.Done():
					conn.WriteMessage(gorilla.CloseMessage, gorilla.FormatCloseMessage(gorilla.CloseNormalClosure, ""))
					return

				case b := <-t.egressMessages:

					if t.debugEnabled {
						zap.L().Debug(fmt.Sprintf("TX->%s", string(b)))
					}

					err := conn.WriteMessage(gorilla.BinaryMessage, b)
					if err != nil {
						errs <- err
						return
					}
				}
			}

		}()
	}

	handleIngress := func(conn *gorilla.Conn, errs chan error) {
		go func() {
			for {
				_, b, err := conn.ReadMessage()

				if t.debugEnabled {
					zap.L().Debug(fmt.Sprintf("RX->%s", string(b)))
				}

				if err != nil {
					errs <- err
					return
				}

				routeMessage(b)
			}
		}()
	}

	handle := func(conn *gorilla.Conn) error {
		errs := make(chan error, 2)
		defer close(errs)
		handleIngress(conn, errs)
		handleEgress(conn, errs)

		select {
		case <-ctx.Done():
			return nil
		case err := <-errs:
			return err
		}

	}

	connect := func() error {
		theURL := url.URL{Scheme: WsScheme, Host: t.hostname, Path: WsPath}
		conn, _, err := gorilla.DefaultDialer.Dial(theURL.String(), nil)

		if err != nil {
			return err
		}

		return handle(conn)
	}

	go func() {
		t.wg.Add(1)
		defer t.wg.Done()

		for {

			zap.L().Debug(fmt.Sprintf("Connecting to %s", t.hostname))

			err := connect()

			if err == nil {
				zap.L().Debug("Connected")
				return
			}

			zap.L().Debug(fmt.Sprintf("Connect error %v; will try again in %v", err, FailWaitDuration))

			select {

			case <-ctx.Done():
				zap.L().Debug("Connect cancelled")
				return

			case <-time.After(FailWaitDuration):
				continue
			}

		}

	}()

}

func (t *Client) NewHandle() MessageHandler {

	zap.L().Debug("(*Client) NewHandle()")

	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.uniqID = t.uniqID + 1

	handle := &Handle{
		Client:  t,
		id:      t.uniqID,
		receive: make(chan *responseWrapper),
		done:    make(chan struct{}),
	}

	t.handleMap[handle.id] = handle

	return handle
}

func (t *Handle) Close() {

	zap.L().Debug("(*Handle) Close()")

	close(t.done)
	close(t.receive)
	t.mutex.Lock()
	defer t.mutex.Unlock()
	delete(t.handleMap, t.id)
}

type responseWrapper struct {
	response *Response
	rawBytes []byte
}

type Handle struct {
	*Client
	id      int
	receive chan *responseWrapper
	done    chan struct{}
}

func (t *Handle) Send(ctx context.Context, request *Request) ([]byte, error) {

	zap.L().Debug("(*Handle) Send(ctx, *Request)")

	request = request.Clone()
	request.ID = t.id

	requestBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	t.egressMessages <- requestBytes

	waitOnResponse := func() (*responseWrapper, error) {

		select {

		case response := <-t.receive:
			return response, nil

		case <-t.done:
			return nil, fmt.Errorf("channel closed by server")

		case <-ctx.Done():
			return nil, fmt.Errorf("channel closed by client")

		case <-time.After(t.sendTimeout):
			return nil, fmt.Errorf("timeout waiting for response")

		}

	}

	response, err := waitOnResponse()

	if err != nil {
		return nil, err
	}

	if response.response.Error != nil {

		if response.response.Error.Code == 401 {

			zap.L().Debug("server responded with auth required")

			if t.username == "" {
				return nil, fmt.Errorf("username is required")
			}

			if t.password == "" {
				return nil, fmt.Errorf("password is required")
			}

			authRequest := &AuthRequest{}
			err = json.Unmarshal([]byte(response.response.Error.Message), authRequest)
			if err != nil {
				return nil, err
			}

			authRequest.Username = t.username
			authRequest.Password = t.password
			authResponse, err := authRequest.ToAuthResponse()

			if err != nil {
				return nil, err
			}

			request.Auth = authResponse

			requestBytes, err := json.Marshal(request)
			if err != nil {
				return nil, err
			}

			t.egressMessages <- requestBytes

			response, err := waitOnResponse()

			if err != nil {
				return nil, err
			}

			return response.rawBytes, nil

		}

		return nil, response.response.Error
	}

	return response.rawBytes, nil
}
