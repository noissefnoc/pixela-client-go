package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// graphcreateCmd represents the graphcreate command
var graphcreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create pixe.la graph.",
	Long: `create pixe.la graph. Usage:

$ pixela graph create <graph id> <graph name> <unit> <type> <color> [timezone]

see official document (https://docs.pixe.la/#/post-graph) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		// TODO: add timezone option later
		if len(args) != 5 && len(args) != 6 {
			fmt.Fprintf(os.Stderr, "argument error: `graph create` requires 5 arguments give %d arguments.\n", len(args))
			os.Exit(1)
		}

		// do request
		client := pixela.Pixela{
			Username: viper.GetString("username"),
			Token:    viper.GetString("token"),
			Debug:    true,
		}

		// optional timezone settings
		var timeZone string

		if len(args) == 5 {
			timeZone = ""
		}

		response, err := client.CreateGraph(args[0], args[1], args[2], args[3], args[4], timeZone)

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
	graphCmd.AddCommand(graphcreateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// graphcreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// graphcreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
