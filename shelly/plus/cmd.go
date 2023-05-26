package plus

import (
	"context"

	"github.com/jodydadescott/shelly-manager/shelly/plus/bluetooth"
	"github.com/jodydadescott/shelly-manager/shelly/plus/cloud"
	"github.com/jodydadescott/shelly-manager/shelly/plus/ethernet"
	"github.com/jodydadescott/shelly-manager/shelly/plus/input"
	"github.com/jodydadescott/shelly-manager/shelly/plus/light"
	"github.com/jodydadescott/shelly-manager/shelly/plus/mqtt"
	"github.com/jodydadescott/shelly-manager/shelly/plus/shelly"
	"github.com/jodydadescott/shelly-manager/shelly/plus/switchx"
	"github.com/jodydadescott/shelly-manager/shelly/plus/system"
	"github.com/jodydadescott/shelly-manager/shelly/plus/websocket"
	"github.com/jodydadescott/shelly-manager/shelly/plus/wifi"
	"github.com/spf13/cobra"
)

type callback interface {
	GetHostname() string
	GetUsername() string
	GetPassword() string
	WriteObject(any) error
	WriteStderr(string)
	ReadInput() ([]byte, error)
	IsDebugEnabled() bool
}

type Cmd struct {
	_client *Client
	*cobra.Command
	callback
}

func NewCmd(callback callback) *cobra.Command {

	d := &Cmd{
		Command: &cobra.Command{
			Use:   "plus",
			Short: "Shelly Plus",
		},
		callback: callback,
	}

	d.AddCommand(system.NewCmd(d), shelly.NewCmd(d), wifi.NewCmd(d), bluetooth.NewCmd(d), mqtt.NewCmd(d))
	d.AddCommand(cloud.NewCmd(d), switchx.NewCmd(d), input.NewCmd(d), websocket.NewCmd(d))
	d.AddCommand(ethernet.NewCmd(d), light.NewCmd(d))
	return d.Command
}

func (t *Cmd) client() (*Client, error) {

	if t._client != nil {
		return t._client, nil
	}

	client, err := New(&Config{
		Hostname:     t.GetHostname(),
		Username:     t.GetUsername(),
		Password:     t.GetPassword(),
		DebugEnabled: t.IsDebugEnabled(),
	})
	if err != nil {
		return nil, err
	}
	t._client = client

	return t._client, nil
}

func (t *Cmd) LogDebug(s string) {
	t.WriteStderr(s)
}

func (t *Cmd) System() (*system.Client, error) {
	client, err := t.client()
	if err != nil {
		return nil, err
	}
	return client.System(), nil
}

func (t *Cmd) Shelly() (*shelly.Client, error) {
	client, err := t.client()
	if err != nil {
		return nil, err
	}
	return client.Shelly(), nil
}

func (t *Cmd) Bluetooth() (*bluetooth.Client, error) {
	client, err := t.client()
	if err != nil {
		return nil, err
	}
	return client.Bluetooth(), nil
}

func (t *Cmd) Mqtt() (*mqtt.Client, error) {
	client, err := t.client()
	if err != nil {
		return nil, err
	}
	return client.Mqtt(), nil
}

func (t *Cmd) Ethernet() (*ethernet.Client, error) {
	client, err := t.client()
	if err != nil {
		return nil, err
	}
	return client.Ethernet(), nil
}

func (t *Cmd) WiFi() (*wifi.Client, error) {
	client, err := t.client()
	if err != nil {
		return nil, err
	}
	return client.WiFi(), nil
}

func (t *Cmd) Cloud() (*cloud.Client, error) {
	client, err := t.client()
	if err != nil {
		return nil, err
	}
	return client.Cloud(), nil
}

func (t *Cmd) Switch() (*switchx.Client, error) {
	client, err := t.client()
	if err != nil {
		return nil, err
	}
	return client.Switch(), nil
}

func (t *Cmd) Light() (*light.Client, error) {
	client, err := t.client()
	if err != nil {
		return nil, err
	}
	return client.Light(), nil
}

func (t *Cmd) Input() (*input.Client, error) {
	client, err := t.client()
	if err != nil {
		return nil, err
	}
	return client.Input(), nil
}

func (t *Cmd) Websocket() (*websocket.Client, error) {
	client, err := t.client()
	if err != nil {
		return nil, err
	}
	return client.Websocket(), nil
}

func (t *Cmd) RebootDevice(ctx context.Context) error {
	shelly, err := t.Shelly()
	if err != nil {
		return err
	}
	return shelly.Reboot(ctx)
}
