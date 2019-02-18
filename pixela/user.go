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

type UpdateUserPayload struct {
	NewToken string `json:"newToken"`
}

// create user
func (pixela *Pixela) CreateUser(agreeTermsOfService, notMinor string) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		AgreeTermsOfService: agreeTermsOfService,
		NotMinor:            notMinor,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`user create`: wrong arguments")
	}

	// create payload
	pl := CreateUserPayload{
		Username:            pixela.Username,
		Token:               pixela.Token,
		AgreeTermsOfService: agreeTermsOfService,
		NotMinor:            notMinor,
	}

	plJSON, err := json.Marshal(pl)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`user create`: can not marshal request payload")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf("%s/v1/users", baseUrl)

	// do request
	responseBody, err := pixela.post(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`user create`: http request failed")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`user create`: http response parse failed")
	}

	return postResponseBody, nil
}

// update user token
func (pixela *Pixela) UpdateUser(newToken string) (NoneGetResponseBody, error) {
	// argument validation
	vf := validateField{
		NewToken: newToken,
	}

	err := pixela.Validator.Validate(vf)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`user update`: wrong arguments")
	}

	// create payload
	pl := UpdateUserPayload{
		NewToken: newToken,
	}

	plJSON, err := json.Marshal(pl)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`user update`: can not marshal request payload")
	}

	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf("%s/v1/users/%s", baseUrl, pixela.Username)

	// do request
	responseBody, err := pixela.put(requestURL, bytes.NewBuffer(plJSON))

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`user update`: http request failed")
	}

	postResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &postResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`user update`: http response parse failed")
	}

	return postResponseBody, nil
}

// delete user
func (pixela *Pixela) DeleteUser() (NoneGetResponseBody, error) {
	// build request url
	// TODO: rewrite by url package
	requestURL := fmt.Sprintf("%s/v1/users/%s", baseUrl, pixela.Username)

	// do request
	responseBody, err := pixela.delete(requestURL)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`user delete`: http request failed")
	}

	deleteResponseBody := NoneGetResponseBody{}
	err = json.Unmarshal(responseBody, &deleteResponseBody)

	if err != nil {
		return NoneGetResponseBody{}, errors.Wrap(err, "`user delete`: http response parse failed")
	}

	return deleteResponseBody, nil
}
