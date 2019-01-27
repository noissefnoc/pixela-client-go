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

// record quantity (insert)
func (pixela *Pixela) CreatePixel(graphId, date, quantity, optionalData string) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphId: graphId,
		Date: date,
		Quantity: quantity,
		OptionalData: optionalData,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel create`: wrong arguments")
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
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel create`: can not marshal request payload.")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s", baseUrl, pixela.Username, graphId)

	// do request
	responseBody, err := pixela.post(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel create`: http request failed.")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	return postResponseBody, nil
}

// get pixel data
func (pixela *Pixela) GetPixel(graphId string, date string) (GetPixelResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphId: graphId,
		Date: date,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return GetPixelResponseBody{}, errors.Wrap(err, "`pixel get`: wrong arguments")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s/%s", baseUrl, pixela.Username, graphId, date)

	// do request
	responseBody, err := pixela.get(requestURL)

	if err != nil {
		return GetPixelResponseBody{}, errors.Wrap(err, "`pixel get`: http request failed.")
	}

	getPixelResponseBody := GetPixelResponseBody{}
	err = json.Unmarshal(responseBody, &getPixelResponseBody)

	if err != nil {
		return GetPixelResponseBody{}, errors.Wrap(err, "`pixel get`: http response parse failed.")
	}

	return getPixelResponseBody, nil
}

// record quantity (upsert)
func (pixela *Pixela) UpdatePixel(graphId, date, quantity, optionalData string) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphId: graphId,
		Date: date,
		Quantity: quantity,
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
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel update`: can not marshal request payload.")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s/%s", baseUrl, pixela.Username, graphId, date)

	// do request
	responseBody, err := pixela.put(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel update`: http request failed.")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	return postResponseBody, nil
}

// increment today's pixel quantity
func (pixela *Pixela) IncPixel(graphId string) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphId: graphId,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel inc`: wrong arguments")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s/increment", baseUrl, pixela.Username, graphId)

	// do request
	responseBody, err := pixela.put(requestURL, nil)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel inc`: http request failed.")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	return postResponseBody, nil
}

// decrement today's pixel quantity
func (pixela *Pixela) DecPixel(graphId string) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphId: graphId,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel dec`: wrong arguments")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s/decrement", baseUrl, pixela.Username, graphId)

	// do request
	responseBody, err := pixela.put(requestURL, nil)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel dec`: http request failed.")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	return postResponseBody, nil
}

// delete pixel
func (pixela *Pixela) DeletePixel(graphId, date string) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphId: graphId,
		Date: date,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel delete`: wrong arguments")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s/%s", baseUrl, pixela.Username, graphId, date)

	// do request
	responseBody, err := pixela.delete(requestURL)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`pixel delete`: http request failed.")
	}

	deleteResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &deleteResponseBody)

	return deleteResponseBody, nil
}
