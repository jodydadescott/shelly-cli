package light

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"gopkg.in/yaml.v2"

	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cobra"
)

var (
	truex  = true
	falsex = false
)

type callback interface {
	WriteObject(any) error
	WriteStderr(s string)
	ReadInput() ([]byte, error)
	Light() (*Client, error)
	RebootDevice(ctx context.Context) error
}

func NewCmd(callback callback) *cobra.Command {

	var autorebootArg bool
	var switchIDArg string

	getSwitchID := func() (*int, error) {

		if switchIDArg == "" {
			return nil, fmt.Errorf("switchID is required")
		}

		switchID, err := strconv.Atoi(switchIDArg)
		if err == nil {
			return &switchID, nil
		}

		return nil, fmt.Errorf("switchID must be an integer")

	}

	rootCmd := &cobra.Command{
		Use:   "light",
		Short: "Light Component",
	}

	rootCmd.PersistentFlags().StringVar(&switchIDArg, "id", "", "switch ID integer")

	getConfigCmd := &cobra.Command{
		Use:   "get-config",
		Short: "Returns config",
		RunE: func(cmd *cobra.Command, args []string) error {

			switchID, err := getSwitchID()
			if err != nil {
				return err
			}

			client, err := callback.Light()
			if err != nil {
				return err
			}

			result, err := client.GetConfig(cmd.Context(), *switchID)
			if err != nil {
				return err
			}

			return callback.WriteObject(result)
		},
	}

	getStatusCmd := &cobra.Command{
		Use:   "get-status",
		Short: "Returns status",
		RunE: func(cmd *cobra.Command, args []string) error {

			switchID, err := getSwitchID()
			if err != nil {
				return err
			}

			client, err := callback.Light()
			if err != nil {
				return err
			}

			result, err := client.GetStatus(cmd.Context(), *switchID)
			if err != nil {
				return err
			}

			return callback.WriteObject(result)
		},
	}

	getExampleConfigCmd := &cobra.Command{
		Use:   "get-example",
		Short: "generates example Config",
		RunE: func(cmd *cobra.Command, args []string) error {
			return callback.WriteObject(ExampleConfig())
		},
	}

	setConfigCmd := &cobra.Command{
		Use:   "set-config",
		Short: "sets config",
		RunE: func(cmd *cobra.Command, args []string) error {

			switchID, err := getSwitchID()
			if err != nil {
				return err
			}

			client, err := callback.Light()
			if err != nil {
				return err
			}

			b, err := callback.ReadInput()
			if err != nil {
				return err
			}

			var config *Config

			var errors *multierror.Error

			err = json.Unmarshal(b, &config)
			if err != nil {
				errors = multierror.Append(errors, err)
				err = yaml.Unmarshal(b, &config)

				if err != nil {
					errors = multierror.Append(errors, err)
					errors = multierror.Append(errors, fmt.Errorf("Invalid format. Expect JSON or YAML"))
					return errors.ErrorOrNil()
				}
			}

			report, err := client.SetConfig(cmd.Context(), *switchID, config)
			if err != nil {
				return err
			}

			if autorebootArg {
				if report.RestartRequired {
					callback.WriteStderr("reboot is required; rebooting ...")
					return callback.RebootDevice(cmd.Context())
				}
			} else {
				if report.RestartRequired {
					callback.WriteStderr("reboot is required!")
				}
			}

			return callback.WriteObject(report)
		},
	}

	setConfigCmd.PersistentFlags().BoolVar(&autorebootArg, "autoreboot", false, "automatically reboot device is necessary")

	setCmd := &cobra.Command{
		Use:   "set",
		Short: "Turn light on, off, or set brigtness level",
	}

	setOnCmd := &cobra.Command{
		Use:   "on",
		Short: "Turn light on",
		RunE: func(cmd *cobra.Command, args []string) error {

			switchID, err := getSwitchID()
			if err != nil {
				return err
			}

			params := &Params{
				ID: *switchID,
				On: &truex,
			}

			client, err := callback.Light()
			if err != nil {
				return err
			}

			report, err := client.Set(cmd.Context(), params)
			if err != nil {
				return err
			}

			return callback.WriteObject(report)
		},
	}

	setOffCmd := &cobra.Command{
		Use:   "off",
		Short: "Turn light off",
		RunE: func(cmd *cobra.Command, args []string) error {

			switchID, err := getSwitchID()
			if err != nil {
				return err
			}

			params := &Params{
				ID: *switchID,
				On: &falsex,
			}

			client, err := callback.Light()
			if err != nil {
				return err
			}

			report, err := client.Set(cmd.Context(), params)
			if err != nil {
				return err
			}

			return callback.WriteObject(report)
		},
	}

	setBrightnessCmd := &cobra.Command{
		Use:   "bright",
		Short: "Sets light brightness",
		RunE: func(cmd *cobra.Command, args []string) error {

			switchID, err := getSwitchID()
			if err != nil {
				return err
			}

			client, err := callback.Light()
			if err != nil {
				return err
			}

			brightness := 0.0

			if len(args) > 0 {
				arg := args[0]
				f, err := strconv.ParseFloat(arg, 8)
				if err != nil {
					return fmt.Errorf("arg %s is not a valid float", arg)
				}
				brightness = f
			}

			params := &Params{
				ID:         *switchID,
				Brightness: &brightness,
			}

			report, err := client.Set(cmd.Context(), params)
			if err != nil {
				return err
			}

			return callback.WriteObject(report)
		},
	}

	toggleCmd := &cobra.Command{
		Use:   "toggle",
		Short: "Toggles switch",
		RunE: func(cmd *cobra.Command, args []string) error {

			switchID, err := getSwitchID()
			if err != nil {
				return err
			}

			client, err := callback.Light()
			if err != nil {
				return err
			}

			report, err := client.Toggle(cmd.Context(), *switchID)
			if err != nil {
				return err
			}

			return callback.WriteObject(report)
		},
	}

	setCmd.AddCommand(toggleCmd, setOnCmd, setOffCmd, setBrightnessCmd)
	rootCmd.AddCommand(getConfigCmd, getStatusCmd, getExampleConfigCmd, setConfigCmd, setCmd)
	return rootCmd
}
