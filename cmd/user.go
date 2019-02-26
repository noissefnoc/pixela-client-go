package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
)

type UserCreateOptions struct {
	AgreeTermsOfService string
	NotMinor            string
}

func newUserCmd() *cobra.Command {
	userCmd := &cobra.Command{
		Use:   "user",
		Short: "handle user subcommands (create, update and delete)",
		Long: `create, update (token information) and delete pixe.la user.
see official document (https://docs.pixe.la) for more detail`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}

	userCmd.AddCommand(newUserCreateCmd())
	userCmd.AddCommand(newUserUpdateCmd())
	userCmd.AddCommand(newUserDeleteCmd())

	return userCmd
}

func newUserCreateCmd() *cobra.Command {
	userCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "create pixe.la user",
		Long: `create pixe.la user. Usage:

$ pixela user create <username> <token> [--agreeTermsOfService yes/no (default:yes)] [--notMinor yes/no (default:yes)]

see official document (https://docs.pixe.la/#/post-user) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			if len(args) != 2 {
				return errors.New(fmt.Sprintf("argument error: `user create` requires 2 arguments give %d arguments.", len(args)))
			}

			username := args[0]
			token := args[1]

			// do request
			client, err := pixela.New(username, token, viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			agreeTermsOfService, _ := cmd.Flags().GetString("agreeTermsOfService")
			notMinor, _ := cmd.Flags().GetString("notMinor")

			response, err := client.CreateUser(agreeTermsOfService, notMinor)

			if err != nil {
				return errors.Wrap(err, "request error")
			}

			responseJSON, err := json.Marshal(response)

			if err != nil {
				return errors.Wrap(err, "response parse error")
			}

			// print result in verbose mode
			if viper.GetBool("verbose") {
				cui.Outputln(string(responseJSON))
			}

			// save authentications into file
			configFilePath := viper.GetString("config")
			err = saveConfigFile(configFilePath, username, token)

			if err != nil {
				return errors.Wrap(err, "save config file error")
			}

			return nil
		},
	}

	userCreateCmd.Flags().StringP("agreeTermsOfService", "", "yes", "agree terms of service.")
	userCreateCmd.Flags().StringP("notMinor", "", "yes", "usage is not minor.")

	return userCreateCmd
}

func newUserUpdateCmd() *cobra.Command {
	userUpdateCmd := &cobra.Command{
		Use:   "update",
		Short: "update user token",
		Long: `update pixe.la user token. Usage:

$ pixela user update <newToken>

see official document (https://docs.pixe.la/#/put-user) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			if len(args) != 1 {
				return errors.New(fmt.Sprintf("argument error: `user update` requires 1 arguments give %d arguments.", len(args)))
			}

			// do request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			response, err := client.UpdateUser(args[0])

			if err != nil {
				return errors.Wrap(err, "request error")
			}

			responseJSON, err := json.Marshal(response)

			if err != nil {
				return errors.Wrap(err, "response parse error")
			}

			// print result in verbose mode
			if viper.GetBool("verbose") {
				cui.Outputln(string(responseJSON))
			}

			return nil
		},
	}

	return userUpdateCmd
}

func newUserDeleteCmd() *cobra.Command {
	userDeleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete pixe.la user",
		Long: `delete pixe.la user. Usage:

$ pixela user delete <username> <token>

see official document (https://docs.pixe.la/#/delete-user) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			if len(args) != 2 {
				return errors.New(fmt.Sprintf("argument error: `user update` requires 2 arguments give %d arguments.", len(args)))
			}

			// do request
			client, err := pixela.New(args[0], args[1], viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			response, err := client.DeleteUser()

			if err != nil {
				return errors.Wrap(err, "request error")
			}

			responseJSON, err := json.Marshal(response)

			if err != nil {
				return errors.Wrap(err, "response parse error")
			}

			// print result in verbose mode
			if viper.GetBool("verbose") {
				cui.Outputln(string(responseJSON))
			}

			return nil
		},
	}

	return userDeleteCmd
}

// check if file exists
func existFile(checkFilePath string) bool {
	_, err := os.Stat(checkFilePath)

	if err == nil {
		return true
	} else {
		return false
	}
}

// default config file path
func getDefaultFilePath() (string, error) {
	homeDirPath, err := homedir.Dir()

	if err != nil {
		return "", errors.Wrap(err, "cannot find home directory path.")
	}

	return filepath.Join(homeDirPath, ".pixela.yaml"), nil
}

// save username and token to file
func saveConfigFile(path, username, token string) error {
	defaultPath, err := getDefaultFilePath()

	if err != nil {
		return err
	}

	if existFile(path) && existFile(defaultPath) {
		cui.Outputln(fmt.Sprintf("`pixel create` successd but there are aleady config file at %s", path))
		cui.Outputln(fmt.Sprintf("You should take a note following settings(save to yaml file)."))
		cui.Outputln(fmt.Sprintf("username: %s", username))
		cui.Outputln(fmt.Sprintf("token: %s", token))

		return errors.New("cannot save configs.")
	}

	saveConfigFilePath := path

	if saveConfigFilePath == "" {
		saveConfigFilePath = defaultPath
	}

	// TODO: apply yaml package
	settings := fmt.Sprintf("username: %s\ntoken: %s\n", username, token)

	err = ioutil.WriteFile(saveConfigFilePath, []byte(settings), 0644)

	if err != nil {
		return err
	}

	return nil
}
