package cmd

import (
	"encoding/json"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type GraphCreateOptions struct {
	Timezone string
}

// graphcreateCmd represents the graphcreate command
var graphcreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create pixe.la graph.",
	Long: `create pixe.la graph. Usage:

$ pixela graph create <graph id> <graph name> <unit> <type> <color> [--timezone timezone] [--selfSufficient selfSufficient]

see official document (https://docs.pixe.la/#/post-graph) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		if len(args) != 5 {
			cmd.Printf("argument error: `graph create` requires 5 arguments give %d arguments.\n\n", len(args))
			cmd.Help()
			os.Exit(1)
		}

		// make request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

		if err != nil {
			cmd.Printf("%v\n", err)
			os.Exit(1)
		}

		// get flags
		timezone, err := cmd.Flags().GetString("timezone")

		if err != nil {
			cmd.Printf("%v\n", err)
			os.Exit(1)
		}

		selfSufficient, err := cmd.Flags().GetString("selfSufficient")

		if err != nil {
			cmd.Printf("%v\n", err)
			os.Exit(1)
		}

		// do request
		response, err := client.CreateGraph(args[0], args[1], args[2], args[3], args[4], timezone, selfSufficient)

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
	graphCmd.AddCommand(graphcreateCmd)

	graphcreateCmd.Flags().StringP( "timezone", "", "", "timezone")
	graphcreateCmd.Flags().StringP( "selfSufficient", "", "none", "selfSufficient")
}
