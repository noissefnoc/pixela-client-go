package cmd

import (
	"github.com/spf13/cobra"
)

// webhookCmd represents the webhook command
var webhookCmd = &cobra.Command{
	Use:   "webhook",
	Short: "handle webhook subcommands (create, get, invoke and delete)",
	Long: `create, get invoke and delete pixe.la webhook.
see official document (https://docs.pixe.la) for more detail`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(webhookCmd)
}
