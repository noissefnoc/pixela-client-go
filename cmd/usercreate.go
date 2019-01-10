package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
)

// usercreateCmd represents the usercreate command
var usercreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create pixe.la user.",
	Long: `create pixe.la user. Usage:

$ pixela user create <username> <token>

see official document (https://docs.pixe.la/#/post-user) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		if len(args) != 2 {
			fmt.Fprintf(
				os.Stderr,"argument error: `user create` requires 2 arguments give %d arguments.\n", len(args))
			os.Exit(1)
		}

		username := args[0]
		token := args[1]

		// do request
		client := pixela.Pixela{
			Username: username,
			Token:    token,
			Debug:    true,
		}

		err := client.CreateUser()

		if err != nil {
			fmt.Fprintf(os.Stderr, "pixela request error:\n%v\n", err)
			os.Exit(1)
		}

		// save authentications into file
		configFilePath := viper.GetString("config")

		if !existFile(configFilePath) && !existFile(getDefaultFilePath()) {
			saveConfigFilePath := configFilePath

			if saveConfigFilePath == "" {
				saveConfigFilePath = getDefaultFilePath()
			}

			// TODO: apply yaml package
			settings := fmt.Sprintf("username: %s\ntoken: %s\n", username, token)

			err := ioutil.WriteFile(saveConfigFilePath, []byte(settings), 0644)

			if err != nil {
				fmt.Fprintf(os.Stderr, "write config error:\n%+v\n", err)
				os.Exit(1)
			}
		} else {
			fmt.Fprintf(os.Stderr, "`pixel create` successd but there are aleady config file at %s\n", configFilePath)
			fmt.Fprintf(os.Stderr, "You should take a note following settings(save to yaml file).\n")
			fmt.Fprintf(os.Stderr, "username: %s\n", username)
			fmt.Fprintf(os.Stderr, "token: %s\n", token)
		}
	},
}

func init() {
	userCmd.AddCommand(usercreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// usercreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// usercreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
func getDefaultFilePath() string {
	homeDirPath, err := homedir.Dir()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return filepath.Join(homeDirPath, ".pixela.yaml")
}
