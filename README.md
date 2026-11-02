# pixela-client-go

[![Test](https://github.com/noissefnoc/pixela-client-go/actions/workflows/test.yml/badge.svg)](https://github.com/noissefnoc/pixela-client-go/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/noissefnoc/pixela-client-go)](https://goreportcard.com/report/github.com/noissefnoc/pixela-client-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

`pixela-client-go` is an unofficial Go API client library & CLI for [Pixela](https://pixe.la).

> [!NOTE]
> [pi](https://github.com/a-know/pi) is the official CLI.

---

## What's New in v2

`v2` is a complete redesign focused on reliability, security, and minimalism:

- **Zero External Dependencies**: Implemented entirely with Go standard libraries (`net/http`, `flag`, `encoding/json`, etc.). Zero third-party packages, zero supply-chain vulnerabilities.
- **Context-Aware**: Full `context.Context` support across all client SDK methods for timeout and cancellation control.
- **Lightweight Binary**: Reduced binary size by ~40% (11MB → 6.7MB) by removing heavy frameworks (`cobra`, `viper`, `validator`).
- **Testable Architecture**: Designed with [Kuniwak's recipe](https://github.com/Kuniwak) for highly testable CLI applications without frameworks.

---

## Installation

### CLI Installation

Install the latest `v2` binary using `go install`:

```bash
go install github.com/noissefnoc/pixela-client-go/v2@latest
```

Or download pre-built binaries from the [GitHub Releases page](https://github.com/noissefnoc/pixela-client-go/releases).

### Go SDK Installation

```bash
go get github.com/noissefnoc/pixela-client-go/v2
```

---

## Go SDK Usage (v2)

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/noissefnoc/pixela-client-go/v2/pixela"
)

func main() {
	ctx := context.Background()

	// 1. Initialize client
	client, err := pixela.New("your-username", "your-secret-token")
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	// 2. Create a graph
	_, err = client.CreateGraph(ctx, pixela.CreateGraphParams{
		ID:    "commit-graph",
		Name:  "Commit Counter",
		Unit:  "commit",
		Type:  "int",
		Color: "shibafu",
	})
	if err != nil {
		log.Printf("graph create error: %v", err)
	}

	// 3. Post a pixel
	result, err := client.CreatePixel(ctx, "commit-graph", pixela.CreatePixelParams{
		Date:     "20261102",
		Quantity: "5",
	})
	if err != nil {
		log.Fatalf("post pixel error: %v", err)
	}
	fmt.Printf("Pixel posted: %s\n", result.Message)

	// 4. Increment today's pixel
	_, err = client.IncrementPixel(ctx, "commit-graph")
	if err != nil {
		log.Fatalf("increment pixel error: %v", err)
	}

	// 5. Retrieve SVG image
	svgBytes, err := client.GetGraph(ctx, "commit-graph")
	if err != nil {
		log.Fatalf("get graph svg error: %v", err)
	}
	fmt.Printf("SVG length: %d bytes\n", len(svgBytes))
}
```

---

## CLI Usage (v2)

```
Usage:
  pixela [flags] <command> [subcommand] [arguments...]

Commands:
  user      handle user subcommands (create, update, delete)
  graph     handle graph subcommands (create, update, delete, get, svg, pixels, detail, stat, stopwatch)
  pixel     handle pixel subcommands (create, get, update, delete, increment, decrement)
  webhook   handle webhook subcommands (create, get, invoke, delete)
  version   show version information
  help      show help

Global Flags:
  -c, --config string     config file (default: $HOME/.pixela.yaml)
  -u, --username string   pixe.la username
  -t, --token string      pixe.la user token
  -n, --verbose           verbose mode
  -h, --help              show help
  -v, --version           show version
```

### Authentication & Configuration

The CLI resolves authentication in the following priority order:
1. Command-line flags (`-u, --username`, `-t, --token`)
2. Environment variables (`PIXELA_USERNAME`, `PIXELA_TOKEN`)
3. Configuration file (`~/.pixela.yaml` or `--config <path>`)

When you run `pixela user create <username> <token>`, the CLI automatically creates `~/.pixela.yaml` with your credentials:

```yaml
username: your-username
token: your-secret-token
```

### Quick Start Workflow

#### 1. Create User
```bash
pixela user create <username> <token> [--agreeTermsOfService yes] [--notMinor yes]
```

#### 2. Create Graph
```bash
pixela graph create <graph-id> <graph-name> <unit> <type> <color> [--timezone UTC]
# Example:
pixela graph create habits "Daily Habit" times int shibafu
```

- `type`: `int` or `float`
- `color`: `shibafu`, `momiji`, `sora`, `ichou`, `ajisai`, or `kuro`

#### 3. Record & Manage Pixels
```bash
# Record a pixel for a date
pixela pixel create habits 20261102 3

# Get pixel value
pixela pixel get habits 20261102

# Increment today's pixel
pixela pixel increment habits

# Decrement today's pixel
pixela pixel decrement habits

# Delete a pixel
pixela pixel delete habits 20261102
```

#### 4. Webhooks
```bash
# Create a webhook
pixela webhook create habits increment

# Invoke a webhook without user token
pixela webhook invoke <webhookHash>

# List all webhooks
pixela webhook get

# Delete a webhook
pixela webhook delete <webhookHash>
```

#### 5. Graphs
```bash
# Get SVG tag
pixela graph svg habits

# Get graph definitions
pixela graph get

# Get graph statistics
pixela graph stat habits

# Get graph detail URL
pixela graph detail habits

# Delete a graph
pixela graph delete habits
```

---

## Migration Guide: v1 to v2

| Feature | v1 | v2 |
|---|---|---|
| **Import Path** | `github.com/noissefnoc/pixela-client-go/pixela` | `github.com/noissefnoc/pixela-client-go/v2/pixela` |
| **Dependencies** | cobra, viper, validator/v10, gocli | **None (Go Standard Library only)** |
| **Context** | Not supported | `context.Context` required on all methods |
| **Method Signatures** | Positional string arguments | Strongly typed parameter structs (`CreateGraphParams`, `CreatePixelParams`, etc.) |
| **Binary Size** | ~11.1 MB | **~6.7 MB (~40% reduction)** |

---

## Development

This repository uses [mise](https://mise.jdx.dev) for tool and task management.

```bash
# Run all tests (v1 and v2)
mise run test

# Run v2 tests only
mise run test:v2

# Build binaries
mise run build

# Cross-compile for all platforms
mise run cross-build:v2

# Run linters
mise exec -- prek run --all-files
```

---

## License

This project is licensed under the [MIT License](LICENSE).

## Author

[noissefnoc](https://github.com/noissefnoc)
