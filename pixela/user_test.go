package pixela

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestPixela_CreateUser(t *testing.T) {
	userCreateUrl := fmt.Sprintf("%s/v1/users", baseUrl)

	respErr := errors.New("`user create`: http request failed: returns none success status code: 400")

	tests := noneGetTestCases{
		{"normal case", sucStatus, scResp, nil, []string{"yes", "yes"}},
		{"status error", errStatus, errResp, respErr, []string{"yes", "yes"}},
	}

	noneGetRequestHelper(t, "user create", tests, userCreateUrl)
}

func TestPixela_UpdateUser(t *testing.T) {
	userUpdateUrl := fmt.Sprintf("%s/v1/users/%s", baseUrl, username)
	respErr := newCommandError("user update", fmt.Sprintf("http request failed: returns none success status code: %d", errStatus))

	tests := noneGetTestCases{
		{"normal case", sucStatus, scResp, nil, []string{"newToken"}},
		{"status error", errStatus, errResp, respErr, []string{"newToken"}},
	}

	noneGetRequestHelper(t, "user update", tests, userUpdateUrl)
}

func TestPixela_DeleteUser(t *testing.T) {
	userDeleteUrl := fmt.Sprintf("%s/v1/users/%s", baseUrl, username)
	respErr := newCommandError("user delete", fmt.Sprintf("http request failed: returns none success status code: %d", errStatus))

	tests := noneGetTestCases{
		{"normal case", sucStatus, scResp, nil, nil},
		{"status error", errStatus, errResp, respErr, nil},
	}

	noneGetRequestHelper(t, "user delete", tests, userDeleteUrl)
}
