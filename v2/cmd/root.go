package cmd

import (
	"context"
	"flag"
	"fmt"
	"io"
	"strings"

	"github.com/noissefnoc/pixela-client-go/v2/pixela"
)

var (
	// Version is the build version of the CLI.
	Version = "v2.0.0-dev"
	// Revision is the git revision of the build.
	Revision = ""
)

const rootUsage = `Simple API Client for pixela (https://pixe.la/).
This command can handle user, graph, pixel and webhook via API.

Usage:
  pixela [flags] <command> [subcommand] [arguments...]

Commands:
  user      handle user subcommands (create, update, delete)
  graph     handle graph subcommands (create, update, delete, get, svg, pixels, detail, stat)
  pixel     handle pixel subcommands (create, get, update, delete, increment, decrement)
  webhook   handle webhook subcommands (create, get, invoke, delete)
  version   show version information
  help      show help

Global Flags:
  -c, --config string     config file (default is $HOME/.pixela.yaml)
  -u, --username string   pixe.la username
  -t, --token string      pixe.la user token
  -n, --verbose           verbose mode
  -h, --help              show help
  -v, --version           show version
`

func printVersion(out io.Writer) {
	if Revision != "" {
		fmt.Fprintf(out, "pixela version %s (rev: %s)\n", Version, Revision)
	} else {
		fmt.Fprintf(out, "pixela version %s\n", Version)
	}
}

func newClient(cfg *Config) (*pixela.Client, error) {
	if cfg.Username == "" || cfg.Token == "" {
		return nil, fmt.Errorf("authentication error: username and token are required (use --username and --token flags, environment variables PIXELA_USERNAME and PIXELA_TOKEN, or ~/.pixela.yaml)")
	}
	opts := []pixela.Option{pixela.WithDebug(cfg.Verbose)}
	if cfg.BaseURL != "" {
		opts = append(opts, pixela.WithBaseURL(cfg.BaseURL))
	}
	return pixela.New(cfg.Username, cfg.Token, opts...)
}

// Run is the main CLI entry point conforming to Kuniwak's recipe1.
func Run(ctx context.Context, args []string, in io.Reader, out, err io.Writer) int {
	// Look ahead for explicit --config or -c flag
	var explicitConfig string
	for i, arg := range args {
		if (arg == "--config" || arg == "-c") && i+1 < len(args) {
			explicitConfig = args[i+1]
			break
		} else if strings.HasPrefix(arg, "--config=") {
			explicitConfig = strings.TrimPrefix(arg, "--config=")
			break
		} else if strings.HasPrefix(arg, "-c=") {
			explicitConfig = strings.TrimPrefix(arg, "-c=")
			break
		}
	}

	cfg, configErr := resolveConfig(explicitConfig)
	if configErr != nil {
		fmt.Fprintln(err, configErr)
		return 1
	}

	rootFS := flag.NewFlagSet("pixela", flag.ContinueOnError)
	rootFS.SetOutput(err)
	addGlobalFlags(rootFS, cfg)

	var showHelp bool
	rootFS.BoolVar(&showHelp, "help", false, "show help")
	rootFS.BoolVar(&showHelp, "h", false, "show help (shorthand)")

	var showVer bool
	rootFS.BoolVar(&showVer, "version", false, "show version")
	rootFS.BoolVar(&showVer, "v", false, "show version (shorthand)")

	rootFS.Usage = func() {
		fmt.Fprint(err, rootUsage)
	}

	if parseErr := rootFS.Parse(args); parseErr != nil {
		return 1
	}

	if showVer {
		printVersion(out)
		return 0
	}

	if showHelp {
		fmt.Fprint(out, rootUsage)
		return 0
	}

	remArgs := rootFS.Args()
	if len(remArgs) == 0 {
		fmt.Fprint(out, rootUsage)
		return 0
	}

	subCmd := remArgs[0]
	subArgs := remArgs[1:]

	switch subCmd {
	case "user":
		return runUser(ctx, subArgs, cfg, in, out, err)
	case "graph":
		return runGraph(ctx, subArgs, cfg, in, out, err)
	case "pixel":
		return runPixel(ctx, subArgs, cfg, in, out, err)
	case "webhook":
		return runWebhook(ctx, subArgs, cfg, in, out, err)
	case "version":
		printVersion(out)
		return 0
	case "help":
		fmt.Fprint(out, rootUsage)
		return 0
	default:
		fmt.Fprintf(err, "unknown command: %q\nRun 'pixela help' for usage.\n", subCmd)
		return 1
	}
}
