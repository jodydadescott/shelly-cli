package bluetooth

import (
	"context"
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cobra"
)

type callback interface {
	WriteObject(any) error
	WriteStderr(s string)
	ReadInput() ([]byte, error)
	Bluetooth() (*Client, error)
	RebootDevice(ctx context.Context) error
}

func NewCmd(callback callback) *cobra.Command {

	// var filenameArg string
	var autorebootArg bool

	rootCmd := &cobra.Command{
		Use:   "ble",
		Short: "Bluetooth Component",
	}

	getConfigCmd := &cobra.Command{
		Use:   "get-config",
		Short: "Returns config",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.Bluetooth()
			if err != nil {
				return err
			}

			result, err := client.GetConfig(cmd.Context())
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

			client, err := callback.Bluetooth()
			if err != nil {
				return err
			}

			result, err := client.GetStatus(cmd.Context())
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

			client, err := callback.Bluetooth()
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

			report, err := client.SetConfig(cmd.Context(), config)
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

	rootCmd.AddCommand(getConfigCmd, getStatusCmd, getExampleConfigCmd, setConfigCmd)
	return rootCmd
}
