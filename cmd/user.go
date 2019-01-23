package cmd

import (
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "handle user subcommands (create, update and delete)",
	Long: `create, update (token information) and delete pixe.la user.
see official document (https://docs.pixe.la) for more detail`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
