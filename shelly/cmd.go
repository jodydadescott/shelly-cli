package shelly

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/hokaccha/go-prettyjson"
	"github.com/jodydadescott/shelly-manager/shelly/plus"
	"github.com/jodydadescott/shelly-manager/shelly/plus/types"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Cmd struct {
	hostnameArg string
	usernameArg string
	passwordArg string
	outputArg   string
	inputArg    string
	formatArg   string
	*cobra.Command
	debugEnabledArg bool
}

func NewCmd() *Cmd {

	d := &Cmd{
		Command: &cobra.Command{},
	}

	d.PersistentFlags().StringVarP(&d.hostnameArg, "hostname", "H", "", "hostname; can also be set with env var '"+ShellyHostnameEnvVar+"'")
	d.PersistentFlags().StringVarP(&d.usernameArg, "username", "u", "", "username; can also be set with env var '"+ShellyUsernameEnvVar+"'. Default is '"+types.ShellyUser+"'")
	d.PersistentFlags().StringVarP(&d.formatArg, "format", "f", "", "Supported formats: prettyjson, json, yaml")
	d.PersistentFlags().StringVarP(&d.outputArg, "output", "o", "", "output filename, default is STDOUT")
	d.PersistentFlags().StringVarP(&d.inputArg, "input", "i", "", "input filename")
	d.PersistentFlags().StringVarP(&d.passwordArg, "password", "p", "", "password; can also be set with env var '"+ShellyPasswordEnvVar+"'")
	d.PersistentFlags().BoolVarP(&d.debugEnabledArg, "debug", "d", false, "debug to STDERR")

	d.AddCommand(plus.NewCmd(d))
	return d
}

func (t *Cmd) WriteObject(obj any) error {

	// If output is STDOUT (default) then default format is Pretty JSON

	switch strings.ToLower(t.outputArg) {

	case "stdout", "":

		switch strings.ToLower(t.formatArg) {

		case "prettyjson", "":
			data, err := prettyjson.Marshal(obj)
			if err != nil {
				return err
			}
			fmt.Println(strings.TrimSpace(string(data)))
			return nil

		case "json":
			data, err := json.Marshal(obj)
			if err != nil {
				return err
			}
			fmt.Println(strings.TrimSpace(string(data)))
			return nil

		case "yaml":
			data, err := yaml.Marshal(obj)
			if err != nil {
				return err
			}
			fmt.Println(strings.TrimSpace(string(data)))
			return nil

		default:
			return fmt.Errorf("format type %s is unknown", t.formatArg)
		}

	}

	// If output is a file then default format is JSON

	switch strings.ToLower(t.formatArg) {

	case "json", "":
		data, err := json.Marshal(obj)
		if err != nil {
			return err
		}
		return os.WriteFile(t.outputArg, data, 0644)

	case "prettyjson":
		data, err := prettyjson.Marshal(obj)
		if err != nil {
			return err
		}
		return os.WriteFile(t.outputArg, data, 0644)

	case "yaml":
		data, err := yaml.Marshal(obj)
		if err != nil {
			return err
		}
		return os.WriteFile(t.outputArg, data, 0644)

	}

	return fmt.Errorf("format type %s is unknown", t.formatArg)
}

func (t *Cmd) WriteStderr(s string) {
	fmt.Fprintln(os.Stderr, s)
}

func (t *Cmd) ReadInput() ([]byte, error) {
	// Read input from either a file or STDIN

	if t.inputArg == "" {
		return io.ReadAll(os.Stdin)
	}

	return os.ReadFile(t.inputArg)
}

func (t *Cmd) GetHostname() string {

	if t.hostnameArg != "" {
		return t.hostnameArg
	}

	return os.Getenv(ShellyHostnameEnvVar)
}

func (t *Cmd) GetUsername() string {

	if t.usernameArg != "" {
		return t.usernameArg
	}

	username := os.Getenv(ShellyUsernameEnvVar)
	if username != "" {
		return username
	}

	return types.ShellyUser
}

func (t *Cmd) GetPassword() string {

	if t.passwordArg != "" {
		return t.passwordArg
	}

	return os.Getenv(ShellyPasswordEnvVar)
}

func (t *Cmd) IsDebugEnabled() bool {
	return t.debugEnabledArg
}
