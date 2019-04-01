package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

// CreatePixelPayload is payload for `pixel create` subcommand
type CreatePixelPayload struct {
	Date         string `json:"date"`
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData,omitempty"`
}

// GetPixelResponseBody is response for `pixel get` subcommand
type GetPixelResponseBody struct {
	Quantity     string `json:"quantity"`
	OptionalData string `json:"optionalData,omitempty"`
}

// PostPixel is method for `pixel post` subcommand
func (pixela *Pixela) PostPixel(graphID, date, quantity, optionalData string) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphID:      graphID,
		Date:         date,
		Quantity:     quantity,
		OptionalData: optionalData,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel post`: wrong arguments")
	}

	// create payload
	pl := CreatePixelPayload{
		Date:     date,
		Quantity: quantity,
	}

	// set optionalData
	if len(optionalData) != 0 {
		pl.OptionalData = optionalData
	}

	plJSON, err := json.Marshal(pl)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel post`: can not marshal request payload")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s", baseURL, pixela.Username, graphID)

	// do request
	responseBody, err := pixela.post(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel post`: http request failed")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel post`: response parse failed")
	}

	return postResponseBody, nil
}

// GetPixel is method for `pixel get` subcommand
func (pixela *Pixela) GetPixel(graphID string, date string) (GetPixelResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphID: graphID,
		Date:    date,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return GetPixelResponseBody{}, errors.Wrap(err, "`pixel get`: wrong arguments")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s/%s", baseURL, pixela.Username, graphID, date)

	// do request
	responseBody, err := pixela.get(requestURL)

	if err != nil {
		return GetPixelResponseBody{}, errors.Wrap(err, "`pixel get`: http request failed")
	}

	getPixelResponseBody := GetPixelResponseBody{}
	err = json.Unmarshal(responseBody, &getPixelResponseBody)

	if err != nil {
		return GetPixelResponseBody{}, errors.Wrap(err, "`pixel get`: http response parse failed")
	}

	return getPixelResponseBody, nil
}

// UpdatePixel is method for `pixel update` subcommand
func (pixela *Pixela) UpdatePixel(graphID, date, quantity, optionalData string) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphID:      graphID,
		Date:         date,
		Quantity:     quantity,
		OptionalData: optionalData,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel update`: wrong arguments")
	}

	// create payload
	pl := CreatePixelPayload{
		Quantity: quantity,
	}

	// set optionalData
	if len(optionalData) != 0 {
		pl.OptionalData = optionalData
	}

	plJSON, err := json.Marshal(pl)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel update`: can not marshal request payload")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s/%s", baseURL, pixela.Username, graphID, date)

	// do request
	responseBody, err := pixela.put(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel update`: http request failed")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel update`: response parse failed")
	}

	return postResponseBody, nil
}

// IncrementPixel is method for `pixel increment` subcommand
func (pixela *Pixela) IncrementPixel(graphID string) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphID: graphID,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel increment`: wrong arguments")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s/increment", baseURL, pixela.Username, graphID)

	// do request
	responseBody, err := pixela.put(requestURL, nil)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel increment`: http request failed")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel increment`: response parse failed")
	}

	return postResponseBody, nil
}

// DecrementPixel is method for `pixel decrement` subcommand
func (pixela *Pixela) DecrementPixel(graphID string) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphID: graphID,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel decrement`: wrong arguments")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s/decrement", baseURL, pixela.Username, graphID)

	// do request
	responseBody, err := pixela.put(requestURL, nil)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel decrement`: http request failed")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel decrement`: response parse failed")
	}

	return postResponseBody, nil
}

// DeletePixel is method for `pixel delete` subcommand
func (pixela *Pixela) DeletePixel(graphID, date string) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphID: graphID,
		Date:    date,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel delete`: wrong arguments")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s/%s", baseURL, pixela.Username, graphID, date)

	// do request
	responseBody, err := pixela.delete(requestURL)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel delete`: http request failed")
	}

	deleteResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &deleteResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel delete`: response parse failed")
	}

	return deleteResponseBody, nil
}
