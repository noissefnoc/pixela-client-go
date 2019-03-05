package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newGraphCmd() *cobra.Command {
	// graphCmd represents the graph command
	graphCmd := &cobra.Command{
		Use:   "graph",
		Short: "handle graph subcommands (create, def, svg, update and delete)",
		Long: `create, def(get definitions), svg, update and delete graph
see official document (https://docs.pixe.la) for more detail`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}

	graphCmd.AddCommand(newGraphCreateCmd())
	graphCmd.AddCommand(newGraphUpdateCmd())
	graphCmd.AddCommand(newGraphDeleteCmd())
	graphCmd.AddCommand(newGraphDefCmd())
	graphCmd.AddCommand(newGraphSvgCmd())
	graphCmd.AddCommand(newGraphPixelsDateCmd())

	return graphCmd
}

func newGraphCreateCmd() *cobra.Command {
	graphCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "create pixe.la graph.",
		Long: `create pixe.la graph. Usage:

$ pixela graph create <graph id> <graph name> <unit> <type> <color> [--timezone timezone] [--selfSufficient selfSufficient]

see official document (https://docs.pixe.la/#/post-graph) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 5 {
				return errors.New(fmt.Sprintf("argument error: `graph create` requires 5 arguments give %d arguments.", len(args)))
			}

			// make request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			timezone, _ := cmd.Flags().GetString("timezone")
			selfSufficient, _ := cmd.Flags().GetString("selfSufficient")

			// do request
			response, err := client.CreateGraph(args[0], args[1], args[2], args[3], args[4], timezone, selfSufficient)

			if err != nil {
				return errors.Wrap(err, "request error: ")
			}

			responseJSON, err := json.Marshal(response)

			if err != nil {
				return errors.Wrap(err, "response parse error: ")
			}

			// print result in verbose mode
			if viper.GetBool("verbose") {
				cui.Outputln(string(responseJSON))
			}

			return nil
		},
	}

	graphCreateCmd.Flags().StringP("timezone", "", "", "timezone")
	graphCreateCmd.Flags().StringP("selfSufficient", "", "none", "selfSufficient")

	return graphCreateCmd
}

func newGraphUpdateCmd() *cobra.Command {
	graphUpdateCmd := &cobra.Command{
		Use:   "update",
		Short: "update graph definition",
		Long: `update graph definition. Usage:

$ pixela graph update <graph id> [--name graph_name] [--unit graph_unit] [--color color_name] [--timezone timezone] [--purgeCacheURLs [url1, url2, ...]]

see official document (https://docs.pixe.la/#/put-graph) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			// TODO: add timezone option later
			if len(args) != 1 {
				return errors.New(fmt.Sprintf("argument error: `graph update` requires 1 arguments give %d arguments.", len(args)))
			}

			// do request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			name, _ := cmd.Flags().GetString("name")
			unit, _ := cmd.Flags().GetString("unit")
			color, _ := cmd.Flags().GetString("color")
			timezone, _ := cmd.Flags().GetString("timezone")
			purgeUrls, err := cmd.Flags().GetStringArray("purge")

			if err != nil {
				return err
			}

			pl := pixela.UpdateGraphPayload{
				Name:           name,
				Unit:           unit,
				Color:          color,
				Timezone:       timezone,
				PurgeCacheURLs: purgeUrls,
			}

			response, err := client.UpdateGraph(args[0], pl)

			if err != nil {
				return errors.Wrap(err, "request error: ")
			}

			responseJSON, err := json.Marshal(response)

			if err != nil {
				return errors.Wrap(err, "response parse error: ")
			}

			// print result in verbose mode
			if viper.GetBool("verbose") {
				cui.Outputln(string(responseJSON))
			}

			return nil
		},
	}

	graphUpdateCmd.Flags().StringP("name", "n", "", "graph name")
	graphUpdateCmd.Flags().StringP("unit", "u", "", "graph unit")
	graphUpdateCmd.Flags().StringP("color", "c", "", "graph color (shibafu/momiji/sora/ichou/ajisai/kuro)")
	graphUpdateCmd.Flags().StringP("timezone", "t", "", "graph timezone")
	graphUpdateCmd.Flags().StringArrayP("purge", "p", nil, "purge cache urls")

	return graphUpdateCmd
}

