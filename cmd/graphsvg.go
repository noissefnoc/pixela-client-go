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
	Short: "get graph SVG HTML tag",
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
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), false)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		response, err := client.GetGraphSvg(args[0])

		if err != nil {
			fmt.Fprintf(os.Stderr, "request error:\n%v\n", err)
			os.Exit(1)
		}

		// print result
		fmt.Printf("%s\n", response)
	},
}

func init() {
	graphCmd.AddCommand(graphsvgCmd)
}
