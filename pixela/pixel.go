package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

// TODO: add OptionalData later
type RecordPayload struct {
	Date     string `json:"date"`
	Quantity string `json:"quantity"`
}

// record quantity
func (pixela *Pixela) PixelRecord(graphId string, date string, quantity string) error {
	// create payload
	pl := RecordPayload{
		Date:     date,
		Quantity: quantity,
	}

	plJSON, err := json.Marshal(pl)

	if err != nil {
		return errors.Wrap(err, "error can not marshal request payload.")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s", BaseUrl, pixela.Username, graphId)

	// do request
	err = pixela.post(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return errors.Wrap(err, "error pixel create http request failed.")
	}

	// TODO: check response body `isSuccess` field is `true`

	return nil
}
