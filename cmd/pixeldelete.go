package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// pixeldeleteCmd represents the pixeldelete command
var pixeldeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete pixel.",
	Long: `delete pixel. Usage:

$ pixela pixel delete <graph id> <date>

see official document (https://docs.pixe.la/#/delete-pixel) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		if len(args) != 2 {
			fmt.Fprintf(os.Stderr, "argument error: `pixel delete` requires 2 arguments give %d arguments.\n", len(args))
			os.Exit(1)
		}

		// do request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), false)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		response, err := client.DeletePixel(args[0], args[1])

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
	pixelCmd.AddCommand(pixeldeleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pixeldeleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pixeldeleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
