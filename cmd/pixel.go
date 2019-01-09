package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pixelCmd represents the pixel command
var pixelCmd = &cobra.Command{
	Use:   "pixel",
	Short: "handle pixel subcommands",
	Long: `record, get, update, increment, decrement and delete pixel

see official document (https://docs.pixe.la) for more detail`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: change this statement print `pixela pixel` sub-commands
		fmt.Println("pixel called")
	},
}

func init() {
	rootCmd.AddCommand(pixelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pixelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pixelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
