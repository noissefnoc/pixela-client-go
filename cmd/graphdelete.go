package cmd

import (
	"encoding/json"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// graphdeleteCmd represents the graphdelete command
var graphdeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete graph",
	Long: `delete graph. Usage:

$ pixela graph delete <graph id>

see official document (https://docs.pixe.la/#/delete-graph) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		if len(args) != 1 {
			cmd.Printf("argument error: `graph delete` requires 1 arguments give %d arguments.\n\n", len(args))
			cmd.Help()
			os.Exit(1)
		}

		// do request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

		if err != nil {
			cmd.Printf("%v\n", err)
			os.Exit(1)
		}

		response, err := client.DeleteGraph(args[0])

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
	graphCmd.AddCommand(graphdeleteCmd)
}
