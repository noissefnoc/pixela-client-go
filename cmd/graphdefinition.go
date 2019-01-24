package cmd

import (
	"encoding/json"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// graphdefinitionCmd represents the graphdefinition command
var graphdefinitionCmd = &cobra.Command{
	Use:   "def",
	Short: "get graph definitions",
	Long: `get graph definitions. Usage:

$ pixela graph def

see official document (https://docs.pixe.la/#/get-graph) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		if len(args) != 0 {
			cmd.Printf("argument error: `graph def` does not accept any argument.\n\n")
			os.Exit(1)
		}

		// do request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

		if err != nil {
			cmd.Printf("%v\n", err)
			os.Exit(1)
		}

		response, err := client.GetGraphDefinition()

		if err != nil {
			cmd.Printf("request error:\n%v\n", err)
			os.Exit(1)
		}

		responseJSON, err := json.Marshal(response)

		if err != nil {
			cmd.Printf("pixela response parse error:\n%v\n", err)
			os.Exit(1)
		}

		// print result in verbose mode
		if viper.GetBool("verbose") {
			cmd.Printf("%s\n", responseJSON)
		}
	},
}

func init() {
	graphCmd.AddCommand(graphdefinitionCmd)
}
