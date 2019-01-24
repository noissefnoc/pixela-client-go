package cmd

import (
	"encoding/json"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// pixelgetCmd represents the pixelget command
var pixelgetCmd = &cobra.Command{
	Use:   "get",
	Short: "get pixel value",
	Long: `get pixel value. Usage:

$ pixela pixel get <graph id> <date>

see official document (https://docs.pixe.la/#/get-pixel) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		if len(args) != 2 {
			cmd.Printf("argument error: `pixel record` requires 2 arguments give %d arguments.\n\n", len(args))
			cmd.Help()
			os.Exit(1)
		}

		// do request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

		if err != nil {
			cmd.Printf("%v\n", err)
			os.Exit(1)
		}

		response, err := client.GetPixel(args[0], args[1])

		if err != nil {
			cmd.Printf("request error:\n%v\n", err)
			os.Exit(1)
		}

		responseJSON, err := json.Marshal(response)

		if err != nil {
			cmd.Printf("response parse error:\n%v\n", err)
			os.Exit(1)
		}

		// print result
		cmd.SetOutput(os.Stdout)
		cmd.Printf("%s\n", responseJSON)
	},
}

func init() {
	pixelCmd.AddCommand(pixelgetCmd)
}
