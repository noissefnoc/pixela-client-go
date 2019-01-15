package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

var baseUrl = "https://pixe.la"

// Pixela is application for pixe.la
type Pixela struct {
	HTTPClient http.Client
	URL        string
	Username   string
	Token      string
	Debug      bool
}

// pixe.la response body that post, put and delete method requested
type NoneGetResponseBody struct {
	Message     string `json:"message"`
	IsSuccess   bool   `json:"isSuccess"`
	WebhookHash string `json:"webhookHash,omitempty"`
}

func New(username, token string, debug bool) (*Pixela, error) {
	if username == "" || token == "" {
		return nil, errors.New("initialization error: username and token required\n")
	}

	return &Pixela{
		HTTPClient: http.Client{
			Timeout: time.Duration(10) * time.Second,
		},
		URL:      baseUrl,
		Username: username,
		Token:    token,
		Debug:    debug,
	}, nil
}

// post request
func (pixela *Pixela) post(url string, payload *bytes.Buffer) ([]byte, error) {
	// create Request
	request := &http.Request{}
	var err error

	if payload == nil {
		request, err = http.NewRequest(http.MethodPost, url, nil)
		request.Header.Set("Content-Length", "0")
	} else {
		request, err = http.NewRequest(http.MethodPost, url, payload)
	}

	if err != nil {
		return nil, errors.Wrap(err, "can not make request")
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-USER-TOKEN", pixela.Token)

	// get response from pixe.la
	response, err := pixela.HTTPClient.Do(request)

	if err != nil {
		return nil, errors.Wrap(err, "pixel record http request failed")
	}

	// parse response
	defer response.Body.Close()

	responseBodyJSON, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.Wrap(err, "response read failed.")
	}

	responseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBodyJSON, &responseBody)

	if err != nil {
		return nil, errors.Wrap(err, "response body parse failed.")
	}

	// check response body if request success
	if !responseBody.IsSuccess {
		return nil, errors.New(fmt.Sprintf("request failed: %s", responseBody.Message))
	}

	return responseBodyJSON, nil
}

// get request
func (pixela *Pixela) get(url string) ([]byte, error) {
	// create Request
	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, errors.Wrap(err, "can not make request")
	}

	request.Header.Set("X-USER-TOKEN", pixela.Token)

	// get response from pixe.la
	response, err := pixela.HTTPClient.Do(request)

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

// put request
func (pixela *Pixela) put(url string, payload *bytes.Buffer) ([]byte, error) {
	request := &http.Request{}
	var err error

	// create Request
	if payload == nil {
		request, err = http.NewRequest(http.MethodPut, url, nil)
		request.Header.Set("Content-Length", "0")
	} else {
		request, err = http.NewRequest(http.MethodPut, url, payload)
	}

	if err != nil {
		return nil, errors.Wrap(err, "can not make request")
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-USER-TOKEN", pixela.Token)

	// get response from pixe.la
	response, err := pixela.HTTPClient.Do(request)

	if err != nil {
		return nil, errors.Wrap(err, "http request failed")
	}

	// parse response
	defer response.Body.Close()

	responseBodyJSON, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.Wrap(err, "response read failed.")
	}

	responseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBodyJSON, &responseBody)

	if err != nil {
		return nil, errors.Wrap(err, "response body parse failed.")
	}

	// check response body if request success
	if !responseBody.IsSuccess {
		return nil, errors.New(fmt.Sprintf("request failed: %s", responseBody.Message))
	}

	return responseBodyJSON, nil
}

// delete request
func (pixela *Pixela) delete(url string) ([]byte, error) {
	// create Request
	request, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		return nil, errors.Wrap(err, "can not make request")
	}

	request.Header.Set("X-USER-TOKEN", pixela.Token)

	// get response from pixe.la
	response, err := pixela.HTTPClient.Do(request)

	if err != nil {
		return nil, errors.Wrap(err, "http delete request failed")
	}

	// parse response
	defer response.Body.Close()

	responseBodyJSON, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.Wrap(err, "delete response read failed.")
	}

	responseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBodyJSON, &responseBody)

	if err != nil {
		return nil, errors.Wrap(err, "response body parse failed.")
	}

	// check response body if request success
	if !responseBody.IsSuccess {
		return nil, errors.New(fmt.Sprintf("request failed: %s", responseBody.Message))
	}

	return responseBodyJSON, nil
}
