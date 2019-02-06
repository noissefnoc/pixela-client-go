package pixela

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	usernameErr := errors.New("initialization error.: " + validationErrorMessages["Username"])
	tokenErr := errors.New("initialization error.: " + validationErrorMessages["Token"])

	tests := []struct {
		name     string
		username string
		token    string
		debug    bool
		wantErr  error
	}{
		{"normal case", "testuser", "testtoken", false, nil},
		{"username empty", "", "testtoken", false, usernameErr},
		{"username invalid", "123", "testtoken", false, usernameErr},
		{"username too short", "a", "testtoken", false, usernameErr},
		{"username too long", strings.Repeat("a", 34), "testtoken", false, usernameErr},
		{"token empty", "testuser", "", false, tokenErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.username, tt.token, tt.debug)

			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but %#v", tt.wantErr, err)
				}
			}
		})
	}
}

func TestPixela_post(t *testing.T) {
	url := "https://examples.com/post"
	username := "testuser"
	token := "testtoken"
	debug := false
	contentType := "application/json"
	contentZeroLen := "0"
	scResp, _ := json.Marshal(NoneGetResponseBody{Message: "success", IsSuccess: true})
	errResp, _ := json.Marshal(NoneGetResponseBody{Message: "errorMessage", IsSuccess: false})

	tests := []struct {
		name       string
		payload    *bytes.Buffer
		statusCode int
		response   *bytes.Buffer
		wantErr    error
	}{
		{"normal case without payload", nil, 200, bytes.NewBuffer(scResp), nil},
		{"normal case with payload", bytes.NewBufferString(`{"key": "value"}`), 200, bytes.NewBuffer(scResp), nil},
		{"some error occurred", nil, 200, bytes.NewBuffer(errResp), errors.New("request failed: errorMessage")},
		{"response status not ok", nil, 403, bytes.NewBuffer(errResp), errors.New("returns none success status code: 403")},
		{"server return invalid response", nil, 200, bytes.NewBufferString("error"), errors.New("response body parse failed.: invalid character 'e' looking for beginning of value")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       ioutil.NopCloser(tt.response),
				Header:     make(http.Header),
			}

			srClient := NewTestClient(func(req *http.Request) *http.Response {
				if tt.payload == nil {
					if req.Header.Get("Content-Length") != contentZeroLen {
						t.Fatalf("want %#v, but %#v", contentZeroLen, req.Header.Get("Content-Length"))
					}
				}

				if req.Header.Get("Content-Type") != contentType {
					t.Fatalf("want %#v, but %#v", contentType, req.Header.Get("Content-Type"))
				}

				if req.Header.Get("X-USER-TOKEN") != token {
					t.Fatalf("want %#v, but %#v", token, req.Header.Get("X-USER-TOKEN"))
				}

				return resp
			})

			pixela, err := New(username, token, debug, OptionHTTPClient(srClient))

			if err != nil {
				t.Fatalf("got error when http client created %#v", err)
			}

			_, err = pixela.post(url, tt.payload)

			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but %#v", tt.wantErr, err)
				}
			}
		})
	}
}
