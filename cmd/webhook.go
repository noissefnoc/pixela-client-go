package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newWebhookCmd() *cobra.Command {
	webhookCmd := &cobra.Command{
		Use:   "webhook",
		Short: "handle webhook subcommands (create, get, invoke and delete)",
		Long: `create, get invoke and delete pixe.la webhook.
see official document (https://docs.pixe.la) for more detail`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Help()
			return nil
		},
	}

	webhookCmd.AddCommand(newWebhookCreateCmd())
	webhookCmd.AddCommand(newWebhookGetCmd())
	webhookCmd.AddCommand(newWebhookInvokeCmd())
	webhookCmd.AddCommand(newWebhookDeleteCmd())

	return webhookCmd
}

func newWebhookCreateCmd() *cobra.Command {
	webhookCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "create webhook",
		Long: `create webhook. Usage:

$ pixela webhook create <graph id> <type>

see official document (https://docs.pixe.la/#/post-webhook) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			// TODO: add timezone option later
			if len(args) != 2 {
				return errors.New(fmt.Sprintf("argument error: `webhook create` requires 2 arguments give %d arguments.", len(args)))
			}

			// do request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			response, err := client.CreateWebhook(args[0], args[1])

			if err != nil {
				return errors.Wrap(err, "request error: ")
			}

			responseJSON, err := json.Marshal(response)

			if err != nil {
				return errors.Wrap(err, "response parse error: ")
			}

			// print result in verbose mode
			if viper.GetBool("verbose") {
				cui.Outputln(responseJSON)
			}

			return nil
		},
	}

	return webhookCreateCmd
}

func newWebhookGetCmd() *cobra.Command {
	webhookGetCmd := &cobra.Command{
		Use:   "get",
		Short: "get user's webhook definitions",
		Long: `get user's webhook definitions'. Usage:

$ pixela webhook get

see official document (https://docs.pixe.la/#/get-webhook) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			// TODO: add timezone option later
			if len(args) != 0 {
				return errors.New(fmt.Sprintf("argument error: `webhook get` requires 0 arguments give %d arguments.", len(args)))
			}

			// do request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			response, err := client.GetWebhookDefinitions()

			if err != nil {
				return errors.Wrap(err, "request error: ")
			}

			responseJSON, err := json.Marshal(response)

			if err != nil {
				return errors.Wrap(err, "response parse error: ")
			}

			// print result
			cui.Outputln(responseJSON)

			return nil
		},
	}

	return webhookGetCmd
}

func newWebhookInvokeCmd() *cobra.Command {
	webhookInvokeCmd := &cobra.Command{
		Use:   "invoke",
		Short: "invoke webhook registered in advance",
		Long: `invoke webhook registered in advance. Usage:

$ pixela webhook invoke <webhook hash>

see official document (https://docs.pixe.la/#/invoke-webhook) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			// TODO: add timezone option later
			if len(args) != 1 {
				return errors.New(fmt.Sprintf("argument error: `webhook create` requires 1 arguments give %d arguments.", len(args)))
			}

			// do request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			response, err := client.InvokeWebhooks(args[0])

			if err != nil {
				return errors.Wrap(err, "request error: ")
			}

			responseJSON, err := json.Marshal(response)

			if err != nil {
				return errors.Wrap(err, "response parse error:")
			}

			// print result in verbose mode
			if viper.GetBool("verbose") {
				cui.Outputln(responseJSON)
			}

			return nil
		},
	}

	return webhookInvokeCmd
}

func newWebhookDeleteCmd() *cobra.Command {
	webhookDeleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete webhook",
		Long: `delete webhook. Usage:

$ pixela webhook delete <webhook hash>

see official document (https://docs.pixe.la/#/delete-webhook) for more detail.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// check arguments
			// TODO: add timezone option later
			if len(args) != 1 {
				return errors.New(fmt.Sprintf("argument error: `webhook delete` requires 1 arguments give %d arguments.", len(args)))
			}

			// do request
			client, err := pixela.New(viper.GetString("username"), viper.GetString("token"), viper.GetBool("verbose"))

			if err != nil {
				return err
			}

			response, err := client.DeleteWebhook(args[0])

			if err != nil {
				return errors.Wrap(err, "request error: ")
			}

			responseJSON, err := json.Marshal(response)

			if err != nil {
				return errors.Wrap(err, "response parse error: ")
			}

			// print result in verbose mode
			if viper.GetBool("verbose") {
				cui.Outputln(responseJSON)
			}

			return nil
		},
	}

	return webhookDeleteCmd
}
