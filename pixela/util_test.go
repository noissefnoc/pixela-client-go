package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

// common values on following tests
var username = "testuser"
var token = "testtoken"
var debug = false
var graphID = "testgraphid"
var graphName = "testgraphname"
var graphUnit = "testunit"
var numType = "int"
var validColor = "shibafu"
var dateStr = "20000102"
var quantityStr = "100"
var webhookHash = "hash"
var webhookType = "increment"

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
var ivResp = []byte("hoge")
var pixelRespWOp, _ = json.Marshal(GetPixelResponseBody{Quantity: quantityStr, OptionalData: `{"key": "value"}`})
var pixelRespWoOp, _ = json.Marshal(GetPixelResponseBody{Quantity: quantityStr})
var webhookResp, _ = json.Marshal(WebhookDefinitions{[]Webhook{{webhookHash, graphID, webhookType}}})
var graphDefResp, _ = json.Marshal(GraphDefinitions{[]Graph{{graphID, graphName, graphUnit, numType, validColor, "Asia/Tokyo", []string{""}}}})
var graphSvgResp = `<sgv>test</svg>`
var graphPixelsResp, _ = json.Marshal(PixelsDateList{[]string{"20190101", "20190102"}})

// sub commands
type subCommand int

const (
	userCreate subCommand = iota
	userUpdate
	userDelete
	pixelPost
	pixelGet
	pixelIncrement
	pixelDecrement
	pixelDelete
	pixelUpdate
	webhookCreate
	webhookGet
	webhookInvoke
	webhookDelete
	graphCreate
	graphUpdate
	graphDelete
	graphGet
	graphSvg
	graphPixels
	graphDetail
)

var subCommandStringMap = map[subCommand]string{
	userCreate:     "user create",
	userUpdate:     "user update",
	userDelete:     "user delete",
	pixelPost:      "pixel post",
	pixelGet:       "pixel get",
	pixelIncrement: "pixel increment",
	pixelDecrement: "pixel decrement",
	pixelDelete:    "pixel delete",
	pixelUpdate:    "pixel update",
	webhookCreate:  "webhook create",
	webhookGet:     "webhook get",
	webhookInvoke:  "webhook invoke",
	webhookDelete:  "webhook delete",
	graphCreate:    "graph create",
	graphUpdate:    "graph update",
	graphDelete:    "graph delete",
	graphGet:       "graph get",
	graphSvg:       "graph svg",
	graphPixels:    "graph pixels",
	graphDetail:    "graph detail",
}

var subCommandMethodMap = map[subCommand]string{
	userCreate:     http.MethodPost,
	userUpdate:     http.MethodPut,
	userDelete:     http.MethodDelete,
	pixelPost:      http.MethodPost,
	pixelGet:       http.MethodGet,
	pixelIncrement: http.MethodPut,
	pixelDecrement: http.MethodPut,
	pixelDelete:    http.MethodDelete,
	pixelUpdate:    http.MethodPut,
	webhookCreate:  http.MethodPost,
	webhookGet:     http.MethodGet,
	webhookInvoke:  http.MethodPost,
	webhookDelete:  http.MethodDelete,
	graphCreate:    http.MethodPost,
	graphUpdate:    http.MethodPut,
	graphDelete:    http.MethodDelete,
	graphGet:       http.MethodGet,
	graphSvg:       http.MethodGet,
	graphPixels:    http.MethodGet,
}

func (c subCommand) String() string {
	if stringer, ok := subCommandStringMap[c]; ok {
		return stringer
	}

	panic("unknown value")
}

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

func newCommandError(command subCommand, message string) error {
	return fmt.Errorf("`%s`: %s", command, message)
}

type testCase struct {
	name       string
	statusCode int
	response   []byte
	wantErr    error
	args       []string
}

type testCases []testCase

