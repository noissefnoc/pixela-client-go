package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type CreateUserPayload struct {
	Username            string `json:"username"`
	Token               string `json:"token"`
	AgreeTermsOfService string `json:"agreeTermsOfService"`
	NotMinor            string `json:"notMinor"`
}

// create user
func (pixela *Pixela) CreateUser() (NoneGetResponseBody, error) {
	// create payload
	pl := CreateUserPayload{
		Username:            pixela.Username,
		Token:               pixela.Token,
		AgreeTermsOfService: "yes",
		NotMinor:            "yes",
	}

	plJSON, err := json.Marshal(pl)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "error `user create`: can not marshal request payload.")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf("%s/v1/users", baseUrl)

	// do request
	responseBody, err := pixela.post(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "error `user create`:http request failed.")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "error `user create`:http response parse failed.")
	}

	return postResponseBody, nil
}
