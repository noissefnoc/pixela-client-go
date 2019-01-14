package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type CreateWebhookPayload struct {
	GraphId string `json:"graphId"`
	Type    string `json:"type"`
}

func (pixela * Pixela) CreateWebhook(graphId, webhookType string) (NoneGetResponseBody, error)  {
	// create payload
	pl := CreateWebhookPayload{
		GraphId: graphId,
		Type: webhookType,
	}

	plJSON, err := json.Marshal(pl)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "error `webhook create`: can not marshal request payload.")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf("%s/v1/users/%s/webhooks", baseUrl, pixela.Username)

	// do request
	responseBody, err := pixela.post(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "error `webhook create`:http request failed.")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "error `webhook create`:http response parse failed.")
	}

	return postResponseBody, nil
}
