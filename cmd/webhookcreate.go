package cmd

import (
	"encoding/json"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// webhookcreateCmd represents the webhookcreate command
var webhookcreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create webhook",
	Long: `create webhook. Usage:

$ pixela webhook create <graph id> <type>

see official document (https://docs.pixe.la/#/post-webhook) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		// TODO: add timezone option later
		if len(args) != 2 {
			cmd.Printf("argument error: `webhook create` requires 2 arguments give %d arguments.\n\n", len(args))
			cmd.Help()
			os.Exit(1)
		}

		// do request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

		if err != nil {
			cmd.Printf("%v\n", err)
			os.Exit(1)
		}

		response, err := client.CreateWebhook(args[0], args[1])

		if err != nil {
			cmd.Printf("request error: %v\n", err)
			os.Exit(1)
		}

		responseJSON, err := json.Marshal(response)

		if err != nil {
			cmd.Printf("response parse error: %v\n", err)
			os.Exit(1)
		}

		// print result in verbose mode
		if viper.GetBool("verbose") {
			cmd.SetOutput(os.Stdout)
			cmd.Printf("%s\n", responseJSON)
		}
	},
}

func init() {
	webhookCmd.AddCommand(webhookcreateCmd)
}
