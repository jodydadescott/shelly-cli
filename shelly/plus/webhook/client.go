package webhook

import (
	"context"
	"encoding/json"
	"fmt"
)

// GenericResponse internal use only
type GenericResponse struct {
	Response
	Result *Webhooks `json:"result,omitempty"`
}

func New(messageHandlerFactory MessageHandlerFactory) *Client {
	return &Client{
		MessageHandlerFactory: messageHandlerFactory,
	}
}

type Client struct {
	MessageHandlerFactory
	_messageHandler MessageHandler
}

func (t *Client) getMessageHandler() MessageHandler {
	if t._messageHandler != nil {
		return t._messageHandler
	}

	t._messageHandler = t.NewHandle()
	return t._messageHandler
}

// ListSupported lists all supported events that can be used to trigger a Webhook
func (t *Client) ListSupported(ctx context.Context) (*Webhooks, error) {

	method := Component + ".ListSupported"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
	})

	if err != nil {
		return nil, err
	}

	response := &GenericResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	if response.Result == nil {
		return nil, fmt.Errorf("Result is missing from response")
	}

	return response.Result, nil
}

// List lists all existing Webhooks for this device.
func (t *Client) List(ctx context.Context) (*Webhooks, error) {

	method := Component + ".List"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
	})

	if err != nil {
		return nil, err
	}

	response := &GenericResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	if response.Result == nil {
		return nil, fmt.Errorf("Result is missing from response")
	}

	return response.Result, nil
}

// Create creates a Webhook instance
func (t *Client) Create(ctx context.Context, params *Params) (*Webhooks, error) {

	method := Component + ".Create"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
		Params: params,
	})

	if err != nil {
		return nil, err
	}

	response := &GenericResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	if response.Result == nil {
		return nil, fmt.Errorf("Result is missing from response")
	}

	return response.Result, nil
}

// Update updates an existing Webhook instance
func (t *Client) Update(ctx context.Context, params *Params) (*Webhooks, error) {

	method := Component + ".Update"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
		Params: params,
	})

	if err != nil {
		return nil, err
	}

	response := &GenericResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	if response.Result == nil {
		return nil, fmt.Errorf("Result is missing from response")
	}

	return response.Result, nil
}

// Delete deletes an existing Webhook instance
func (t *Client) Delete(ctx context.Context, params *Webhooks) (*Webhooks, error) {

	method := Component + ".Update"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
		Params: params,
	})

	if err != nil {
		return nil, err
	}

	response := &GenericResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	if response.Result == nil {
		return nil, fmt.Errorf("Result is missing from response")
	}

	return response.Result, nil
}

// Delete deletes all existing Webhooks
func (t *Client) DeleteAll(ctx context.Context) (*Webhooks, error) {

	method := Component + ".DeleteAll"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
	})

	if err != nil {
		return nil, err
	}

	response := &GenericResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	if response.Result == nil {
		return nil, fmt.Errorf("Result is missing from response")
	}

	return response.Result, nil
}

func (t *Client) Close() {
	if t._messageHandler != nil {
		t._messageHandler.Close()
	}
}
