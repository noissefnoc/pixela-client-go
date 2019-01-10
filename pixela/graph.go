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

func (pixela *Pixela) CreateGraph(id, name, unit, numType, color, timezone string) (PostResponseBody, error) {
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
		return PostResponseBody{}, errors.Wrap(err, "error `graph create`: can not marshal request payload.")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf("%s/v1/users/%s/graphs", BaseUrl, pixela.Username)

	// do request
	responseBody, err := pixela.post(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return PostResponseBody{}, errors.Wrap(err, "error `graph create`:http request failed.")
	}

	postResponseBody := PostResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return PostResponseBody{}, errors.Wrap(err, "error `graph create`:http response parse failed.")
	}

	return postResponseBody, nil
}