func subCommandTestHelper(t *testing.T, cmd subCommand, tests testCases, urlStr string) {
	orgURLStr := urlStr

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       ioutil.NopCloser(bytes.NewBuffer(tt.response)),
				Header:     make(http.Header),
			}

			switch cmd {
			case graphSvg:
				urlStr = queryBuilder(orgURLStr, "date", tt.args[1], "mode", tt.args[2])
			case graphPixels:
				urlStr = queryBuilder(orgURLStr, "from", tt.args[1], "to", tt.args[2])
			}

			c := NewTestClient(func(req *http.Request) *http.Response {
				if req.Header.Get(tokenHeader) != token {
					t.Fatalf("want %#v, but got %#v", token, req.Header.Get(tokenHeader))
				}

				if req.URL.String() != urlStr {
					t.Fatalf("want %#v, but got %#v", urlStr, req.URL.String())
				}

				if method, ok := subCommandMethodMap[cmd]; ok {
					if req.Method != method {
						t.Fatalf("want %#v, but got %#v", method, req.Method)
					}

					if cmd == pixelIncrement || cmd == pixelDecrement {
						if req.Header.Get(contentLength) != contentZeroLen {
							t.Fatalf("want %#v, but got %#v", contentZeroLen, req.Header.Get(contentLength))
						}
					}
				}

				return resp
			})

			// skip checking instance creation error
			pixela, err := New(username, token, debug, OptionHTTPClient(c))
			err = subCommandMethodCall(pixela, tt, cmd)

			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but got %#v", tt.wantErr.Error(), err.Error())
				}
			}
		})
	}
}

func subCommandMethodCall(pixela *Pixela, tt testCase, cmd subCommand) error {
	var err error

	switch cmd {
	case userCreate:
		_, err = pixela.CreateUser(tt.args[0], tt.args[1])
	case userUpdate:
		_, err = pixela.UpdateUser(tt.args[0])
	case userDelete:
		_, err = pixela.DeleteUser()
	case pixelPost:
		_, err = pixela.PostPixel(tt.args[0], tt.args[1], tt.args[2], tt.args[3])
	case pixelGet:
		_, err = pixela.GetPixel(tt.args[0], tt.args[1])
	case pixelIncrement:
		_, err = pixela.IncrementPixel(tt.args[0])
	case pixelDecrement:
		_, err = pixela.DecrementPixel(tt.args[0])
	case pixelDelete:
		_, err = pixela.DeletePixel(tt.args[0], tt.args[1])
	case pixelUpdate:
		_, err = pixela.UpdatePixel(tt.args[0], tt.args[1], tt.args[2], tt.args[3])
	case webhookCreate:
		_, err = pixela.CreateWebhook(tt.args[0], tt.args[1])
	case webhookGet:
		_, err = pixela.GetWebhookDefinitions()
	case webhookInvoke:
		_, err = pixela.InvokeWebhooks(tt.args[0])
	case webhookDelete:
		_, err = pixela.DeleteWebhook(tt.args[0])
	case graphCreate:
		_, err = pixela.CreateGraph(tt.args[0], tt.args[1], tt.args[2], tt.args[3], tt.args[4], tt.args[5], tt.args[6])
	case graphUpdate:
		payload := UpdateGraphPayload{
			tt.args[1],
			tt.args[2],
			tt.args[3],
			tt.args[4],
			[]string{tt.args[5]},
		}
		_, err = pixela.UpdateGraph(tt.args[0], payload)
	case graphDelete:
		_, err = pixela.DeleteGraph(tt.args[0])
	case graphGet:
		_, err = pixela.GetGraphDefinition()
	case graphSvg:
		_, err = pixela.GetGraphSvg(tt.args[0], tt.args[1], tt.args[2])
	case graphPixels:
		_, err = pixela.GetGraphPixelsDateList(tt.args[0], tt.args[1], tt.args[2])
	default:
		err = fmt.Errorf("unexpected cmd %s", cmd)
	}

	return err
}

func queryBuilder(urlStr, key1, value1, key2, value2 string) string {
	u, _ := url.Parse(urlStr)

	if len(value1) != 0 || len(value2) != 0 {
		q := u.Query()

		if len(value1) != 0 {
			q.Set(key1, value1)
		}

		if len(value2) != 0 {
			q.Set(key2, value2)
		}
		u.RawQuery = q.Encode()
	}

	return u.String()
}
