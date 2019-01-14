package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// webhookcreateCmd represents the webhookcreate command
var webhookcreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create webhook.",
	Long: `create webhook. Usage:

$ pixela webhook create <graph id> <type>

see official document (https://docs.pixe.la/#/post-webhook) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		// TODO: add timezone option later
		if len(args) != 2 {
			fmt.Fprintf(os.Stderr, "argument error: `webhook create` requires 2 arguments give %d arguments.\n", len(args))
			os.Exit(1)
		}

		// do request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), false)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		response, err := client.CreateWebhook(args[0], args[1])

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
	webhookCmd.AddCommand(webhookcreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webhookcreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webhookcreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
