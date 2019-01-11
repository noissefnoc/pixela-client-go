package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// graphdefinitionCmd represents the graphdefinition command
var graphdefinitionCmd = &cobra.Command{
	Use:   "def",
	Short: "get graph definitions.",
	Long: `get graph definitions. Usage:

$ pixela graph def

see official document (https://docs.pixe.la/#/get-graph) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		if len(args) != 0 {
			fmt.Fprintf(os.Stderr, "argument error: `graph def` does not accept any argument.\n")
			os.Exit(1)
		}

		// do request
		client := pixela.Pixela{
			Username: viper.GetString("username"),
			Token: viper.GetString("token"),
			Debug: true,
		}

		response, err := client.GetGraphDefinition()

		if err != nil {
			fmt.Fprintf(os.Stderr, "pixela request error:\n%v\n", err)
			os.Exit(1)
		}

		responseJSON, err := json.Marshal(response)

		if err != nil {
			fmt.Fprintf(os.Stderr, "pixela response parse error:\n%v\n", err)
			os.Exit(1)
		}

		// print result
		fmt.Printf("%s\n", responseJSON)
	},
}

func init() {
	graphCmd.AddCommand(graphdefinitionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// graphdefinitionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// graphdefinitionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
