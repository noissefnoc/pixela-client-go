package pixela

import (
	"bytes"
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
	} {
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
	client := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response {
			StatusCode: 200,
			Body: ioutil.NopCloser(bytes.NewBufferString(`{"message":"success", "isSuccess": true}`)),
			Header: make(http.Header),
		}
	})

	pixela, err := New("testuser", "testtoken", false, OptionHTTPClient(client))

	if err != nil {
		t.Fatalf("got error when test http client created %#v", err)
	}

	got, err := pixela.post("https://examples.com/posttest", bytes.NewBufferString(`{"key": "value"}`))

	if err != nil {
		t.Fatalf("got error when test http client created %#v", err)
	}

	if got == nil {
		t.Fatalf("but %#v", got)
	}
}
