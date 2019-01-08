package pixel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/noissefnoc/pixela-client-go/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// record quantity
func Record(cmd *cobra.Command, args []string) {
	// TODO: add OptionalData later
	type Payload struct {
		Date     string `json:"date"`
		Quantity string `json:"quantity"`
	}

	if len(args) != 3 {
		fmt.Fprintf(
			os.Stderr,
			`error "pixel record" requires 3 arguments: give %d arguments.\n`,
			len(args))
		os.Exit(1)
	}

	// TODO: arguments format validation
	graphId := args[0]
	date := args[1]
	quantity := args[2]

	username := viper.GetString("username")

	pl := Payload{
		Date:     date,
		Quantity: quantity,
	}

	plJSON, err := json.Marshal(pl)

	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			`error can not marshal request json.: %v\n`,
			err)
		os.Exit(1)
	}

	// TODO: extract url into one file
	// TODO: look for how to get global value in cobra (in case `username` and `graph id`)
	requestURL := fmt.Sprintf("https://pixe.la/v1/users/%s/graphs/%s", username, graphId)

	err = util.DoRequest("POST", requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"error pixel create request failed.: %v\n",
			err)
		os.Exit(1)
	}
}
