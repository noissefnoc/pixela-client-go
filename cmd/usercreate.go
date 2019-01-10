package cmd

import (
	"fmt"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"os"
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
			fmt.Fprintf(os.Stderr, "argument error: `user create` requires 2 arguments give %d arguments.\n", len(args))
			os.Exit(1)
		}

		// do request
		client := pixela.Pixela{
			Username: args[0],
			Token:    args[1],
			Debug:    true,
		}

		err := client.CreateUser()

		if err != nil {
			fmt.Fprintf(os.Stderr, "pixela request error:\n%v\n", err)
			os.Exit(1)
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
