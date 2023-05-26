package shelly

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/jodydadescott/shelly-manager/shelly/plus/bluetooth"
	"github.com/jodydadescott/shelly-manager/shelly/plus/cloud"
	"github.com/jodydadescott/shelly-manager/shelly/plus/ethernet"
	"github.com/jodydadescott/shelly-manager/shelly/plus/input"
	"github.com/jodydadescott/shelly-manager/shelly/plus/light"
	"github.com/jodydadescott/shelly-manager/shelly/plus/mqtt"
	"github.com/jodydadescott/shelly-manager/shelly/plus/switchx"
	"github.com/jodydadescott/shelly-manager/shelly/plus/system"
	"github.com/jodydadescott/shelly-manager/shelly/plus/types"
	"github.com/jodydadescott/shelly-manager/shelly/plus/websocket"
	"github.com/jodydadescott/shelly-manager/shelly/plus/wifi"
	"github.com/jodydadescott/shelly-manager/shelly/util"

	"github.com/hashicorp/go-multierror"
)

// Result internal use only
type Result struct {
	RestartRequired bool   `json:"restart_required,omitempty"`
	Error           *Error `json:"error,omitempty"`
}

// SetConfigResponse internal use only
type SetConfigResponse struct {
	Response
	Result *Result `json:"result,omitempty"`
}

// GetConfigResponse internal use only
type GetConfigResponse struct {
	Response
	Result *ShellyConfig `json:"result,omitempty"`
}

// GetStatusResponse internal use only
type GetStatusResponse struct {
	Response
	Result *ShellyStatus `json:"result,omitempty"`
}

// DeviceInfoResponse internal use only
type DeviceInfoResponse struct {
	Response
	Result *DeviceInfo `json:"result,omitempty"`
}

// CheckForUpdateResponse Shelly component object
type CheckForUpdateResponse struct {
	Response
	Result *SystemAvailableUpdates `json:"result,omitempty"`
}

// ListMethodsResponse internal use only
type ListMethodsResponse struct {
	Response
	Result *ShellyRPCMethods `json:"result,omitempty"`
}

type clientContract interface {
	System() *system.Client
	Bluetooth() *bluetooth.Client
	Mqtt() *mqtt.Client
	WiFi() *wifi.Client
	Cloud() *cloud.Client
	Switch() *switchx.Client
	Input() *input.Client
	Light() *light.Client
	Websocket() *websocket.Client
	Ethernet() *ethernet.Client
	NewHandle() MessageHandler
}

func New(clientContract clientContract) *Client {
	return &Client{
		clientContract: clientContract,
	}
}

type Client struct {
	clientContract
	_messageHandler MessageHandler
}

func (t *Client) getMessageHandler() MessageHandler {
	if t._messageHandler != nil {
		return t._messageHandler
	}

	t._messageHandler = t.NewHandle()
	return t._messageHandler
}

// GetStatus returns the status of all the components of the device.
func (t *Client) GetStatus(ctx context.Context) (*ShellyStatus, error) {

	method := Component + ".GetStatus"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
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

// ListMethods lists all available RPC methods. It takes into account both ACL and authentication restrictions
// and only lists the methods allowed for the particular user/channel that's making the request.
func (t *Client) ListMethods(ctx context.Context) (*ShellyRPCMethods, error) {

	// Do NOT validate command here because it would be recursive

	method := Component + ".ListMethods"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
	})
	if err != nil {
		return nil, err
	}

	response := &ListMethodsResponse{}
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

