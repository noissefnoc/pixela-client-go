package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// graphCmd represents the graph command
var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "handle graph subcommands",
	Long: `create, definition, get, update and delete graph

see official document (https://docs.pixe.la) for more detail`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: change this statement print `pixela graph` subcommands
		fmt.Println("graph called")
	},
}

func init() {
	rootCmd.AddCommand(graphCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// graphCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// graphCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
