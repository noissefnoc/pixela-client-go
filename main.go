package main

import (
	"os"

	"github.com/goark/gocli/rwi"

	"github.com/noissefnoc/pixela-client-go/cmd"
)

var (
	// Version is build version
	Version string
	// Revision is build revision
	Revision string
)

func main() {
	cmd.Execute(
		rwi.New(
			rwi.WithReader(os.Stdin),
			rwi.WithWriter(os.Stdout),
			rwi.WithErrorWriter(os.Stderr),
		),
		os.Args[1:],
	).Exit()
}
