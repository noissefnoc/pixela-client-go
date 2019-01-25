package cmd

import (
	"encoding/json"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// pixelpostCmd represents the pixelrecord command
var pixelpostCmd = &cobra.Command{
	Use:   "create",
	Short: "create pixel",
	Long:  `create pixel. Usage:

$ pixela pixel create <graph id> <date> <quantity>

see official document (https://docs.pixe.la/#/post-pixel) for more detail.`,
	Run:   func(cmd *cobra.Command, args []string) {
		// check arguments
		if len(args) != 3 {
			cmd.Printf("argument error: `pixel create` requires 3 arguments give %d arguments.\n\n", len(args))
			cmd.Help()
			os.Exit(1)
		}

		// do request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

		if err != nil {
			cmd.Printf("%v\n", err)
			os.Exit(1)
		}

		response, err := client.CreatePixel(args[0], args[1], args[2])

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
	pixelCmd.AddCommand(pixelpostCmd)
}
