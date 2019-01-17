package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type CreatePixelPayload struct {
	Date         string `json:"date"`
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData,omitempty"`
}

type GetPixelResponseBody struct {
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData,omitempty"`
}

// record quantity
func (pixela *Pixela) CreatePixel(graphId string, date string, quantity string) (NoneGetResponseBody, error) {
	// create payload
	pl := CreatePixelPayload{
		Date:     date,
		Quantity: quantity,
	}

	plJSON, err := json.Marshal(pl)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "error `pixel create`: can not marshal request payload.")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s", baseUrl, pixela.Username, graphId)

	// do request
	responseBody, err := pixela.post(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "error `pixel create`:http request failed.")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	return postResponseBody, nil
}

// get pixel data
func (pixela *Pixela) GetPixel(graphId string, date string) (GetPixelResponseBody, error) {
	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s/%s", baseUrl, pixela.Username, graphId, date)

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

// record quantity
func (pixela *Pixela) UpdatePixel(graphId string, date string, quantity string) (NoneGetResponseBody, error) {
	// create payload
	pl := CreatePixelPayload{
		Quantity: quantity,
	}

	plJSON, err := json.Marshal(pl)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "error `pixel update`: can not marshal request payload.")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s/%s", baseUrl, pixela.Username, graphId, date)

	// do request
	responseBody, err := pixela.put(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "error `pixel update`:http request failed.")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	return postResponseBody, nil
}

// increment today's pixel quantity
func (pixela *Pixela) IncPixel(graphId string) (NoneGetResponseBody, error) {
	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s/increment", baseUrl, pixela.Username, graphId)

	// do request
	responseBody, err := pixela.put(requestURL, nil)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "error `pixel inc`:http request failed.")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	return postResponseBody, nil
}

// decrement today's pixel quantity
func (pixela *Pixela) DecPixel(graphId string) (NoneGetResponseBody, error) {
	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s/decrement", baseUrl, pixela.Username, graphId)

	// do request
	responseBody, err := pixela.put(requestURL, nil)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "error `pixel dec`:http request failed.")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	return postResponseBody, nil
}

// delete pixel
func (pixela *Pixela) DeletePixel(graphId, date string) (NoneGetResponseBody, error) {
	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s/%s", baseUrl, pixela.Username, graphId, date)

	// do request
	responseBody, err := pixela.delete(requestURL)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "error `pixel delete`:http request failed.")
	}

	deleteResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &deleteResponseBody)

	return deleteResponseBody, nil
}