func newGraphDeleteCmd() *cobra.Command {
	graphDeleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete graph",
		Long: `delete graph. Usage:

$ pixela graph delete <graph id>

see official document (https://docs.pixe.la/#/delete-graph) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			if len(args) != 1 {
				return errors.New(fmt.Sprintf("argument error: `graph delete` requires 1 arguments give %d arguments.", len(args)))
			}

			// do request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			response, err := client.DeleteGraph(args[0])

			if err != nil {
				return errors.Wrap(err, "request error: ")
			}

			responseJSON, err := json.Marshal(response)

			if err != nil {
				return errors.Wrap(err, "response parse error: ")
			}

			// print result in verbose mode
			if viper.GetBool("verbose") {
				cui.Outputln(string(responseJSON))
			}

			return nil
		},
	}

	return graphDeleteCmd
}

func newGraphDefCmd() *cobra.Command {
	graphDefinitionCmd := &cobra.Command{
		Use:   "def",
		Short: "get graph definitions",
		Long: `get graph definitions. Usage:

$ pixela graph def

see official document (https://docs.pixe.la/#/get-graph) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			if len(args) != 0 {
				return errors.New("argument error: `graph def` does not accept any argument.")
			}

			// do request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			response, err := client.GetGraphDefinition()

			if err != nil {
				return errors.Wrap(err, "request error: ")
			}

			responseJSON, err := json.Marshal(response)

			if err != nil {
				return errors.Wrap(err, "response parse error: ")
			}

			// print result
			cui.Outputln(string(responseJSON))

			return nil
		},
	}

	return graphDefinitionCmd
}

func newGraphSvgCmd() *cobra.Command {
	graphSvgCmd := &cobra.Command{
		Use:   "svg",
		Short: "get graph SVG HTML tag",
		Long: `get graph SVG HTML tag. Usage:

$ pixela graph svg <graph id>

see official document (https://docs.pixe.la/#/get-svg) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			if len(args) != 1 {
				return errors.New(fmt.Sprintf("argument error: `graph svg` requires 1 arguments give %d arguments.", len(args)))
			}

			// do request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			dateStr, _ := cmd.Flags().GetString("date")
			mode, _ := cmd.Flags().GetString("mode")

			response, err := client.GetGraphSvg(args[0], dateStr, mode)

			if err != nil {
				return errors.Wrap(err, "request error: ")
			}

			// print result
			cui.Outputln(string(response))

			return nil
		},
	}

	graphSvgCmd.Flags().StringP("date", "", "", "date")
	graphSvgCmd.Flags().StringP("mode", "", "", "mode")

	return graphSvgCmd
}

func newGraphPixelsDateCmd() *cobra.Command {
	graphPixelsDateCmd := &cobra.Command{
		Use:   "pixels",
		Short: "get graph pixels list",
		Long: `get graph pixels list. Usage:

$ pixela graph pixels <graph id> [--from yyyyMMdd] [--to yyyyMMdd]

see official document (https://docs.pixe.la/#/get-graph-pixels) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			if len(args) != 1 {
				return errors.New(fmt.Sprintf("argument error: `graph pixels` requires 1 argument give %d arguments.", len(args)))
			}

			// do request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			from, _ := cmd.Flags().GetString("from")
			to, _ := cmd.Flags().GetString("to")

			response, err := client.GetGraphPixelsDateList(args[0], from, to)

			if err != nil {
				return errors.Wrap(err, "request error: ")
			}

			responseJSON, err := json.Marshal(response)

			if err != nil {
				return errors.Wrap(err, "response parse error: ")
			}

			// print result
			cui.Outputln(string(responseJSON))

			return nil
		},
	}

	graphPixelsDateCmd.Flags().StringP("from", "", "", "from")
	graphPixelsDateCmd.Flags().StringP("to", "", "", "to")

	return graphPixelsDateCmd
}
