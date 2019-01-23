package cmd

import (
	"github.com/spf13/cobra"
)

// graphCmd represents the graph command
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "handle graph subcommands (create, def, svg, update and delete)",
	Long: `create, def(get definitions), svg, update and delete graph
see official document (https://docs.pixe.la) for more detail`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(graphCmd)
}
