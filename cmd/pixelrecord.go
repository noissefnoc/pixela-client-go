package cmd

import (
	"fmt"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// pixelrecordCmd represents the pixelrecord command
var pixelrecordCmd = &cobra.Command{
	Use:   "record",
	Short: "record quantity to graph.",
	Long:  `record quantity to graph. Usage:

$ pixela pixel record <graph id> <date> <quantity>

see official document (https://docs.pixe.la/#/post-pixel) for more detail.`,
	Run:   func(cmd *cobra.Command, args []string) {
		// check arguments
		if len(args) != 3 {
			fmt.Fprintf(os.Stderr, "argument error: `pixel record` requires 3 arguments give %d arguments.\n", len(args))
			os.Exit(1)
		}

		// rdo request
		client := pixela.Pixela{
			Username: viper.GetString("username"),
			Token: viper.GetString("token"),
			Debug: true,
		}

		err := client.RecordPixel(args[0], args[1], args[2])

		if err != nil {
			fmt.Fprintf(os.Stderr, "pixela request error:\n%v\n", err)
			os.Exit(1)
		}
	},
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
