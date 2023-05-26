package plus

import (
	"go.uber.org/zap"

	"github.com/jodydadescott/shelly-manager/shelly/logging"
	"github.com/jodydadescott/shelly-manager/shelly/plus/bluetooth"
	"github.com/jodydadescott/shelly-manager/shelly/plus/cloud"
	"github.com/jodydadescott/shelly-manager/shelly/plus/ethernet"
	"github.com/jodydadescott/shelly-manager/shelly/plus/input"
	"github.com/jodydadescott/shelly-manager/shelly/plus/light"
	"github.com/jodydadescott/shelly-manager/shelly/plus/mqtt"
	"github.com/jodydadescott/shelly-manager/shelly/plus/msghandlers"
	"github.com/jodydadescott/shelly-manager/shelly/plus/shelly"
	"github.com/jodydadescott/shelly-manager/shelly/plus/switchx"
	"github.com/jodydadescott/shelly-manager/shelly/plus/system"
	"github.com/jodydadescott/shelly-manager/shelly/plus/types"
	"github.com/jodydadescott/shelly-manager/shelly/plus/webhook"
	"github.com/jodydadescott/shelly-manager/shelly/plus/websocket"
	"github.com/jodydadescott/shelly-manager/shelly/plus/wifi"
)

type Config struct {
	Hostname     string
	Password     string
	Username     string
	DebugEnabled bool
	ZapLogger    *zap.Logger
}

type Client struct {
	_system      *system.Client
	_shelly      *shelly.Client
	_wifi        *wifi.Client
	_bluetooth   *bluetooth.Client
	_mqtt        *mqtt.Client
	_cloud       *cloud.Client
	_switch      *switchx.Client
	_light       *light.Client
	_input       *input.Client
	_websocket   *websocket.Client
	_ethernet    *ethernet.Client
	_webhook     *webhook.Client
	debugEnabled bool
	types.MessageHandlerFactory
}

func New(config *Config) (*Client, error) {

	t := &Client{
		debugEnabled: config.DebugEnabled,
	}

	zapLogger := config.ZapLogger

	if zapLogger == nil {
		if t.debugEnabled {
			zapLogger = logging.GetDebugZapLogger()
		} else {
			zapLogger = logging.GetDefaultZapLogger()
		}
	}

	zap.ReplaceGlobals(zapLogger)

	if t.debugEnabled {
		zap.L().Debug("debug is enabled")
	}

	username := config.Username
	if username == "" {
		username = DefaultUsername
	}

	wsClient, err := msghandlers.NewWS(&msghandlers.Config{
		Hostname:     config.Hostname,
		Password:     config.Password,
		Username:     username,
		DebugEnabled: config.DebugEnabled,
	})

	if err != nil {
		return nil, err
	}

	t.MessageHandlerFactory = wsClient
	return t, nil
}

func (t *Client) System() *system.Client {
	if t._system == nil {
		t._system = system.New(t)
	}
	return t._system
}

func (t *Client) Shelly() *shelly.Client {
	if t._shelly == nil {
		t._shelly = shelly.New(t)
	}
	return t._shelly
}

func (t *Client) Bluetooth() *bluetooth.Client {
	if t._bluetooth == nil {
		t._bluetooth = bluetooth.New(t)
	}
	return t._bluetooth
}

func (t *Client) Mqtt() *mqtt.Client {
	if t._mqtt == nil {
		t._mqtt = mqtt.New(t)
	}
	return t._mqtt
}

func (t *Client) Ethernet() *ethernet.Client {
	if t._ethernet == nil {
		t._ethernet = ethernet.New(t)
	}
	return t._ethernet
}

func (t *Client) WiFi() *wifi.Client {
	if t._wifi == nil {
		t._wifi = wifi.New(t)
	}
	return t._wifi
}

func (t *Client) Cloud() *cloud.Client {
	if t._cloud == nil {
		t._cloud = cloud.New(t)
	}
	return t._cloud
}

func (t *Client) Switch() *switchx.Client {
	if t._switch == nil {
		t._switch = switchx.New(t)
	}
	return t._switch
}

func (t *Client) Light() *light.Client {
	if t._light == nil {
		t._light = light.New(t)
	}
	return t._light
}

func (t *Client) Input() *input.Client {
	if t._input == nil {
		t._input = input.New(t)
	}
	return t._input
}

func (t *Client) Websocket() *websocket.Client {
	if t._websocket == nil {
		t._websocket = websocket.New(t)
	}
	return t._websocket
}

func (t *Client) WebHook() *webhook.Client {
	if t._webhook == nil {
		t._webhook = webhook.New(t)
	}
	return t._webhook
}

func (t *Client) Close() {

	zap.L().Debug("(*Client) Close()")

	if t._system != nil {
		t._system.Close()
	}

	if t._system != nil {
		t._system.Close()
	}

	if t._shelly != nil {
		t._shelly.Close()
	}

	if t._wifi != nil {
		t._wifi.Close()
	}

	if t._bluetooth != nil {
		t._bluetooth.Close()
	}

	if t._mqtt != nil {
		t._mqtt.Close()
	}

	if t._cloud != nil {
		t._cloud.Close()
	}

	if t._switch != nil {
		t._switch.Close()
	}

	if t._light != nil {
		t._light.Close()
	}

	if t._input != nil {
		t._input.Close()
	}

	if t._websocket != nil {
		t._websocket.Close()
	}

	if t._ethernet != nil {
		t._ethernet.Close()
	}

	if t._webhook != nil {
		t._webhook.Close()
	}

	t.MessageHandlerFactory.Close()
}
