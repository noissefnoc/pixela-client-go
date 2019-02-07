package pixela

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// common values on following tests
var urlStr = "https://examples.com/"
var username = "testuser"
var token = "testtoken"
var debug = false
var graphId = "testgraphid"
var dateStr = "20000102"
var quantityStr = "100"

// urls
var pixelCreateUrl = fmt.Sprintf("%s/v1/users/%s/graphs/%s", urlStr, username, graphId)
var pixelGetUrl = fmt.Sprintf("%s/v1/users/%s/graphs/%s/%s", urlStr, username, graphId, dateStr)

// request
var contentType = "application/json"
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
