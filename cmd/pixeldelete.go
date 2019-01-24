package cmd

import (
	"encoding/json"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// pixeldeleteCmd represents the pixeldelete command
var pixeldeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete pixel",
	Long: `delete pixel. Usage:

$ pixela pixel delete <graph id> <date>

see official document (https://docs.pixe.la/#/delete-pixel) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		if len(args) != 2 {
			cmd.Printf("argument error: `pixel delete` requires 2 arguments give %d arguments.\n\n", len(args))
			cmd.Help()
			os.Exit(1)
		}

		// do request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

		if err != nil {
			cmd.Printf("%v\n", err)
			os.Exit(1)
		}

		response, err := client.DeletePixel(args[0], args[1])

		if err != nil {
			cmd.Printf("request error:\n%v\n", err)
			os.Exit(1)
		}

		responseJSON, err := json.Marshal(response)

		if err != nil {
			cmd.Printf("response parse error:\n%v\n", err)
			os.Exit(1)
		}

		// print result in verbose mode
		if viper.GetBool("verbose") {
			cmd.Printf("%s\n", responseJSON)
		}
	},
}

func init() {
	pixelCmd.AddCommand(pixeldeleteCmd)
}
