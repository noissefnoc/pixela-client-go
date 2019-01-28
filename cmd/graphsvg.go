package cmd

import (
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type graphSvgOptions struct {
	Date string
	Mode string
}

var gsOptions = &graphSvgOptions{}

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
			cmd.Printf("argument error: `graph svg` requires 1 arguments give %d arguments.\n\n", len(args))
			cmd.Help()
			os.Exit(1)
		}

		// do request
		client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

		if err != nil {
			cmd.Printf("%v\n", err)
			os.Exit(1)
		}

		response, err := client.GetGraphSvg(args[0], gsOptions.Date, gsOptions.Mode)

		if err != nil {
			cmd.Printf("request error: %v\n", err)
			os.Exit(1)
		}

		// print result
		cmd.SetOutput(os.Stdout)
		cmd.Printf("%s\n", response)
	},
}

func init() {
	graphCmd.AddCommand(graphsvgCmd)

	graphsvgCmd.Flags().StringVarP(&gsOptions.Date, "date", "", "", "date")
	graphsvgCmd.Flags().StringVarP(&gsOptions.Mode, "mode", "", "", "mode")
}
