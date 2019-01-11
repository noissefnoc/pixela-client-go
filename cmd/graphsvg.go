package cmd

import (
	"fmt"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// graphsvgCmd represents the graphsvg command
var graphsvgCmd = &cobra.Command{
	Use:   "svg",
	Short: "get graph SVG HTML tag.",
	Long: `get graph SVG HTML tag. Usage:

$ pixela graph svg <graph id>

see official document (https://docs.pixe.la/#/get-svg) for more detail.`,
	Run: func(cmd *cobra.Command, args []string) {
		// check arguments
		if len(args) != 1 {
			fmt.Fprintf(os.Stderr, "argument error: `pixel record` requires 1 arguments give %d arguments.\n", len(args))
			os.Exit(1)
		}

		// do request
		client := pixela.Pixela{
			Username: viper.GetString("username"),
			Token: viper.GetString("token"),
			Debug: true,
		}

		response, err := client.GetGraphSvg(args[0])

		if err != nil {
			fmt.Fprintf(os.Stderr, "pixela request error:\n%v\n", err)
			os.Exit(1)
		}

		// print result
		fmt.Printf("%s\n", response)
	},
}

func init() {
	graphCmd.AddCommand(graphsvgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// graphsvgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// graphsvgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
