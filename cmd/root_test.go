package cmd

import (
	"bytes"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"testing"
)

func TestRootNormal(t *testing.T) {
	tests := []struct {
		args []string
		want string
	} {
		{args: []string{""}, want: ""},
	}

	for _, tt := range tests {
		out := new(bytes.Buffer)
		errOut := new(bytes.Buffer)

		ui := rwi.New(
			rwi.WithErrorWriter(out),
			rwi.WithErrorWriter(errOut),
		)

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
	}
}
