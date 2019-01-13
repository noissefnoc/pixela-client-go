package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type GrapupdateOptions struct {
	Name string
	Unit string
	Color string
	Timezone string
	purgeCacheURLs []string
}

var (
	guOptions = &GrapupdateOptions{}
)

// graphupdateCmd represents the graphupdate command
var graphupdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update graph definition.",
	Long: `update graph definition. Usage:

$ pixela graph update <graph id> [--name graph_name] [--unit graph_unit] [--color color_name] [--timezone timezone] [--purgeCacheURLs [url1, url2, ...]]

see official document (https://docs.pixe.la/#/put-graph) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		// TODO: add timezone option later
		if len(args) != 1 {
			fmt.Fprintf(os.Stderr, "argument error: `graph update` requires 1 arguments give %d arguments.\n", len(args))
			os.Exit(1)
		}

		// do request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), false)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
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
			fmt.Fprintf(os.Stderr, "request error:\n%v\n", err)
			os.Exit(1)
		}

		responseJSON, err := json.Marshal(response)

		if err != nil {
			fmt.Fprintf(os.Stderr, "response parse error:\n%v\n", err)
			os.Exit(1)
		}

		// print result
		fmt.Printf("%s\n", responseJSON)
	},
}

func init() {
	graphCmd.AddCommand(graphupdateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// graphupdateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// graphupdateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	graphupdateCmd.Flags().StringVarP(&guOptions.Name, "name", "n", "", "graph name")
	graphupdateCmd.Flags().StringVarP(&guOptions.Unit, "unit", "u", "", "graph unit")
	graphupdateCmd.Flags().StringVarP(&guOptions.Color, "color", "c", "", "graph color (shibafu/momiji/sora/ichou/ajisai/kuro)")
	graphupdateCmd.Flags().StringVarP(&guOptions.Timezone, "timezone", "t", "", "graph timezone")
	graphupdateCmd.Flags().StringArrayVarP(&guOptions.purgeCacheURLs, "purge", "p", nil, "purge cache urls")
}