// GetConfig returns the configuration of all the components of the device.
func (t *Client) GetConfig(ctx context.Context) (*ShellyConfig, error) {

	method := Component + ".GetConfig"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
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

// SetConfig sets the configuration for each component with non nil config. Note that this function
// calls into each componenet as necessary.
func (t *Client) SetConfig(ctx context.Context, config *ShellyConfig) (*ShellyReport, error) {

	mresp := &ShellyReport{}

	var errors *multierror.Error

	if config.Bluetooth != nil {
		resp, err := t.Bluetooth().SetConfig(ctx, config.Bluetooth)
		mresp.Bluetooth = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Bluetooth :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Cloud != nil {
		resp, err := t.Cloud().SetConfig(ctx, config.Cloud)
		mresp.Cloud = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Cloud :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Mqtt != nil {
		resp, err := t.Mqtt().SetConfig(ctx, config.Mqtt)
		mresp.Mqtt = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Mqtt :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Light0 != nil {
		resp, err := t.Light().SetConfig(ctx, 0, config.Light0)
		mresp.Light0 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Light0 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Light1 != nil {
		resp, err := t.Light().SetConfig(ctx, 1, config.Light1)
		mresp.Light1 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Light1 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Light2 != nil {
		resp, err := t.Light().SetConfig(ctx, 2, config.Light2)
		mresp.Light2 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Light2 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Light3 != nil {
		resp, err := t.Light().SetConfig(ctx, 3, config.Light3)
		mresp.Light3 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Light3 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Light4 != nil {
		resp, err := t.Light().SetConfig(ctx, 4, config.Light4)
		mresp.Light4 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Light4 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Light5 != nil {
		resp, err := t.Light().SetConfig(ctx, 5, config.Light5)
		mresp.Light5 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Light5 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Light6 != nil {
		resp, err := t.Light().SetConfig(ctx, 6, config.Light6)
		mresp.Light6 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Light6 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Light7 != nil {
		resp, err := t.Light().SetConfig(ctx, 7, config.Light7)
		mresp.Light7 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Light7 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Input0 != nil {
		resp, err := t.Input().SetConfig(ctx, 0, config.Input0)
		mresp.Input0 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Input0 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Input1 != nil {
		resp, err := t.Input().SetConfig(ctx, 1, config.Input1)
		mresp.Input1 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Input1 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Input2 != nil {
		resp, err := t.Input().SetConfig(ctx, 2, config.Input2)
		mresp.Input2 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Input2 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Input3 != nil {
		resp, err := t.Input().SetConfig(ctx, 3, config.Input3)
		mresp.Input3 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Input3 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Input4 != nil {
		resp, err := t.Input().SetConfig(ctx, 4, config.Input4)
		mresp.Input4 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Input4 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Input5 != nil {
		resp, err := t.Input().SetConfig(ctx, 5, config.Input5)
		mresp.Input5 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Input5 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Input6 != nil {
		resp, err := t.Input().SetConfig(ctx, 6, config.Input6)
		mresp.Input6 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Input6 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Input7 != nil {
		resp, err := t.Input().SetConfig(ctx, 7, config.Input7)
		mresp.Input7 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Input7 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Switch0 != nil {
		resp, err := t.Switch().SetConfig(ctx, 0, config.Switch0)
		mresp.Switch0 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Switch0 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Switch1 != nil {
		resp, err := t.Switch().SetConfig(ctx, 1, config.Switch1)
		mresp.Switch1 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Switch1 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Switch2 != nil {
		resp, err := t.Switch().SetConfig(ctx, 2, config.Switch2)
		mresp.Switch2 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Switch2 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Switch3 != nil {
		resp, err := t.Switch().SetConfig(ctx, 3, config.Switch3)
		mresp.Switch3 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Switch3 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Switch4 != nil {
		resp, err := t.Switch().SetConfig(ctx, 4, config.Switch4)
		mresp.Switch4 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Switch4 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Switch5 != nil {
		resp, err := t.Switch().SetConfig(ctx, 5, config.Switch5)
		mresp.Switch5 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Switch5 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Switch6 != nil {
		resp, err := t.Switch().SetConfig(ctx, 6, config.Switch6)
		mresp.Switch6 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Switch6 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Switch7 != nil {
		resp, err := t.Switch().SetConfig(ctx, 7, config.Switch7)
		mresp.Switch7 = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Switch7 :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.System != nil {
		resp, err := t.System().SetConfig(ctx, config.System)
		mresp.System = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("System :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Websocket != nil {
		resp, err := t.Websocket().SetConfig(ctx, config.Websocket)
		mresp.Websocket = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Websocket :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Ethernet != nil {
		resp, err := t.Ethernet().SetConfig(ctx, config.Ethernet)
		mresp.Ethernet = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Ethernet :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	// We set WiFi last because we may lose connectivity after the change

	if config.Wifi != nil {
		resp, err := t.WiFi().SetConfig(ctx, config.Wifi)
		mresp.Wifi = resp

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("WiFi :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	if config.Auth != nil {
		resp, err := t.SetAuth(ctx, &ShellyParams{
			Ha1:  config.Auth.Pass,
			User: &config.Auth.User,
		})

		if err != nil {
			errors = multierror.Append(errors, fmt.Errorf("Auth :: %v", err))
		} else {
			if resp.RestartRequired {
				mresp.RestartRequired = true
			}
		}
	}

	return mresp, errors.ErrorOrNil()
}

// GetDeviceInfo returns information about the device.
func (t *Client) GetDeviceInfo(ctx context.Context) (*DeviceInfo, error) {

	method := Component + ".GetDeviceInfo"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
	})
	if err != nil {
		return nil, err
	}

	response := &DeviceInfoResponse{}
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

// CheckForUpdate checks for new firmware version for the device and returns information about it.
// If no update is available returns empty JSON object as result.
func (t *Client) CheckForUpdate(ctx context.Context) (*UpdatesReport, error) {

	method := Component + ".CheckForUpdate"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
	})
	if err != nil {
		return nil, err
	}

	response := &CheckForUpdateResponse{}
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

	return &UpdatesReport{
		Src:              response.Src,
		AvailableUpdates: response.Result,
	}, nil
}

// Update updates the firmware version of the device.
func (t *Client) Update(ctx context.Context, params *ShellyParams) (*SetReport, error) {

	method := Component + ".Update"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
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
		Src:             response.Src,
		RestartRequired: response.Result.RestartRequired,
	}, nil
}

// FactoryReset resets the configuration to its default state
func (t *Client) FactoryReset(ctx context.Context) (*SetReport, error) {

	method := Component + ".FactoryReset"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
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
		Src:             response.Src,
		RestartRequired: response.Result.RestartRequired,
	}, nil
}

// ResetWiFiConfig resets the WiFi configuration of the device
func (t *Client) ResetWiFiConfig(ctx context.Context) (*SetReport, error) {

	method := Component + ".ResetWiFiConfig"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
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
		Src:             response.Src,
		RestartRequired: response.Result.RestartRequired,
	}, nil
}

// Reboot reboots the device
func (t *Client) Reboot(ctx context.Context) error {

	method := Component + ".Reboot"

	respBytes, err := t.getMessageHandler().Send(ctx, &Request{
		Method: method,
	})
	if err != nil {
		return err
	}

	response := &SetConfigResponse{}
	err = json.Unmarshal(respBytes, response)
	if err != nil {
		return err
	}

	if response.Error != nil {
		return response.Error
	}

	return nil
}

// HashAuth checks to see if the Ha1 attribute is hashed. If so no action is taken and a nil error is returned.
// Otherwise the hash is created and the Ha1 attribute is updated. This function is exposed so that the caller
// may create configuration.
func HashAuth(params *ShellyParams) error {

	if params.Ha1 == nil || *params.Ha1 == "" {
		return nil
	}

	if !util.IsHexadecimal(*params.Ha1) {
		if len(*params.Ha1) > 20 {
			return nil
		}
	}

	if params.User == nil || *params.User == "" {
		return fmt.Errorf("User is required")
	}

	if params.Realm == nil || *params.Realm == "" {
		return fmt.Errorf("Realm is required")
	}

	hashInput := *params.User + ":" + *params.Realm + ":" + *params.Ha1
	h := sha256.New()
	h.Write([]byte(hashInput))
	b := h.Sum(nil)

	ha1 := hex.EncodeToString(b)
	params.Ha1 = &ha1
	return nil
}

// SetAuth sets authentication details (password) for the device
func (t *Client) SetAuth(ctx context.Context, params *ShellyParams) (*SetReport, error) {

	method := Component + ".SetAuth"

	params = params.Clone()

	if params.Realm == nil {
		deviceInfo, err := t.GetDeviceInfo(ctx)
		if err != nil {
			return nil, err
		}
		params.Realm = &deviceInfo.ID
	}

	if params.User == nil || *params.User == "" {
		user := types.ShellyUser
		params.User = &user
	}

	err := HashAuth(params)
	if err != nil {
		return nil, err
	}

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
		Src:             response.Src,
		RestartRequired: response.Result.RestartRequired,
	}, nil
}

func (t *Client) PutTLSClientCert(ctx context.Context, params *ShellyParams) (*SetReport, error) {

	method := Component + ".PutTLSClientCert"

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
		Src:             response.Src,
		RestartRequired: response.Result.RestartRequired,
	}, nil
}

func (t *Client) PutTLSClientKey(ctx context.Context, params *ShellyParams) (*SetReport, error) {

	method := Component + ".PutTLSClientKey"

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
		Src:             response.Src,
		RestartRequired: response.Result.RestartRequired,
	}, nil
}

func (t *Client) PutUserCA(ctx context.Context, params *ShellyParams) (*SetReport, error) {

	method := Component + ".PutUserCA"

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
		Src:             response.Src,
		RestartRequired: response.Result.RestartRequired,
	}, nil
}

func (t *Client) Close() {
	if t._messageHandler != nil {
		t._messageHandler.Close()
	}
}
