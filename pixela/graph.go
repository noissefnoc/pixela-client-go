package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type CreateGraphPayload struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Unit     string `json:"unit"`
	NumType  string `json:"type"`
	Color    string `json:"color"`
	Timezone string `json:"timezone,omitempty"`
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

func (pixela *Pixela) CreateGraph(id, name, unit, numType, color, timezone string) (NoneGetResponseBody, error) {
	// create payload
	pl := CreateGraphPayload{
		Id: id,
		Name: name,
		Unit: unit,
		NumType: numType,
		Color: color,
	}

	if timezone != "" {
		pl.Timezone = timezone
	}

	plJSON, err := json.Marshal(pl)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "error `graph create`: can not marshal request payload.")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf("%s/v1/users/%s/graphs", BaseUrl, pixela.Username)

	// do request
	responseBody, err := pixela.post(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "error `graph create`:http request failed.")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "error `graph create`:http response parse failed.")
	}

	return postResponseBody, nil
}

// get graph definition
func (pixela *Pixela) GetGraphDefinition() (GraphDefinitions, error) {
	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs", BaseUrl, pixela.Username)

	// do request
	responseBody, err := pixela.get(requestURL)

	if err != nil {
		return GraphDefinitions{}, errors.Wrap(err, "error `graph def`:http request failed.")
	}

	graphDefinitions := GraphDefinitions{}
	err = json.Unmarshal(responseBody, &graphDefinitions)

	if err != nil {
		return GraphDefinitions{}, errors.Wrap(err, "error `graph def`:http response parse failed.")
	}

	return graphDefinitions, nil
}

// get graph svg html tag
func (pixela *Pixela) GetGraphSvg(graphId string) ([]byte, error) {
	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf(
		"%s/v1/users/%s/graphs/%s", BaseUrl, pixela.Username, graphId)

	// do request
	responseBody, err := pixela.get(requestURL)

	if err != nil {
		return nil, errors.Wrap(err, "error `graph svg`:http request failed.")
	}

	return responseBody, nil
}
