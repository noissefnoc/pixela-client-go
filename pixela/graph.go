package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/url"
	"path"
)

// CreateGraphPayload is payload for `graph create` subcommand
type CreateGraphPayload struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Unit           string `json:"unit"`
	NumType        string `json:"type"`
	Color          string `json:"color"`
	Timezone       string `json:"timezone,omitempty"`
	SelfSufficient string `json:"selfSufficient,omitempty"`
}

// GraphDefinitions is response for `graph def` subcommand
type GraphDefinitions struct {
	Graphs []Graph `json:"graphs"`
}

// Graph is part of response for `graph def` subcommand
type Graph struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Unit           string   `json:"unit"`
	Type           string   `json:"type"`
	Color          string   `json:"color"`
	Timezone       string   `json:"timezone"`
	PurgeCacheURLs []string `json:"purgeCacheURLs"`
}

// UpdateGraphPayload is payload for `graph update` subcommand
type UpdateGraphPayload struct {
	Name           string   `json:"name,omitempty"`
	Unit           string   `json:"unit,omitempty"`
	Color          string   `json:"color,omitempty"`
	Timezone       string   `json:"timezone,omitempty"`
	PurgeCacheURLs []string `json:"purgeCacheURLs,omitempty"`
}

// PixelsDateList is response for `graph pixels` subcommand
type PixelsDateList struct {
	Pixels []string `json:"pixels"`
}

// CreateGraph is method for `graph create` subcommand
func (pixela *Pixela) CreateGraph(id, name, unit, numType, color, timezone, selfSufficient string) (NoneGetResponseBody, error) {
	// create payload
	pl := CreateGraphPayload{
		ID:      id,
		Name:    name,
		Unit:    unit,
		NumType: numType,
		Color:   color,
	}

	// argument validation
	vf := validateField{
		GraphID:  id,
		UnitType: numType,
		Color:    color,
	}

	if len(timezone) != 0 {
		pl.Timezone = timezone
	}

	if len(selfSufficient) != 0 {
		pl.SelfSufficient = selfSufficient
		vf.SelfSufficient = selfSufficient
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph create`: wrong arguments")
	}

	plJSON, err := json.Marshal(pl)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph create`: can not marshal request payload")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf("%s/v1/users/%s/graphs", baseURL, pixela.Username)

	// do request
	responseBody, err := pixela.post(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph create`: http request failed")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph create`: http response parse failed")
	}

	return postResponseBody, nil
}

// GetGraphDefinition is method for `graph def` subcommand
func (pixela *Pixela) GetGraphDefinition() (GraphDefinitions, error) {
	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs", baseURL, pixela.Username)

	// do request
	responseBody, err := pixela.get(requestURL)

	if err != nil {
		return GraphDefinitions{}, errors.Wrap(err, "`graph def`: http request failed")
	}

	graphDefinitions := GraphDefinitions{}
	err = json.Unmarshal(responseBody, &graphDefinitions)

	if err != nil {
		return GraphDefinitions{}, errors.Wrap(err, "`graph def`: http response parse failed")
	}

	return graphDefinitions, nil
}

// GetGraphSvg is method for `graph svg` subcommand
func (pixela *Pixela) GetGraphSvg(graphID, date, mode string) ([]byte, error) {
	// argument validation
	vf := validateField{
		GraphID: graphID,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return nil, errors.Wrap(err, "`graph svg`: wrong arguments")
	}

	// build request url
	u, _ := url.Parse(baseURL)
	u.Path = path.Join(u.Path, "v1", "users", pixela.Username, "graphs", graphID)

	// set query
	if len(date) != 0 || len(mode) != 0 {
		q := u.Query()

		if len(date) != 0 {
			q.Set("date", date)
		}

		if len(mode) != 0 {
			q.Set("mode", mode)
		}
		u.RawQuery = q.Encode()
	}

	requestURL := u.String()

	// do request
	responseBody, err := pixela.get(requestURL)

	if err != nil {
		return nil, errors.Wrap(err, "`graph svg`: http request failed")
	}

	return responseBody, nil
}

// UpdateGraph is method for `graph update` subcommand
func (pixela *Pixela) UpdateGraph(graphID string, payload UpdateGraphPayload) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphID: graphID,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph update`: wrong arguments")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s", baseURL, pixela.Username, graphID)

	plJSON, err := json.Marshal(payload)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph update`: can not marshal request payload")
	}

	// do request
	responseBody, err := pixela.put(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph update`: http request failed")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph update`: http response parse failed")
	}

	return postResponseBody, nil
}

// DeleteGraph is method for `graph delete` subcommand
func (pixela *Pixela) DeleteGraph(graphID string) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphID: graphID,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph delete`: wrong arguments")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s", baseURL, pixela.Username, graphID)

	// do request
	responseBody, err := pixela.delete(requestURL)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph delete`: http request failed")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph delete`: http response parse failed")
	}

	return postResponseBody, nil
}

// GetGraphPixelsDateList is method for `graph pixels` subcommand
func (pixela *Pixela) GetGraphPixelsDateList(graphID, from, to string) (PixelsDateList, error) {
	// argument validation
	vf := validateField{
		GraphID: graphID,
		From:    from,
		To:      to,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return PixelsDateList{}, errors.Wrap(err, "`graph pixels`: wrong arguments")
	}

	// build request url
	u, _ := url.Parse(baseURL)
	u.Path = path.Join(u.Path, "v1", "users", pixela.Username, "graphs", graphID, "pixels")

	// set query
	if len(from) != 0 || len(to) != 0 {
		q := u.Query()

		if len(from) != 0 {
			q.Set("from", from)
		}

		if len(to) != 0 {
			q.Set("to", to)
		}
		u.RawQuery = q.Encode()
	}

	requestURL := u.String()

	// do request
	responseBody, err := pixela.get(requestURL)

	if err != nil {
		return PixelsDateList{}, errors.Wrap(err, "`graph pixels`: http request failed")
	}

	pixelsDateList := PixelsDateList{}
	err = json.Unmarshal(responseBody, &pixelsDateList)

	if err != nil {
		return PixelsDateList{}, errors.Wrap(err, "`graph pixels`: http response parse failed")
	}

	return pixelsDateList, nil
}
