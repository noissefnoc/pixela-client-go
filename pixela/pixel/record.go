package pixel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/noissefnoc/pixela-client-go/pixela"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// TODO: add OptionalData later
type RecordPayload struct {
	Date     string `json:"date"`
	Quantity string `json:"quantity"`
}

// record quantity
func Record(cmd *cobra.Command, args []string) {
	if len(args) != 3 {
		fmt.Fprintf(os.Stderr, `error "pixel record" requires 3 arguments: give %d arguments.\n`, len(args))
		os.Exit(1)
	}

	// TODO: arguments format validation
	graphId := args[0]
	date := args[1]
	quantity := args[2]

	pl := RecordPayload{
		Date:     date,
		Quantity: quantity,
	}

	plJSON, err := json.Marshal(pl)

	if err != nil {
		fmt.Fprintf(os.Stderr, `error can not marshal request json.: %v\n`, err)
		os.Exit(1)
	}

	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s", pixela.BaseUrl, viper.GetString("username"), graphId)

	client := pixela.Pixela{
		Url: requestURL,
		Token: viper.GetString("token"),
		Debug: true,
	}

	err = client.DoAPI("POST", requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		fmt.Fprintf(os.Stderr, "error pixel create request failed.: %v\n", err)
		os.Exit(1)
	}
}
