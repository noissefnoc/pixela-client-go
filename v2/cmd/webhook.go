package cmd

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"github.com/noissefnoc/pixela-client-go/v2/pixela"
)

const webhookUsage = `create, get, invoke and delete pixe.la webhook.
see official document (https://docs.pixe.la) for more detail

Usage:
  pixela webhook <command> [arguments...]

Commands:
  create   create webhook
  get      get user's webhook definitions
  invoke   invoke webhook registered in advance
  delete   delete webhook
`

func runWebhook(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(out, webhookUsage)
		return 0
	}

	switch args[0] {
	case "create":
		return runWebhookCreate(ctx, args[1:], cfg, in, out, err)
	case "get":
		return runWebhookGet(ctx, args[1:], cfg, in, out, err)
	case "invoke":
		return runWebhookInvoke(ctx, args[1:], cfg, in, out, err)
	case "delete":
		return runWebhookDelete(ctx, args[1:], cfg, in, out, err)
	case "help", "-h", "--help":
		fmt.Fprint(out, webhookUsage)
		return 0
	default:
		fmt.Fprintf(err, "unknown webhook command: %q\nRun 'pixela webhook help' for usage.\n", args[0])
		return 1
	}
}

func runWebhookCreate(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela webhook create", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	fs.Usage = func() {
		fmt.Fprintln(err, `create webhook. Usage:

  pixela webhook create <graph id> <type>

see official document (https://docs.pixe.la/#/post-webhook) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 2 {
		fmt.Fprintf(err, "argument error: `webhook create` requires 2 arguments, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	params := pixela.CreateWebhookParams{
		GraphID: remArgs[0],
		Type:    remArgs[1],
	}

	resp, reqErr := client.CreateWebhook(ctx, params)
	if reqErr != nil {
		fmt.Fprintf(err, "request error: %v\n", reqErr)
		return 1
	}

	respJSON, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		fmt.Fprintf(err, "response parse error: %v\n", jsonErr)
		return 1
	}

	if cfg.Verbose {
		fmt.Fprintln(out, string(respJSON))
	}

	return 0
}

func runWebhookGet(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela webhook get", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	fs.Usage = func() {
		fmt.Fprintln(err, `get user's webhook definitions. Usage:

  pixela webhook get

see official document (https://docs.pixe.la/#/get-webhook) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 0 {
		fmt.Fprintln(err, "argument error: `webhook get` does not accept any arguments")
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	resp, reqErr := client.GetWebhooks(ctx)
	if reqErr != nil {
		fmt.Fprintf(err, "request error: %v\n", reqErr)
		return 1
	}

	respJSON, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		fmt.Fprintf(err, "response parse error: %v\n", jsonErr)
		return 1
	}

	fmt.Fprintln(out, string(respJSON))
	return 0
}

func runWebhookInvoke(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela webhook invoke", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	fs.Usage = func() {
		fmt.Fprintln(err, `invoke webhook registered in advance. Usage:

  pixela webhook invoke <webhook hash>

see official document (https://docs.pixe.la/#/invoke-webhook) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 1 {
		fmt.Fprintf(err, "argument error: `webhook invoke` requires 1 argument, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	resp, reqErr := client.InvokeWebhook(ctx, remArgs[0])
	if reqErr != nil {
		fmt.Fprintf(err, "request error: %v\n", reqErr)
		return 1
	}

	respJSON, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		fmt.Fprintf(err, "response parse error: %v\n", jsonErr)
		return 1
	}

	if cfg.Verbose {
		fmt.Fprintln(out, string(respJSON))
	}

	return 0
}

func runWebhookDelete(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela webhook delete", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	fs.Usage = func() {
		fmt.Fprintln(err, `delete webhook. Usage:

  pixela webhook delete <webhook hash>

see official document (https://docs.pixe.la/#/delete-webhook) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 1 {
		fmt.Fprintf(err, "argument error: `webhook delete` requires 1 argument, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	resp, reqErr := client.DeleteWebhook(ctx, remArgs[0])
	if reqErr != nil {
		fmt.Fprintf(err, "request error: %v\n", reqErr)
		return 1
	}

	respJSON, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		fmt.Fprintf(err, "response parse error: %v\n", jsonErr)
		return 1
	}

	if cfg.Verbose {
		fmt.Fprintln(out, string(respJSON))
	}

	return 0
}
