package cmd

import (
	"encoding/json"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type GraphUpdateOptions struct {
	Name string
	Unit string
	Color string
	Timezone string
	purgeCacheURLs []string
}

var (
	guOptions = &GraphUpdateOptions{}
)

// graphupdateCmd represents the graphupdate command
var graphupdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update graph definition",
	Long: `update graph definition. Usage:

$ pixela graph update <graph id> [--name graph_name] [--unit graph_unit] [--color color_name] [--timezone timezone] [--purgeCacheURLs [url1, url2, ...]]

see official document (https://docs.pixe.la/#/put-graph) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		// TODO: add timezone option later
		if len(args) != 1 {
			cmd.Printf("argument error: `graph update` requires 1 arguments give %d arguments.\n\n", len(args))
			cmd.Help()
			os.Exit(1)
		}

		// do request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

		if err != nil {
			cmd.Printf("%v\n", err)
			os.Exit(1)
		}

		pl := pixela.UpdateGraphPayload{
			Name: guOptions.Name,
			Unit: guOptions.Unit,
			Color: guOptions.Color,
			Timezone: guOptions.Timezone,
			PurgeCacheURLs: guOptions.purgeCacheURLs,
		}

		response, err := client.UpdateGraph(args[0], pl)

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
	graphCmd.AddCommand(graphupdateCmd)

	graphupdateCmd.Flags().StringVarP(&guOptions.Name, "name", "n", "", "graph name")
	graphupdateCmd.Flags().StringVarP(&guOptions.Unit, "unit", "u", "", "graph unit")
	graphupdateCmd.Flags().StringVarP(&guOptions.Color, "color", "c", "", "graph color (shibafu/momiji/sora/ichou/ajisai/kuro)")
	graphupdateCmd.Flags().StringVarP(&guOptions.Timezone, "timezone", "t", "", "graph timezone")
	graphupdateCmd.Flags().StringArrayVarP(&guOptions.purgeCacheURLs, "purge", "p", nil, "purge cache urls")
}
