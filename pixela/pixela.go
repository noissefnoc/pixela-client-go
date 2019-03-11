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

var baseURL = "https://pixe.la"

// Pixela is application for pixe.la
type Pixela struct {
	HTTPClient *http.Client
	URL        string
	Username   string
	Validator  Validator
	Token      string
	Debug      bool
}

// Option is customize Pixela properties function
type Option func(*Pixela)

// OptionHTTPClient - provide a custom http client to the HTTPClient
func OptionHTTPClient(c *http.Client) Option {
	return func(pixela *Pixela) {
		pixela.HTTPClient = c
	}
}

// NoneGetResponseBody - pixe.la response body that post, put and delete method requested
type NoneGetResponseBody struct {
	Message     string `json:"message"`
	IsSuccess   bool   `json:"isSuccess"`
	WebhookHash string `json:"webhookHash,omitempty"`
}

// New creates pixe.la api client instance
func New(username, token string, debug bool, opts ...Option) (*Pixela, error) {
	// validate arguments
	vf := newInstanceValidateField{
		Username: username,
		Token:    token,
	}

	validate := newValidator()
	err := validate.Validate(vf)

	if err != nil {
		return nil, errors.Wrap(err, "initialization error")
	}

	// create instance
	pixela := &Pixela{
		HTTPClient: &http.Client{
			Timeout: time.Duration(10) * time.Second,
		},
		URL:       baseURL,
		Username:  username,
		Token:     token,
		Validator: validate,
		Debug:     debug,
	}

	for _, opt := range opts {
		opt(pixela)
	}

	return pixela, nil
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
		return nil, errors.Wrap(err, "http post request failed")
	}

	// parse response
	defer response.Body.Close()

	responseBodyJSON, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.Wrap(err, "post response read failed")
	}

	responseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBodyJSON, &responseBody)

	if err != nil {
		return nil, errors.Wrap(err, "post response body parse failed")
	}

	// check response body if request success
	if response.StatusCode != http.StatusOK && !responseBody.IsSuccess {
		return nil, fmt.Errorf("post request failed: %s", responseBody.Message)
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
		return nil, errors.Wrap(err, "http get request failed")
	}

	// parse response
	defer response.Body.Close()

	responseBodyJSON, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.Wrap(err, "get response read failed")
	}

	if response.StatusCode != http.StatusOK {
		responseBody := NoneGetResponseBody{}
		err = json.Unmarshal(responseBodyJSON, &responseBody)

		if err != nil {
			return nil, errors.Wrap(err, "get response parse failed")
		}

		return nil, fmt.Errorf("get request failed: %s", responseBody.Message)
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
		return nil, errors.Wrap(err, "http put request failed")
	}

	// parse response
	defer response.Body.Close()

	responseBodyJSON, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.Wrap(err, "put response read failed")
	}

	responseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBodyJSON, &responseBody)

	if err != nil {
		return nil, errors.Wrap(err, "put response body parse failed")
	}

	// check response body if request success
	if response.StatusCode != http.StatusOK && !responseBody.IsSuccess {
		return nil, fmt.Errorf("put request failed: %s", responseBody.Message)
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
		return nil, errors.Wrap(err, "delete response read failed")
	}

	responseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBodyJSON, &responseBody)

	if err != nil {
		return nil, errors.Wrap(err, "delete response body parse failed")
	}

	// check response body if request success
	if response.StatusCode != http.StatusOK && !responseBody.IsSuccess {
		return nil, fmt.Errorf("delete request failed: %s", responseBody.Message)
	}

	return responseBodyJSON, nil
}
