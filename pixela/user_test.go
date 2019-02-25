package pixela

import (
	"fmt"
	"testing"
)

func TestPixela_CreateUser(t *testing.T) {
	userCreateUrl := fmt.Sprintf("%s/v1/users", baseUrl)
	respErr := newCommandError(userCreate, fmt.Sprintf("http request failed: returns none success status code: %d", errStatus))

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{"yes", "yes"}},
		{"status error", errStatus, errResp, respErr, []string{"yes", "yes"}},
	}

	subCommandTestHelper(t, userCreate, tests, userCreateUrl)
}

func TestPixela_UpdateUser(t *testing.T) {
	userUpdateUrl := fmt.Sprintf("%s/v1/users/%s", baseUrl, username)
	respErr := newCommandError(userUpdate, fmt.Sprintf("http request failed: returns none success status code: %d", errStatus))

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{"newToken"}},
		{"status error", errStatus, errResp, respErr, []string{"newToken"}},
	}

	subCommandTestHelper(t, userUpdate, tests, userUpdateUrl)
}

func TestPixela_DeleteUser(t *testing.T) {
	userDeleteUrl := fmt.Sprintf("%s/v1/users/%s", baseUrl, username)
	respErr := newCommandError(userDelete, fmt.Sprintf("http request failed: returns none success status code: %d", errStatus))

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, nil},
		{"status error", errStatus, errResp, respErr, nil},
	}

	subCommandTestHelper(t, userDelete, tests, userDeleteUrl)
}
