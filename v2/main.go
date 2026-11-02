package main

import (
	"context"
	"os"

	"github.com/noissefnoc/pixela-client-go/v2/cmd"
)

var (
	// Version is build version
	Version string
	// Revision is build revision
	Revision string
)

func main() {
	if Version != "" {
		cmd.Version = Version
	}
	if Revision != "" {
		cmd.Revision = Revision
	}
	os.Exit(cmd.Run(context.Background(), os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
