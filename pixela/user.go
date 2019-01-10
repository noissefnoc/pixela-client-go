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
func (pixela *Pixela) CreateUser() error {
	// create payload
	pl := CreateUserPayload{
		Username:            pixela.Username,
		Token:               pixela.Token,
		AgreeTermsOfService: "yes",
		NotMinor:            "yes",
	}

	plJSON, err := json.Marshal(pl)

	if err != nil {
		return errors.Wrap(err, "error `userl create`: can not marshal request payload.")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf("%s/v1/users", BaseUrl)

	// do request
	err = pixela.post(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return errors.Wrap(err, "error `user create`:http request failed.")
	}

	return nil
}
