package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var optionalData string

func newPixelCmd() *cobra.Command {
	pixelCmd := &cobra.Command{
		Use:   "pixel",
		Short: "handle pixel subcommands (create, get, update, increment, decrement and delete)",
		Long: `record, get, update, increment, decrement and delete pixel
see official document (https://docs.pixe.la) for more detail`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}

	pixelCmd.AddCommand(newPixelCreateCmd())
	pixelCmd.AddCommand(newPixelGetCmd())
	pixelCmd.AddCommand(newPixelUpdateCmd())
	pixelCmd.AddCommand(newPixelDeleteCmd())
	pixelCmd.AddCommand(newPixelIncCmd())
	pixelCmd.AddCommand(newPixelDecCmd())

	return pixelCmd
}

func newPixelCreateCmd() *cobra.Command {
	pixelCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "create pixel",
		Long: `create pixel. Usage:

$ pixela pixel create <graph id> <date> <quantity>

see official document (https://docs.pixe.la/#/post-pixel) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			if len(args) != 3 {
				return errors.New(fmt.Sprintf("argument error: `pixel create` requires 3 arguments give %d arguments.", len(args)))
			}

			// do request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			response, err := client.CreatePixel(args[0], args[1], args[2], optionalData)

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

	return pixelCreateCmd
}

func newPixelGetCmd() *cobra.Command {
	pixelGetCmd := &cobra.Command{
		Use:   "get",
		Short: "get pixel value",
		Long: `get pixel value. Usage:

$ pixela pixel get <graph id> <date>

see official document (https://docs.pixe.la/#/get-pixel) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			if len(args) != 2 {
				return errors.New(fmt.Sprintf("argument error: `pixel record` requires 2 arguments give %d arguments.", len(args)))
			}

			// do request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			response, err := client.GetPixel(args[0], args[1])

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

	return pixelGetCmd
}

func newPixelUpdateCmd() *cobra.Command {
	pixelUpdateCmd := &cobra.Command{
		Use:   "update",
		Short: "update pixel",
		Long: `update pixel. Usage:

$ pixela pixel update <graph id> <date> <quantity>

see official document (https://docs.pixe.la/#/put-pixel) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			if len(args) != 3 {
				return errors.New(fmt.Sprintf("argument error: `pixel update` requires 3 arguments give %d arguments.", len(args)))
			}

			// do request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			optionalData, err := cmd.Flags().GetString("optionalData")

			if err != nil {
				return err
			}

			response, err := client.UpdatePixel(args[0], args[1], args[2], optionalData)

			if err != nil {
				return errors.Wrap(err, "request error: ")
			}

			responseJSON, err := json.Marshal(response)

			if err != nil {
				return errors.Wrap(err, "response parse error: %v\n")
			}

			// print result in verbose mode
			if viper.GetBool("verbose") {
				cui.Outputln(string(responseJSON))
			}

			return nil
		},
	}

	pixelUpdateCmd.Flags().StringP("optionalData", "", "", "optionalData (JSON format)")

	return pixelUpdateCmd
}

func newPixelDeleteCmd() *cobra.Command {
	pixelDeleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete pixel",
		Long: `delete pixel. Usage:

$ pixela pixel delete <graph id> <date>

see official document (https://docs.pixe.la/#/delete-pixel) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			if len(args) != 2 {
				return errors.New(fmt.Sprintf("argument error: `pixel delete` requires 2 arguments give %d arguments.", len(args)))
			}

			// do request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			response, err := client.DeletePixel(args[0], args[1])

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

	return pixelDeleteCmd
}

func newPixelIncCmd() *cobra.Command {
	pixelIncCmd := &cobra.Command{
		Use:   "inc",
		Short: "increment pixel quantity",
		Long: `increment pixel quantity. Usage:

$ pixela pixel inc <graph id>

see official document (https://docs.pixe.la/#/increment-pixel) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			if len(args) != 1 {
				return errors.New(fmt.Sprintf("argument error: `pixel inc` requires 1 arguments give %d arguments.\n\n", len(args)))
			}

			// do request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			response, err := client.IncPixel(args[0])

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

	return pixelIncCmd
}

func newPixelDecCmd() *cobra.Command {
	pixelDecCmd := &cobra.Command{
		Use:   "dec",
		Short: "decrement today's pixel quantity",
		Long: `decrement today's pixel quantity. Usage:

$ pixela pixel dec <graph id>

see official document (https://docs.pixe.la/#/decrement-pixel) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			if len(args) != 1 {
				return errors.New(fmt.Sprintf("argument error: `pixel dec` requires 1 arguments give %d arguments.", len(args)))
			}

			// do request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			response, err := client.DecPixel(args[0])

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

	return pixelDecCmd
}
