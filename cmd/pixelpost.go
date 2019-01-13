package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// pixelpostCmd represents the pixelrecord command
var pixelpostCmd = &cobra.Command{
	Use:   "post",
	Short: "post quantity to graph.",
	Long:  `post quantity to graph. Usage:

$ pixela pixel post <graph id> <date> <quantity>

see official document (https://docs.pixe.la/#/post-pixel) for more detail.`,
	Run:   func(cmd *cobra.Command, args []string) {
		// check arguments
		if len(args) != 3 {
			fmt.Fprintf(os.Stderr, "argument error: `pixel post` requires 3 arguments give %d arguments.\n", len(args))
			os.Exit(1)
		}

		// do request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), false)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		response, err := client.PostPixel(args[0], args[1], args[2])

		if err != nil {
			fmt.Fprintf(os.Stderr, "request error:\n%v\n", err)
			os.Exit(1)
		}

		responseJSON, err := json.Marshal(response)

		if err != nil {
			fmt.Fprintf(os.Stderr, "response parse error:\n%v\n", err)
			os.Exit(1)
		}

		// print result
		fmt.Printf("%s\n", responseJSON)
	},
}

func init() {
	pixelCmd.AddCommand(pixelpostCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pixelpostCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pixelpostCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
