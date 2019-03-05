package pixela

import (
	"fmt"
	"testing"
)

func TestPixela_CreateUser(t *testing.T) {
	userCreateUrl := fmt.Sprintf("%s/v1/users", baseUrl)

	ivAToSErr := newCommandError(userCreate, "wrong arguments: "+validationErrorMessages["AgreeTermsOfService"])
	ivNMErr := newCommandError(userCreate, "wrong arguments: "+validationErrorMessages["NotMinor"])
	respDataErr := newCommandError(userCreate, "http request failed: post request failed: errorMessage")

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{"yes", "yes"}},
		{"status error", errStatus, errResp, respDataErr, []string{"yes", "yes"}},
		{"not agree statement", errStatus, errResp, respDataErr, []string{"no", "yes"}},
		{"not minor use", errStatus, errResp, respDataErr, []string{"yes", "no"}},
		{"entire disagree statement", errStatus, errResp, respDataErr, []string{"no", "no"}},
		{"invalid agree statement", 0, errResp, ivAToSErr, []string{"hoge", "yes"}},
		{"invalid usage", 0, errResp, ivNMErr, []string{"yes", "hoge"}},
	}

	subCommandTestHelper(t, userCreate, tests, userCreateUrl)
}

func TestPixela_UpdateUser(t *testing.T) {
	userUpdateUrl := fmt.Sprintf("%s/v1/users/%s", baseUrl, username)

	respDataErr := newCommandError(userUpdate, "http request failed: put request failed: errorMessage")

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{"newToken"}},
		{"status error", errStatus, errResp, respDataErr, []string{"newToken"}},
	}

	subCommandTestHelper(t, userUpdate, tests, userUpdateUrl)
}

func TestPixela_DeleteUser(t *testing.T) {
	userDeleteUrl := fmt.Sprintf("%s/v1/users/%s", baseUrl, username)

	respDataErr := newCommandError(userDelete, "http request failed: delete request failed: errorMessage")

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, nil},
		{"status error", errStatus, errResp, respDataErr, nil},
	}

	subCommandTestHelper(t, userDelete, tests, userDeleteUrl)
}
