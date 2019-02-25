package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"testing"
)

// common values on following tests
var username = "testuser"
var token = "testtoken"
var debug = false
var graphId = "testgraphid"
var graphName = "testgraphname"
var graphUnit = "testunit"
var numType = "int"
var validColor = "shibafu"
var dateStr = "20000102"
var quantityStr = "100"

// request
var contentType = "application/json"
var contentLength = "Content-Length"
var tokenHeader = "X-USER-TOKEN"
var contentZeroLen = "0"

// http status code
var sucStatus = http.StatusOK
var errStatus = http.StatusBadRequest

// response
var scResp, _ = json.Marshal(NoneGetResponseBody{Message: "success", IsSuccess: true})
var errResp, _ = json.Marshal(NoneGetResponseBody{Message: "errorMessage", IsSuccess: false})
var pixelRespWOp, _ = json.Marshal(GetPixelResponseBody{Quantity: quantityStr, OptionalData: `{"key": "value"}`})
var pixelRespWoOp, _ = json.Marshal(GetPixelResponseBody{Quantity: quantityStr})

// RoundTripFunc
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func newCommandError(command, message string) error {
	return errors.New(fmt.Sprintf("`%s`: %s", command, message))
}

type testCase struct {
	name       string
	statusCode int
	response   []byte
	wantErr    error
	args       []string
}

type noneGetTestCases []testCase

func noneGetRequestHelper(t *testing.T, command string, tests noneGetTestCases, url string) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       ioutil.NopCloser(bytes.NewBuffer(tt.response)),
				Header:     make(http.Header),
			}

			c := NewTestClient(func(req *http.Request) *http.Response {
				if req.URL.String() != url {
					t.Fatalf("want %#v, but got %#v", url, req.URL.String())
				}

				if req.Header.Get(tokenHeader) != token {
					t.Fatalf("want %#v, but got %#v", token, req.Header.Get(tokenHeader))
				}

				return resp
			})

			// skip checking instance creation error
			pixela, err := New(username, token, debug, OptionHTTPClient(c))

			switch command {
			case "user create":
				_, err = pixela.CreateUser(tt.args[0], tt.args[1])
			case "user update":
				_, err = pixela.UpdateUser(tt.args[0])
			case "user delete":
				_, err = pixela.DeleteUser()
			default:
				t.Fatalf("unexpected command %s", command)
			}

			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but got %#v", tt.wantErr.Error(), err.Error())
				}
			}
		})
	}
}
