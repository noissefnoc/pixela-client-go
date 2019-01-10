package cmd

import (
	"fmt"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// pixelgetCmd represents the pixelget command
var pixelgetCmd = &cobra.Command{
	Use:   "get",
	Short: "get pixel value.",
	Long: `get pixel value. Usage:

$ pixela pixel get <graph id> <date>

see official document (https://docs.pixe.la/#/get-pixel) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		if len(args) != 2 {
			fmt.Fprintf(os.Stderr, "argument error: `pixel record` requires 2 arguments give %d arguments.\n", len(args))
			os.Exit(1)
		}

		// do request
		client := pixela.Pixela{
			Username: viper.GetString("username"),
			Token: viper.GetString("token"),
			Debug: true,
		}

		response, err := client.GetPixel(args[0], args[1])

		if err != nil {
			fmt.Fprintf(os.Stderr, "pixela request error:\n%v\n", err)
			os.Exit(1)
		}

		// print result
		fmt.Printf("%+v\n", response)
	},
}

func init() {
	pixelCmd.AddCommand(pixelgetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pixelgetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pixelgetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
