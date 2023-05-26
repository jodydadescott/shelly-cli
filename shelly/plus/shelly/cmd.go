package shelly

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/jodydadescott/shelly-manager/shelly/plus/types"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type callback interface {
	WriteObject(any) error
	WriteStderr(s string)
	ReadInput() ([]byte, error)
	Shelly() (*Client, error)
	RebootDevice(ctx context.Context) error
}

func NewCmd(callback callback) *cobra.Command {

	var stageArg string
	var urlArg string
	var newUserArg string
	var newPasswordArg string
	var appendArg string
	var autorebootArg bool

	rootCmd := &cobra.Command{
		Use:   "shelly",
		Short: "Shelly Component",
	}

	getConfigCmd := &cobra.Command{
		Use:   "get-config",
		Short: "Returns config",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.Shelly()
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

			client, err := callback.Shelly()
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

	getInfoCmd := &cobra.Command{
		Use:   "get-info",
		Short: "Returns device info",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.Shelly()
			if err != nil {
				return err
			}

			result, err := client.GetDeviceInfo(cmd.Context())
			if err != nil {
				return err
			}

			return callback.WriteObject(result)
		},
	}

	getMethodsCmd := &cobra.Command{
		Use:   "get-methods",
		Short: "Returns all available RPC methods for device",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.Shelly()
			if err != nil {
				return err
			}

			result, err := client.ListMethods(cmd.Context())
			if err != nil {
				return err
			}

			return callback.WriteObject(result)
		},
	}

	getUpdatesCmd := &cobra.Command{
		Use:   "get-updates",
		Short: "Returns available update info",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.Shelly()
			if err != nil {
				return err
			}

			result, err := client.CheckForUpdate(cmd.Context())
			if err != nil {
				return err
			}

			return callback.WriteObject(result)
		},
	}

	getExampleConfigCmd := &cobra.Command{
		Use:   "get-example-config",
		Short: "generates example Config",
		RunE: func(cmd *cobra.Command, args []string) error {
			return callback.WriteObject(ExampleConfig())
		},
	}

	rebootCmd := &cobra.Command{
		Use:   "reboot",
		Short: "Executes device reboot",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.Shelly()
			if err != nil {
				return err
			}

			err = client.Reboot(cmd.Context())
			if err != nil {
				return err
			}

			return nil
		},
	}

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Returns available update info",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.Shelly()
			if err != nil {
				return err
			}

			report, err := client.Update(cmd.Context(), &ShellyParams{
				Stage: &stageArg,
				Url:   &urlArg,
			})
			if err != nil {
				return err
			}

			return callback.WriteObject(report)

		},
	}

	updateCmd.PersistentFlags().StringVar(&stageArg, "stage", "", "The type of the new version - either stable or beta. By default updates to stable version. Optional")
	updateCmd.PersistentFlags().StringVar(&urlArg, "url", "", "Url address of the update. Optional")

	factoryResetCmd := &cobra.Command{
		Use:   "factory-reset",
		Short: "Executes factory reset",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.Shelly()
			if err != nil {
				return err
			}

			report, err := client.FactoryReset(cmd.Context())
			if err != nil {
				return err
			}

			return callback.WriteObject(report)
		},
	}

	resetWifiConfigCmd := &cobra.Command{
		Use:   "reset-wifi-config",
		Short: "Executes Wifi config reset",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.Shelly()
			if err != nil {
				return err
			}

			report, err := client.ResetWiFiConfig(cmd.Context())
			if err != nil {
				return err
			}

			return callback.WriteObject(report)
		},
	}

	setConfigCmd := &cobra.Command{
		Use:   "set-config",
		Short: "Sets config",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.Shelly()
			if err != nil {
				return err
			}

			b, err := callback.ReadInput()
			if err != nil {
				return err
			}

			var config *ShellyConfig

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

	//TODO

	setAuthCmd := &cobra.Command{
		Use:   "set-auth",
		Short: "Sets Auth credentials",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.Shelly()
			if err != nil {
				return err
			}

			if newPasswordArg == "" {
				return fmt.Errorf("newpassword is required")
			}

			params := &ShellyParams{
				Ha1: &newPasswordArg,
			}

			if newUserArg != "" {
				params.User = &newUserArg
			}

			report, err := client.SetAuth(cmd.Context(), params)
			if err != nil {
				return err
			}

			return callback.WriteObject(report)
		},
	}

	setAuthCmd.PersistentFlags().StringVar(&newUserArg, "newuser", "", "New user. Default is "+types.ShellyUser)
	setAuthCmd.PersistentFlags().StringVar(&newPasswordArg, "newpassword", "", "new cleartext password or SHA256(user:realm:password). If cleartext hash will be done automatically.")

	resetAuthCmd := &cobra.Command{
		Use:   "reset-auth",
		Short: "Resets Auth credentials to factory default (no auth)",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.Shelly()
			if err != nil {
				return err
			}

			params := &ShellyParams{}

			report, err := client.SetAuth(cmd.Context(), params)
			if err != nil {
				return err
			}

			return callback.WriteObject(report)
		},
	}

	getAppend := func() (bool, error) {

		switch strings.ToLower(appendArg) {

		case "true":
			return true, nil

		case "false":
			return false, nil

		default:
			return false, fmt.Errorf("append must be set to true or false")

		}

	}

	putTlsClientCertCmd := &cobra.Command{
		Use:   "put-tls-client-cert",
		Short: "Sets TLS Client Cert",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.Shelly()
			if err != nil {
				return err
			}

			append, err := getAppend()
			if err != nil {
				return err
			}

			b, err := callback.ReadInput()
			if err != nil {
				return err
			}

			data := string(b)

			report, err := client.PutTLSClientCert(cmd.Context(), &ShellyParams{
				Data:   &data,
				Append: &append,
			})
			if err != nil {
				return err
			}

			return callback.WriteObject(report)
		},
	}

	appendMsg := "true if more data will be appended afterwards, default false"

	putTlsClientCertCmd.PersistentFlags().StringVar(&appendArg, "append", "", appendMsg)

	putTlsClientKeyCmd := &cobra.Command{
		Use:   "put-tls-client-key",
		Short: "Sets TLS Client Key",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.Shelly()
			if err != nil {
				return err
			}

			append, err := getAppend()
			if err != nil {
				return err
			}

			b, err := callback.ReadInput()
			if err != nil {
				return err
			}

			data := string(b)

			report, err := client.PutTLSClientKey(cmd.Context(), &ShellyParams{
				Data:   &data,
				Append: &append,
			})
			if err != nil {
				return err
			}

			return callback.WriteObject(report)
		},
	}

	putTlsClientKeyCmd.PersistentFlags().StringVar(&appendArg, "append", "", appendMsg)

	putUserCACmd := &cobra.Command{
		Use:   "put-user-ca",
		Short: "Sets Users CA",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.Shelly()
			if err != nil {
				return err
			}

			append, err := getAppend()
			if err != nil {
				return err
			}

			b, err := callback.ReadInput()
			if err != nil {
				return err
			}

			data := string(b)

			report, err := client.PutUserCA(cmd.Context(), &ShellyParams{
				Data:   &data,
				Append: &append,
			})
			if err != nil {
				return err
			}

			return callback.WriteObject(report)
		},
	}

	putUserCACmd.PersistentFlags().StringVar(&appendArg, "append", "", appendMsg)

	rootCmd.AddCommand(getConfigCmd, getStatusCmd, getInfoCmd, getMethodsCmd,
		getUpdatesCmd, getExampleConfigCmd, rebootCmd, updateCmd,
		factoryResetCmd, resetWifiConfigCmd, setConfigCmd, setAuthCmd, resetAuthCmd,
		putTlsClientCertCmd, putTlsClientKeyCmd, putUserCACmd)
	return rootCmd
}
