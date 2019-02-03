package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/url"
	"path"
)

type CreateGraphPayload struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Unit           string `json:"unit"`
	NumType        string `json:"type"`
	Color          string `json:"color"`
	Timezone       string `json:"timezone,omitempty"`
	SelfSufficient string `json:"selfSufficient,omitempty"`
}

type GraphDefinitions struct {
	Graphs []struct {
		ID             string   `json:"id"`
		Name           string   `json:"name"`
		Unit           string   `json:"unit"`
		Type           string   `json:"type"`
		Color          string   `json:"color"`
		Timezone       string   `json:"timezone"`
		PurgeCacheURLs []string `json:"purgeCacheURLs"`
	} `json:"graphs"`
}

type UpdateGraphPayload struct {
	Name           string `json:"name,omitempty"`
	Unit           string `json:"unit,omitempty"`
	Color          string `json:"color,omitempty"`
	Timezone       string `json:"timezone,omitempty"`
	PurgeCacheURLs []string `json:"purgeCacheURLs,omitempty"`
}

// create graph
func (pixela *Pixela) CreateGraph(id, name, unit, numType, color, timezone, selfSufficient string) (NoneGetResponseBody, error) {
	// create payload
	pl := CreateGraphPayload{
		Id: id,
		Name: name,
		Unit: unit,
		NumType: numType,
		Color: color,
	}

	// argument validation
	vf := validateField{
		GraphId: id,
		UnitType: numType,
		Color: color,
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
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph create`: can not marshal request payload.")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf("%s/v1/users/%s/graphs", baseUrl, pixela.Username)

	// do request
	responseBody, err := pixela.post(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph create`: http request failed.")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph create`: http response parse failed.")
	}

	return postResponseBody, nil
}

// get graph definition
func (pixela *Pixela) GetGraphDefinition() (GraphDefinitions, error) {
	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs", baseUrl, pixela.Username)

	// do request
	responseBody, err := pixela.get(requestURL)

	if err != nil {
		return GraphDefinitions{}, errors.Wrap(err, "`graph def`: http request failed.")
	}

	graphDefinitions := GraphDefinitions{}
	err = json.Unmarshal(responseBody, &graphDefinitions)

	if err != nil {
		return GraphDefinitions{}, errors.Wrap(err, "`graph def`: http response parse failed.")
	}

	return graphDefinitions, nil
}

// get graph svg html tag
func (pixela *Pixela) GetGraphSvg(graphId, date, mode string) ([]byte, error) {
	// argument validation
	vf := validateField{
		GraphId: graphId,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return nil, errors.Wrap(err, "`graph svg`: wrong arguments")
	}

	// build request url
	u, _ := url.Parse(baseUrl)
	u.Path = path.Join(u.Path, "v1", "users", pixela.Username, "graphs", graphId)

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
		return nil, errors.Wrap(err, "`graph svg`: http request failed.")
	}

	return responseBody, nil
}

// update graph
func (pixela *Pixela) UpdateGraph(graphId string, payload UpdateGraphPayload) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphId: graphId,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph update`: wrong arguments")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s", baseUrl, pixela.Username, graphId)

	plJSON, err := json.Marshal(payload)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph update`: can not marshal request payload.")
	}

	// do request
	responseBody, err := pixela.put(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph update`: http request failed.")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph update`: http response parse failed.")
	}

	return postResponseBody, nil
}

// delete graph
func (pixela *Pixela) DeleteGraph(graphId string) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphId: graphId,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph delete`: wrong arguments")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s", baseUrl, pixela.Username, graphId)

	// do request
	responseBody, err := pixela.delete(requestURL)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph delete`: http request failed.")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`graph delete`: http response parse failed.")
	}

	return postResponseBody, nil
}
