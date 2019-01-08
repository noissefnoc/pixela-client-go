package cmd

import (
	"github.com/noissefnoc/pixela-client-go/pixel"

	"github.com/spf13/cobra"
)

// pixelrecordCmd represents the pixelrecord command
var pixelrecordCmd = &cobra.Command{
	Use:   "record",
	Short: "record quantity to graph.",
	Long:  `record quantity to graph.`,
	Run:   pixel.Record,
}

func init() {
	pixelCmd.AddCommand(pixelrecordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pixelrecordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pixelrecordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
