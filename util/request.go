package util

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type ResponseBody struct {
	Message string `json:"message"`
	IsSuccess string `json:"isSuccess"`
}

func DoRequest(method string, url string, payload *bytes.Buffer) error {
	token := viper.GetString("token")

	request, err := http.NewRequest(method, url, payload)

	if err != nil {
		return errors.Wrap(err, "can not make request")
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-USER-TOKEN", token)

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return errors.Wrap(err, "pixel record http request failed")
	}

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
