package cmd

import (
	"encoding/json"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// pixelincCmd represents the pixelinc command
var pixelincCmd = &cobra.Command{
	Use:   "inc",
	Short: "increment pixel quantity",
	Long: `increment pixel quantity. Usage:

$ pixela pixel inc <graph id>

see official document (https://docs.pixe.la/#/increment-pixel) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		if len(args) != 1 {
			cmd.Printf("argument error: `pixel inc` requires 1 arguments give %d arguments.\n\n", len(args))
			cmd.Help()
			os.Exit(1)
		}

		// do request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

		if err != nil {
			cmd.Printf("%v\n", err)
			os.Exit(1)
		}

		response, err := client.IncPixel(args[0])

		if err != nil {
			cmd.Printf("request error: %v\n", err)
			os.Exit(1)
		}

		responseJSON, err := json.Marshal(response)

		if err != nil {
			cmd.Printf("response parse error: %v\n", err)
			os.Exit(1)
		}

		// print result in verbose mode
		if viper.GetBool("verbose") {
			cmd.SetOutput(os.Stdout)
			cmd.Printf("%s\n", responseJSON)
		}
	},
}

func init() {
	pixelCmd.AddCommand(pixelincCmd)
}
