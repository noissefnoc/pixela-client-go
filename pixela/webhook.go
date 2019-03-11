package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

// CreateWebhookPayload is `webhook create` subcommand payload
type CreateWebhookPayload struct {
	GraphID string `json:"graphID"`
	Type    string `json:"type"`
}

// WebhookDefinitions is `webhook get` response
type WebhookDefinitions struct {
	Webhooks []Webhook `json:"webhooks"`
}

// Webhook is internal part of `webhook get` response
type Webhook struct {
	WebhookHash string `json:"webhookHash"`
	GraphID     string `json:"graphID"`
	Type        string `json:"type"`
}

// CreateWebhook is method for `webhook create` subcommand
func (pixela *Pixela) CreateWebhook(graphID, webhookType string) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		GraphID:     graphID,
		WebhookType: webhookType,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`webhook create`: wrong arguments")
	}

	// create payload
	pl := CreateWebhookPayload{
		GraphID: graphID,
		Type:    webhookType,
	}

	plJSON, err := json.Marshal(pl)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`webhook create`: can not marshal request payload")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf("%s/v1/users/%s/webhooks", baseURL, pixela.Username)

	// do request
	responseBody, err := pixela.post(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`webhook create`: http request failed")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`webhook create`: http response parse failed")
	}

	return postResponseBody, nil
}

// GetWebhookDefinitions is method for `webhook get` subcommand
func (pixela *Pixela) GetWebhookDefinitions() (WebhookDefinitions, error) {
	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf("%s/v1/users/%s/webhooks", baseURL, pixela.Username)

	// do request
	responseBody, err := pixela.get(requestURL)

	if err != nil {
		return WebhookDefinitions{}, errors.Wrap(err, "`webhook get`: http request failed")
	}

	getResponseBody := WebhookDefinitions{}
	err = json.Unmarshal(responseBody, &getResponseBody)

	if err != nil {
		return WebhookDefinitions{}, errors.Wrap(err, "`webhook get`: http response parse failed")
	}

	return getResponseBody, nil
}

// InvokeWebhooks is method for `webhook invoke` subcommand
func (pixela *Pixela) InvokeWebhooks(webhookHash string) (NoneGetResponseBody, error) {
	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf("%s/v1/users/%s/webhooks/%s", baseURL, pixela.Username, webhookHash)

	// do request
	responseBody, err := pixela.post(requestURL, nil)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`webhook invoke`: http request failed")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`webhook invoke`: http response parse failed")
	}

	return postResponseBody, nil
}

// DeleteWebhook is method for `webhook delete` subcommand
func (pixela *Pixela) DeleteWebhook(webhookHash string) (NoneGetResponseBody, error) {
	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf("%s/v1/users/%s/webhooks/%s", baseURL, pixela.Username, webhookHash)

	// do request
	responseBody, err := pixela.delete(requestURL)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`webhook delete`: http request failed")
	}

	deleteResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &deleteResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`webhook delete`: http response parse failed")
	}

	return deleteResponseBody, nil
}
