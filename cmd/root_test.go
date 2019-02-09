package cmd

import (
	"bytes"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"testing"
)

func TestRootCmd(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"without option", []string{""}, ""},
		{"verbose option", []string{"--verbose"}, ""},
	}

	for _, tt := range tests {
		out := new(bytes.Buffer)
		errOut := new(bytes.Buffer)

		ui := rwi.New(
			rwi.WithErrorWriter(out),
			rwi.WithErrorWriter(errOut),
		)

		t.Run(tt.name, func(t *testing.T) {
			exit := Execute(ui, tt.args)

			if exit != exitcode.Normal {
				t.Errorf("Execute() err = \"%v\", want \"%v\".", exit, exitcode.Normal)
			}

			if out.String() != tt.want {
				t.Errorf("Execute() Stdout = \"%v\", want \"%v\".", out.String(), tt.want)
			}

			if errOut.String() != "" {
				t.Errorf("Execute() Stderr = \"%v\", want \"%v\".", errOut.String(), "")
			}
		})

	}
}
