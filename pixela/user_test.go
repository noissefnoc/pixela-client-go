package pixela

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPixela_CreateUser(t *testing.T) {
	userCreateUrl := fmt.Sprintf("%s/v1/users", baseUrl)

	respErr := errors.New("`user create`: http request failed.: returns none success status code: 400")

	tests := []struct {
		name                string
		token               string
		username            string
		agreeTermsOfService string
		notMinor            string
		statusCode          int
		response            *bytes.Buffer
		wantErr             error
	}{
		{"normal case", username, token, "yes", "yes", http.StatusOK, bytes.NewBuffer(scResp), nil},
		{"return error response", username, token, "yes", "yes", http.StatusBadRequest, bytes.NewBuffer(errResp), respErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       ioutil.NopCloser(tt.response),
				Header:     make(http.Header),
			}

			c := NewTestClient(func(req *http.Request) *http.Response {
				if req.URL.String() != userCreateUrl {
					t.Fatalf("want %#v, but got %#v", userCreateUrl, req.URL.String())
				}

				if req.Header.Get(tokenHeader) != token {
					t.Fatalf("want %#v, but got %#v", token, req.Header.Get(tokenHeader))
				}

				return resp
			})

			// skip checking instance creation error
			pixela, _ := New(username, token, debug, OptionHTTPClient(c))

			_, err := pixela.CreateUser(tt.agreeTermsOfService, tt.notMinor)

			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but got %#v", tt.wantErr.Error(), err.Error())
				}
			}
		})
	}
}