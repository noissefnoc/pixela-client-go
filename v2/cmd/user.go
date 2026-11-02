package cmd

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"github.com/noissefnoc/pixela-client-go/v2/pixela"
)

const userUsage = `create, update (token information) and delete pixe.la user.
see official document (https://docs.pixe.la) for more detail

Usage:
  pixela user <command> [arguments...]

Commands:
  create   create pixe.la user
  update   update user token
  delete   delete pixe.la user
`

func runUser(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(out, userUsage)
		return 0
	}

	switch args[0] {
	case "create":
		return runUserCreate(ctx, args[1:], cfg, in, out, err)
	case "update":
		return runUserUpdate(ctx, args[1:], cfg, in, out, err)
	case "delete":
		return runUserDelete(ctx, args[1:], cfg, in, out, err)
	case "help", "-h", "--help":
		fmt.Fprint(out, userUsage)
		return 0
	default:
		fmt.Fprintf(err, "unknown user command: %q\nRun 'pixela user help' for usage.\n", args[0])
		return 1
	}
}

func runUserCreate(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela user create", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	var agreeTermsOfService string
	fs.StringVar(&agreeTermsOfService, "agreeTermsOfService", "yes", "agree terms of service")

	var notMinor string
	fs.StringVar(&notMinor, "notMinor", "yes", "usage is not minor")

	fs.Usage = func() {
		fmt.Fprintln(err, `create pixe.la user. Usage:

  pixela user create <username> <token> [--agreeTermsOfService yes/no] [--notMinor yes/no]

see official document (https://docs.pixe.la/#/post-user) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 2 {
		fmt.Fprintf(err, "argument error: `user create` requires 2 arguments, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	username := remArgs[0]
	token := remArgs[1]

	opts := []pixela.Option{pixela.WithDebug(cfg.Verbose)}
	if cfg.BaseURL != "" {
		opts = append(opts, pixela.WithBaseURL(cfg.BaseURL))
	}
	client, clientErr := pixela.New(username, token, opts...)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	resp, reqErr := client.CreateUser(ctx, pixela.CreateUserParams{
		AgreeTermsOfService: agreeTermsOfService == "yes",
		NotMinor:            notMinor == "yes",
	})
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

	if saveErr := saveConfigFile(cfg.ConfigFile, username, token, out); saveErr != nil {
		fmt.Fprintf(err, "save config file error: %v\n", saveErr)
		return 1
	}

	return 0
}

func runUserUpdate(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela user update", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	fs.Usage = func() {
		fmt.Fprintln(err, `update pixe.la user token. Usage:

  pixela user update <newToken>

see official document (https://docs.pixe.la/#/put-user) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 1 {
		fmt.Fprintf(err, "argument error: `user update` requires 1 argument, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	resp, reqErr := client.UpdateUser(ctx, pixela.UpdateUserParams{NewToken: remArgs[0]})
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

func runUserDelete(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela user delete", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	fs.Usage = func() {
		fmt.Fprintln(err, `delete pixe.la user. Usage:

  pixela user delete <username> <token>

see official document (https://docs.pixe.la/#/delete-user) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 2 {
		fmt.Fprintf(err, "argument error: `user delete` requires 2 arguments, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	opts := []pixela.Option{pixela.WithDebug(cfg.Verbose)}
	if cfg.BaseURL != "" {
		opts = append(opts, pixela.WithBaseURL(cfg.BaseURL))
	}
	client, clientErr := pixela.New(remArgs[0], remArgs[1], opts...)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	resp, reqErr := client.DeleteUser(ctx)
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
