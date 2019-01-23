package cmd

import (
	"github.com/spf13/cobra"
)

// pixelCmd represents the pixel command
var pixelCmd = &cobra.Command{
	Use:   "pixel",
	Short: "handle pixel subcommands (create, get, update, increment, decrement and delete)",
	Long: `record, get, update, increment, decrement and delete pixel
see official document (https://docs.pixe.la) for more detail`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(pixelCmd)
}
