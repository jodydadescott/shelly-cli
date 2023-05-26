package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type callback interface {
	WriteObject(any) error
	WriteStderr(s string)
	ReadInput() ([]byte, error)
	WebHook() (*Client, error)
	RebootDevice(ctx context.Context) error
}

func NewCmd(callback callback) *cobra.Command {

	rootCmd := &cobra.Command{
		Use:   "proto",
		Short: "Proto Component",
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List WebHooks",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.WebHook()
			if err != nil {
				return err
			}

			result, err := client.List(cmd.Context())
			if err != nil {
				return err
			}

			return callback.WriteObject(result)
		},
	}

	listSupportedCmd := &cobra.Command{
		Use:   "list-supported",
		Short: "List supported events",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.WebHook()
			if err != nil {
				return err
			}

			result, err := client.ListSupported(cmd.Context())
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

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Creates WebHook",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.WebHook()
			if err != nil {
				return err
			}

			b, err := callback.ReadInput()
			if err != nil {
				return err
			}

			var params *Params

			var errors *multierror.Error

			err = json.Unmarshal(b, &params)
			if err != nil {
				errors = multierror.Append(errors, err)
				err = yaml.Unmarshal(b, &params)

				if err != nil {
					errors = multierror.Append(errors, err)
					errors = multierror.Append(errors, fmt.Errorf("Invalid format. Expect JSON or YAML"))
					return errors.ErrorOrNil()
				}
			}

			report, err := client.Create(cmd.Context(), params)
			if err != nil {
				return err
			}

			return callback.WriteObject(report)
		},
	}

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Updates WebHook",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.WebHook()
			if err != nil {
				return err
			}

			b, err := callback.ReadInput()
			if err != nil {
				return err
			}

			var params *Params

			var errors *multierror.Error

			err = json.Unmarshal(b, &params)
			if err != nil {
				errors = multierror.Append(errors, err)
				err = yaml.Unmarshal(b, &params)

				if err != nil {
					errors = multierror.Append(errors, err)
					errors = multierror.Append(errors, fmt.Errorf("Invalid format. Expect JSON or YAML"))
					return errors.ErrorOrNil()
				}
			}

			report, err := client.Update(cmd.Context(), params)
			if err != nil {
				return err
			}

			return callback.WriteObject(report)
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Deletes WebHook with specified ID",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.WebHook()
			if err != nil {
				return err
			}

			if len(args) <= 0 {
				return fmt.Errorf("Webhook ID is required")
			}

			arg := args[0]
			id, err := strconv.Atoi(arg)
			if err != nil {
				return fmt.Errorf("ID must be an integer, %s is not valid", arg)
			}

			report, err := client.Delete(cmd.Context(), &Webhooks{
				ID: id,
			})
			if err != nil {
				return err
			}

			return callback.WriteObject(report)
		},
	}

	deleteAllCmd := &cobra.Command{
		Use:   "delete-all",
		Short: "Deletes all WebHooks",
		RunE: func(cmd *cobra.Command, args []string) error {

			client, err := callback.WebHook()
			if err != nil {
				return err
			}

			report, err := client.DeleteAll(cmd.Context())
			if err != nil {
				return err
			}

			return callback.WriteObject(report)
		},
	}

	rootCmd.AddCommand(listCmd, listSupportedCmd, getExampleConfigCmd, createCmd, updateCmd, deleteCmd, deleteAllCmd)
	return rootCmd
}
