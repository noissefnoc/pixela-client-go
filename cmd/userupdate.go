package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// userupdateCmd represents the userupdate command
var userupdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update user token",
	Long: `update pixe.la user token. Usage:

$ pixela user update <newToken>

see official document (https://docs.pixe.la/#/put-user) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		if len(args) != 1 {
			fmt.Fprintf(
				os.Stderr,"argument error: `user update` requires 1 arguments give %d arguments.\n", len(args))
			os.Exit(1)
		}

		// do request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), false)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		response, err := client.UpdateUser(args[0])

		if err != nil {
			fmt.Fprintf(os.Stderr, "request error:\n%v\n", err)
			os.Exit(1)
		}

		responseJSON, err := json.Marshal(response)

		if err != nil {
			fmt.Fprintf(os.Stderr, "response parse error:\n%v\n", err)
			os.Exit(1)
		}

		// print result
		fmt.Printf("%s\n", responseJSON)
	},
}

func init() {
	userCmd.AddCommand(userupdateCmd)
}
