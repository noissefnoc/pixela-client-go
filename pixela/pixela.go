package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

var BaseUrl = "https://pixe.la"

// Pixela is application for pixe.la
type Pixela struct {
	Username string
	Token    string
	Debug    bool
}

type PostResponseBody struct {
	Message     string `json:"message"`
	IsSuccess   bool   `json:"isSuccess"`
	WebhookHash string `json:"webhookHash,omitempty"`
}

func (pixela *Pixela) post(url string, payload *bytes.Buffer) ([]byte, error) {
	// create Request
	request, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		return nil, errors.Wrap(err, "can not make request")
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-USER-TOKEN", pixela.Token)

	// get response from pixe.la
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, errors.Wrap(err, "pixel record http request failed")
	}

	// parse response
	defer response.Body.Close()

	responseBodyJSON, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.Wrap(err, "response read failed.")
	}

	responseBody := PostResponseBody{}
	err = json.Unmarshal(responseBodyJSON, &responseBody)

	if err != nil {
		return nil, errors.Wrap(err, "response body parse failed.")
	}

	// check response body if request success
	if !responseBody.IsSuccess {
		return nil, errors.Wrap(err, fmt.Sprintf("request failed: %s", responseBody.Message))
	}

	return responseBodyJSON, nil
}

func (pixela *Pixela) get(url string) ([]byte, error) {
	// create Request
	request, err := http.NewRequest(http.MethodGet, url,nil)

	if err != nil {
		return nil, errors.Wrap(err, "can not make request")
	}

	request.Header.Set("X-USER-TOKEN", pixela.Token)

	// get response from pixe.la
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, errors.Wrap(err, "pixel record http request failed")
	}

	// parse response
	defer response.Body.Close()

	responseBodyJSON, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.Wrap(err, "response read failed.")
	}

	return responseBodyJSON, nil
}
