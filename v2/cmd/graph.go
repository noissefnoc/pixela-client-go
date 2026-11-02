package cmd

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"

	"github.com/noissefnoc/pixela-client-go/v2/pixela"
)

const graphUsage = `create, def(get definitions), svg, update and delete graph.
see official document (https://docs.pixe.la) for more detail

Usage:
  pixela graph <command> [arguments...]

Commands:
  create    create pixe.la graph
  update    update graph definition
  delete    delete graph
  get       get graph definitions
  svg       get graph SVG HTML tag
  pixels    get graph pixels list
  detail    get graph detail URL
  stat      get graph stat
  stopwatch stop graph stopwatch
`

func runGraph(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(out, graphUsage)
		return 0
	}

	switch args[0] {
	case "create":
		return runGraphCreate(ctx, args[1:], cfg, in, out, err)
	case "update":
		return runGraphUpdate(ctx, args[1:], cfg, in, out, err)
	case "delete":
		return runGraphDelete(ctx, args[1:], cfg, in, out, err)
	case "get", "def":
		return runGraphGet(ctx, args[1:], cfg, in, out, err)
	case "svg":
		return runGraphSvg(ctx, args[1:], cfg, in, out, err)
	case "pixels":
		return runGraphPixels(ctx, args[1:], cfg, in, out, err)
	case "detail":
		return runGraphDetail(ctx, args[1:], cfg, in, out, err)
	case "stat":
		return runGraphStat(ctx, args[1:], cfg, in, out, err)
	case "stopwatch":
		return runGraphStopwatch(ctx, args[1:], cfg, in, out, err)
	case "help", "-h", "--help":
		fmt.Fprint(out, graphUsage)
		return 0
	default:
		fmt.Fprintf(err, "unknown graph command: %q\nRun 'pixela graph help' for usage.\n", args[0])
		return 1
	}
}

