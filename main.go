package main

import "github.com/noissefnoc/pixela-client-go/cmd"

var (
	// Version is build version
	Version string
	// Revision is build revision
	Revision string
)

func main() {
	cmd.Execute()
}
