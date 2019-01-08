package pixela

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

var BaseUrl string = "https://pixe.la"

// Pixela is application for pixe.la
type Pixela struct {
	Url      string
	Token    string
	Debug    bool
}

type ResponseBody struct {
	Message   string `json:"message"`
	IsSuccess bool   `json:"isSuccess"`
}

func (pixela *Pixela) DoAPI(method string, url string, payload *bytes.Buffer) error {
	// Create Request
	request, err := http.NewRequest(method, url, payload)

	if err != nil {
		return errors.Wrap(err, "can not make request")
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-USER-TOKEN", pixela.Token)

	// Get response from pixe.la
	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return errors.Wrap(err, "pixel record http request failed")
	}

	// Parse response
	defer response.Body.Close()

	responseBodyJSON, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return errors.Wrap(err, "response read failed.")
	}

	responseBody := ResponseBody{}
	err = json.Unmarshal(responseBodyJSON, &responseBody)

	if err != nil {
		return errors.Wrap(err, "response body parse failed.")
	}

	return nil
}
