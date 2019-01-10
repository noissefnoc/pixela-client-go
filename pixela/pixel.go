package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type RecordPayload struct {
	Date         string `json:"date"`
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData,omitempty"`
}

type GetPixelResponseBody struct {
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData,omitempty"`
}

// record quantity
func (pixela *Pixela) RecordPixel(graphId string, date string, quantity string) (PostResponseBody, error) {
	// create payload
	pl := RecordPayload{
		Date:     date,
		Quantity: quantity,
	}

	plJSON, err := json.Marshal(pl)

	if err != nil {
		return PostResponseBody{}, errors.Wrap(err, "error `pixel create`: can not marshal request payload.")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s", BaseUrl, pixela.Username, graphId)

	// do request
	responseBody, err := pixela.post(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return PostResponseBody{}, errors.Wrap(err, "error `pixel create`:http request failed.")
	}

	postResponseBody := PostResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	return postResponseBody, nil
}

func (pixela *Pixela) GetPixel(graphId string, date string) (GetPixelResponseBody, error) {
	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s/%s", BaseUrl, pixela.Username, graphId, date)

	// do request
	responseBody, err := pixela.get(requestURL)

	if err != nil {
		return GetPixelResponseBody{}, errors.Wrap(err, "error `pixel get`:http request failed.")
	}

	getPixelResponseBody := GetPixelResponseBody{}
	err = json.Unmarshal(responseBody, &getPixelResponseBody)

	if err != nil {
		return GetPixelResponseBody{}, errors.Wrap(err, "error `pixel get`:http response parse failed.")
	}

	return getPixelResponseBody, nil
}
