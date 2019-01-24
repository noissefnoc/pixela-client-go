package cmd

import (
	"encoding/json"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// pixelupdateCmd represents the pixelupdate command
var pixelupdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update pixel",
	Long: `update pixel. Usage:

$ pixela pixel update <graph id> <date> <quantity>

see official document (https://docs.pixe.la/#/put-pixel) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		if len(args) != 3 {
			cmd.Printf("argument error: `pixel update` requires 3 arguments give %d arguments.\n\n", len(args))
			cmd.Help()
			os.Exit(1)
		}

		// do request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

		if err != nil {
			cmd.Printf("%v\n", err)
			os.Exit(1)
		}

		response, err := client.UpdatePixel(args[0], args[1], args[2])

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
	pixelCmd.AddCommand(pixelupdateCmd)
}
