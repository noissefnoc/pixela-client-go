package cmd

import (
	"encoding/json"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// webhookdeleteCmd represents the webhookdelete command
var webhookdeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete webhook",
	Long: `delete webhook. Usage:

$ pixela webhook delete <webhook hash>

see official document (https://docs.pixe.la/#/delete-webhook) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		// TODO: add timezone option later
		if len(args) != 1 {
			cmd.Printf("argument error: `webhook delete` requires 1 arguments give %d arguments.\n\n", len(args))
			cmd.Help()
			os.Exit(1)
		}

		// do request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

		if err != nil {
			cmd.Printf("%v\n", err)
			os.Exit(1)
		}

		response, err := client.DeleteWebhook(args[0])

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
	webhookCmd.AddCommand(webhookdeleteCmd)
}
