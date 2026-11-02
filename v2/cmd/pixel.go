package cmd

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"github.com/noissefnoc/pixela-client-go/v2/pixela"
)

const pixelUsage = `record, get, update, increment, decrement and delete pixel.
see official document (https://docs.pixe.la) for more detail

Usage:
  pixela pixel <command> [arguments...]

Commands:
  create    create pixel
  get       get pixel value
  update    update pixel
  delete    delete pixel
  increment increment today's pixel quantity (alias: inc)
  decrement decrement today's pixel quantity (alias: dec)
`

func runPixel(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(out, pixelUsage)
		return 0
	}

	switch args[0] {
	case "create", "post":
		return runPixelCreate(ctx, args[1:], cfg, in, out, err)
	case "get":
		return runPixelGet(ctx, args[1:], cfg, in, out, err)
	case "update":
		return runPixelUpdate(ctx, args[1:], cfg, in, out, err)
	case "delete":
		return runPixelDelete(ctx, args[1:], cfg, in, out, err)
	case "increment", "inc":
		return runPixelIncrement(ctx, args[1:], cfg, in, out, err)
	case "decrement", "dec":
		return runPixelDecrement(ctx, args[1:], cfg, in, out, err)
	case "help", "-h", "--help":
		fmt.Fprint(out, pixelUsage)
		return 0
	default:
		fmt.Fprintf(err, "unknown pixel command: %q\nRun 'pixela pixel help' for usage.\n", args[0])
		return 1
	}
}

func runPixelCreate(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela pixel create", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	var optionalData string
	fs.StringVar(&optionalData, "optionalData", "", "optionalData (JSON format)")

	fs.Usage = func() {
		fmt.Fprintln(err, `create pixel. Usage:

  pixela pixel create <graph id> <date> <quantity> [--optionalData json]

see official document (https://docs.pixe.la/#/post-pixel) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 3 {
		fmt.Fprintf(err, "argument error: `pixel create` requires 3 arguments, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	params := pixela.CreatePixelParams{
		Date:         remArgs[1],
		Quantity:     remArgs[2],
		OptionalData: optionalData,
	}

	resp, reqErr := client.CreatePixel(ctx, remArgs[0], params)
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

func runPixelGet(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela pixel get", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	fs.Usage = func() {
		fmt.Fprintln(err, `get pixel value. Usage:

  pixela pixel get <graph id> <date>

see official document (https://docs.pixe.la/#/get-pixel) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 2 {
		fmt.Fprintf(err, "argument error: `pixel get` requires 2 arguments, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	resp, reqErr := client.GetPixel(ctx, remArgs[0], remArgs[1])
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

func runPixelUpdate(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela pixel update", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	var optionalData string
	fs.StringVar(&optionalData, "optionalData", "", "optionalData (JSON format)")

	fs.Usage = func() {
		fmt.Fprintln(err, `update pixel. Usage:

  pixela pixel update <graph id> <date> <quantity> [--optionalData json]

see official document (https://docs.pixe.la/#/put-pixel) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 3 {
		fmt.Fprintf(err, "argument error: `pixel update` requires 3 arguments, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	params := pixela.UpdatePixelParams{
		Quantity:     remArgs[2],
		OptionalData: optionalData,
	}

	resp, reqErr := client.UpdatePixel(ctx, remArgs[0], remArgs[1], params)
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

func runPixelDelete(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela pixel delete", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	fs.Usage = func() {
		fmt.Fprintln(err, `delete pixel. Usage:

  pixela pixel delete <graph id> <date>

see official document (https://docs.pixe.la/#/delete-pixel) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 2 {
		fmt.Fprintf(err, "argument error: `pixel delete` requires 2 arguments, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	resp, reqErr := client.DeletePixel(ctx, remArgs[0], remArgs[1])
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

func runPixelIncrement(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela pixel increment", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	fs.Usage = func() {
		fmt.Fprintln(err, `increment today's pixel quantity. Usage:

  pixela pixel increment <graph id>

see official document (https://docs.pixe.la/#/increment-pixel) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 1 {
		fmt.Fprintf(err, "argument error: `pixel increment` requires 1 argument, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	resp, reqErr := client.IncrementPixel(ctx, remArgs[0])
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

func runPixelDecrement(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela pixel decrement", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	fs.Usage = func() {
		fmt.Fprintln(err, `decrement today's pixel quantity. Usage:

  pixela pixel decrement <graph id>

see official document (https://docs.pixe.la/#/decrement-pixel) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 1 {
		fmt.Fprintf(err, "argument error: `pixel decrement` requires 1 argument, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	resp, reqErr := client.DecrementPixel(ctx, remArgs[0])
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
