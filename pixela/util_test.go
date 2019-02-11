package pixela

import (
	"encoding/json"
	"net/http"
)

// common values on following tests
var username = "testuser"
var token = "testtoken"
var debug = false
var graphId = "testgraphid"
var dateStr = "20000102"
var quantityStr = "100"

// request
var contentType = "application/json"
var contentLength = "Content-Length"
var tokenHeader = "X-USER-TOKEN"
var contentZeroLen = "0"

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
