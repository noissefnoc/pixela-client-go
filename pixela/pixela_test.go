package pixela

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

// test for pixela.New
func TestNew(t *testing.T) {
	usernameErr := errors.New("initialization error: " + validationErrorMessages["Username"])
	tokenErr := errors.New("initialization error: " + validationErrorMessages["Token"])

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

// test for pixela.post
func TestPixela_post(t *testing.T) {
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
		{"response status not ok", nil, 403, bytes.NewBuffer(errResp), errors.New("post request failed: errorMessage")},
		{"server return invalid response", nil, 200, bytes.NewBufferString("error"), errors.New("post response body parse failed: invalid character 'e' looking for beginning of value")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       ioutil.NopCloser(tt.response),
				Header:     make(http.Header),
			}

			c := NewTestClient(func(req *http.Request) *http.Response {
				if req.Method != http.MethodPost {
					t.Fatalf("want %#v, but %#v", http.MethodPost, req.Method)
				}

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

			pixela, err := New(username, token, debug, OptionHTTPClient(c))

			if err != nil {
				t.Fatalf("got error when http client created %#v", err)
			}

			_, err = pixela.post(baseURL, tt.payload)

			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but %#v", tt.wantErr, err)
				}
			}
		})
	}
}

// test for pixela.get
func TestPixela_get(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		response   *bytes.Buffer
		wantErr    error
	}{
		{"normal case", 200, bytes.NewBufferString("success"), nil},
		{"response status not ok", 403, bytes.NewBuffer(errResp), errors.New("get request failed: errorMessage")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       ioutil.NopCloser(tt.response),
				Header:     make(http.Header),
			}

			c := NewTestClient(func(req *http.Request) *http.Response {
				if req.Method != http.MethodGet {
					t.Fatalf("want %#v, but %#v", http.MethodGet, req.Method)
				}

				if req.Header.Get("X-USER-TOKEN") != token {
					t.Fatalf("want %#v, but %#v", token, req.Header.Get("X-USER-TOKEN"))
				}

				return resp
			})

			pixela, err := New(username, token, debug, OptionHTTPClient(c))

			if err != nil {
				t.Fatalf("got error when http client created %#v", err)
			}

			_, err = pixela.get(baseURL)

			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but %#v", tt.wantErr, err)
				}
			}
		})
	}
}

// test for pixela.put
func TestPixela_put(t *testing.T) {
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
		{"response status not ok", nil, 403, bytes.NewBuffer(errResp), errors.New("put request failed: errorMessage")},
		{"server return invalid response", nil, 200, bytes.NewBufferString("error"), errors.New("put response body parse failed: invalid character 'e' looking for beginning of value")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       ioutil.NopCloser(tt.response),
				Header:     make(http.Header),
			}

			c := NewTestClient(func(req *http.Request) *http.Response {
				if req.Method != http.MethodPut {
					t.Fatalf("want %#v, but %#v", http.MethodPut, req.Method)
				}

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

			pixela, err := New(username, token, debug, OptionHTTPClient(c))

			if err != nil {
				t.Fatalf("got error when http client created %#v", err)
			}

			_, err = pixela.put(baseURL, tt.payload)

			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but %#v", tt.wantErr, err)
				}
			}
		})
	}
}

// test for pixela.delete
func TestPixela_delete(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		response   *bytes.Buffer
		wantErr    error
	}{
		{"normal case", 200, bytes.NewBuffer(scResp), nil},
		{"some error occurred", 200, bytes.NewBuffer(errResp), errors.New("request failed: errorMessage")},
		{"response status not ok", 403, bytes.NewBuffer(errResp), errors.New("delete request failed: errorMessage")},
		{"server return invalid response", 200, bytes.NewBufferString("error"), errors.New("delete response body parse failed: invalid character 'e' looking for beginning of value")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       ioutil.NopCloser(tt.response),
				Header:     make(http.Header),
			}

			c := NewTestClient(func(req *http.Request) *http.Response {
				if req.Method != http.MethodDelete {
					t.Fatalf("want %#v, but %#v", http.MethodDelete, req.Method)
				}

				if req.Header.Get("X-USER-TOKEN") != token {
					t.Fatalf("want %#v, but %#v", token, req.Header.Get("X-USER-TOKEN"))
				}

				return resp
			})

			pixela, err := New(username, token, debug, OptionHTTPClient(c))

			if err != nil {
				t.Fatalf("got error when http client created %#v", err)
			}

			_, err = pixela.delete(baseURL)

			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but %#v", tt.wantErr, err)
				}
			}
		})
	}
}
