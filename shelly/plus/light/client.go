package light

import (
	"context"
	"encoding/json"
	"fmt"
)

// Result internal use only
type Result struct {
	RestartRequired bool   `json:"restart_required,omitempty"`
	Error           *Error `json:"error,omitempty"`
}

// GetConfigResponse internal use only
type GetConfigResponse struct {
	Response
	Result *Config `json:"result,omitempty"`
	Params *Params `json:"params,omitempty"`
}

// SetConfigResponse internal use only
type SetConfigResponse struct {
	Response
	Result *Result `json:"result,omitempty"`
}

// GetStatusResponse internal use only
type GetStatusResponse struct {
	Response
	Result *Status `json:"result,omitempty"`
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

func (t *Client) GetStatus(ctx context.Context, switchId int) (*Status, error) {

	method := Component + ".GetStatus"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
		Params: &Params{
			ID: switchId,
		},
	})

	if err != nil {
		return nil, err
	}

	response := &GetStatusResponse{}
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

func (t *Client) GetConfig(ctx context.Context, switchId int) (*Config, error) {

	method := Component + ".GetConfig"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
		Params: &Params{
			ID: switchId,
		},
	})
	if err != nil {
		return nil, err
	}

	response := &GetConfigResponse{}
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

func (t *Client) SetConfig(ctx context.Context, switchId int, config *Config) (*SetReport, error) {

	method := Component + ".SetConfig"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
		Params: &Params{
			ID:     switchId,
			Config: config,
		},
	})

	if err != nil {
		return nil, err
	}

	response := &SetConfigResponse{}
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

	return &SetReport{
		Src:             response.Src,
		RestartRequired: response.Result.RestartRequired,
	}, nil
}

func (t *Client) Set(ctx context.Context, params *Params) (*SetReport, error) {

	method := Component + ".Set"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
		Params: params,
	})

	if err != nil {
		return nil, err
	}

	response := &SetConfigResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return &SetReport{
		Src: response.Src,
	}, nil
}

func (t *Client) Toggle(ctx context.Context, switchId int) (*SetReport, error) {

	method := Component + ".Toggle"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
		Params: &Params{
			ID: switchId,
		},
	})

	if err != nil {
		return nil, err
	}

	response := &SetConfigResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return &SetReport{
		Src: response.Src,
	}, nil
}

func (t *Client) Close() {
	if t._messageHandler != nil {
		t._messageHandler.Close()
	}
}