func runGraphCreate(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela graph create", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	var timezone string
	fs.StringVar(&timezone, "timezone", "", "timezone")

	var selfSufficient string
	fs.StringVar(&selfSufficient, "selfSufficient", "none", "selfSufficient (none/increment/decrement)")

	fs.Usage = func() {
		fmt.Fprintln(err, `create pixe.la graph. Usage:

  pixela graph create <graph id> <graph name> <unit> <type> <color> [--timezone timezone] [--selfSufficient selfSufficient]

see official document (https://docs.pixe.la/#/post-graph) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 5 {
		fmt.Fprintf(err, "argument error: `graph create` requires 5 arguments, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	params := pixela.CreateGraphParams{
		ID:             remArgs[0],
		Name:           remArgs[1],
		Unit:           remArgs[2],
		Type:           remArgs[3],
		Color:          remArgs[4],
		Timezone:       timezone,
		SelfSufficient: selfSufficient,
	}

	resp, reqErr := client.CreateGraph(ctx, params)
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

func runGraphUpdate(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela graph update", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	var name, unit, color, timezone string
	var purgeUrls stringSliceFlag

	fs.StringVar(&name, "name", "", "graph name")
	fs.StringVar(&unit, "unit", "", "graph unit")
	fs.StringVar(&color, "color", "", "graph color (shibafu/momiji/sora/ichou/ajisai/kuro)")
	fs.StringVar(&timezone, "timezone", "", "graph timezone")
	fs.Var(&purgeUrls, "purge", "purge cache URLs")

	fs.Usage = func() {
		fmt.Fprintln(err, `update graph definition. Usage:

  pixela graph update <graph id> [--name graph_name] [--unit graph_unit] [--color color_name] [--timezone timezone] [--purge url]

see official document (https://docs.pixe.la/#/put-graph) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 1 {
		fmt.Fprintf(err, "argument error: `graph update` requires 1 argument, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	pl := pixela.UpdateGraphParams{
		Name:           name,
		Unit:           unit,
		Color:          color,
		Timezone:       timezone,
		PurgeCacheURLs: purgeUrls,
	}

	resp, reqErr := client.UpdateGraph(ctx, remArgs[0], pl)
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

func runGraphDelete(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela graph delete", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	fs.Usage = func() {
		fmt.Fprintln(err, `delete graph. Usage:

  pixela graph delete <graph id>

see official document (https://docs.pixe.la/#/delete-graph) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 1 {
		fmt.Fprintf(err, "argument error: `graph delete` requires 1 argument, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	resp, reqErr := client.DeleteGraph(ctx, remArgs[0])
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

func runGraphGet(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela graph get", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	fs.Usage = func() {
		fmt.Fprintln(err, `get graph definitions. Usage:

  pixela graph get

see official document (https://docs.pixe.la/#/get-graph) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 0 {
		fmt.Fprintln(err, "argument error: `graph get` does not accept any arguments")
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	resp, reqErr := client.GetGraphDefinitions(ctx)
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

func runGraphSvg(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela graph svg", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	var date, mode string
	fs.StringVar(&date, "date", "", "date (yyyyMMdd)")
	fs.StringVar(&mode, "mode", "", "mode (short/line)")

	fs.Usage = func() {
		fmt.Fprintln(err, `get graph SVG HTML tag. Usage:

  pixela graph svg <graph id> [--date yyyyMMdd] [--mode short/line]

see official document (https://docs.pixe.la/#/get-svg) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 1 {
		fmt.Fprintf(err, "argument error: `graph svg` requires 1 argument, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	svgData, reqErr := client.GetGraph(ctx, remArgs[0], pixela.GetGraphOptions{Date: date, Mode: mode})
	if reqErr != nil {
		fmt.Fprintf(err, "request error: %v\n", reqErr)
		return 1
	}

	fmt.Fprintln(out, string(svgData))
	return 0
}

func runGraphPixels(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela graph pixels", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	var from, to string
	fs.StringVar(&from, "from", "", "from date (yyyyMMdd)")
	fs.StringVar(&to, "to", "", "to date (yyyyMMdd)")

	fs.Usage = func() {
		fmt.Fprintln(err, `get graph pixels list. Usage:

  pixela graph pixels <graph id> [--from yyyyMMdd] [--to yyyyMMdd]

see official document (https://docs.pixe.la/#/get-graph-pixels) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 1 {
		fmt.Fprintf(err, "argument error: `graph pixels` requires 1 argument, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	resp, reqErr := client.GetPixelDates(ctx, remArgs[0], from, to)
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

func runGraphDetail(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela graph detail", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	fs.Usage = func() {
		fmt.Fprintln(err, `get graph detail URL. Usage:

  pixela graph detail <graph id>

see official document (https://docs.pixe.la/#/get-graph-html) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 1 {
		fmt.Fprintf(err, "argument error: `graph detail` requires 1 argument, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	urlStr := fmt.Sprintf("%s/v1/users/%s/graphs/%s.html", client.BaseURL(), client.Username(), remArgs[0])
	fmt.Fprintln(out, urlStr)
	return 0
}

func runGraphStat(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela graph stat", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	fs.Usage = func() {
		fmt.Fprintln(err, `get graph stat (such as total pixels count). Usage:

  pixela graph stat <graph id>

see official document (https://docs.pixe.la/#/get-graph-stat) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 1 {
		fmt.Fprintf(err, "argument error: `graph stat` requires 1 argument, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	resp, reqErr := client.GetGraphStats(ctx, remArgs[0])
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

func runGraphStopwatch(ctx context.Context, args []string, cfg *Config, in io.Reader, out, err io.Writer) int {
	fs := flag.NewFlagSet("pixela graph stopwatch", flag.ContinueOnError)
	fs.SetOutput(err)
	addGlobalFlags(fs, cfg)

	fs.Usage = func() {
		fmt.Fprintln(err, `stop graph stopwatch. Usage:

  pixela graph stopwatch <graph id>

see official document (https://docs.pixe.la/#/get-stopwatch) for more detail.`)
	}

	remArgs, parseErr := parseFlagSet(fs, args)
	if parseErr != nil {
		return 1
	}

	if len(remArgs) != 1 {
		fmt.Fprintf(err, "argument error: `graph stopwatch` requires 1 argument, got %d arguments\n", len(remArgs))
		fs.Usage()
		return 1
	}

	client, clientErr := newClient(cfg)
	if clientErr != nil {
		fmt.Fprintln(err, clientErr)
		return 1
	}

	resp, reqErr := client.StopGraphStopwatch(ctx, remArgs[0])
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
