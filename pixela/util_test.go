package pixela

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
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
	pixelCreate
	pixelGet
	pixelInc
	pixelDec
	pixelDelete
	pixelUpdate
	webhookCreate
	webhookGet
	webhookInvoke
	webhookDelete
	graphCreate
	graphUpdate
	graphDelete
	graphDef
	graphSvg
	graphPixels
)

func (c subCommand) String() string {
	switch c {
	case userCreate:
		return "user create"
	case userUpdate:
		return "user update"
	case userDelete:
		return "user delete"
	case pixelCreate:
		return "pixel create"
	case pixelGet:
		return "pixel get"
	case pixelInc:
		return "pixel inc"
	case pixelDec:
		return "pixel dec"
	case pixelDelete:
		return "pixel delete"
	case pixelUpdate:
		return "pixel update"
	case webhookCreate:
		return "webhook create"
	case webhookGet:
		return "webhook get"
	case webhookInvoke:
		return "webhook invoke"
	case webhookDelete:
		return "webhook delete"
	case graphCreate:
		return "graph create"
	case graphUpdate:
		return "graph update"
	case graphDelete:
		return "graph delete"
	case graphDef:
		return "graph def"
	case graphSvg:
		return "graph svg"
	case graphPixels:
		return "graph pixels"
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
	return errors.New(fmt.Sprintf("`%s`: %s", command, message))
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       ioutil.NopCloser(bytes.NewBuffer(tt.response)),
				Header:     make(http.Header),
			}

			c := NewTestClient(func(req *http.Request) *http.Response {
				if req.Header.Get(tokenHeader) != token {
					t.Fatalf("want %#v, but got %#v", token, req.Header.Get(tokenHeader))
				}

				switch cmd {
				case graphSvg:
					urlStr = queryBuilder(urlStr, "date", tt.args[1], "mode", tt.args[2])
				case graphPixels:
					urlStr = queryBuilder(urlStr, "from", tt.args[1], "to", tt.args[2])
				default:
					if req.URL.String() != urlStr {
						t.Fatalf("want %#v, but got %#v", urlStr, req.URL.String())
					}
				}

				switch cmd {
				case pixelGet, webhookGet, graphDef, graphSvg:
					if req.Method != http.MethodGet {
						t.Fatalf("want %#v, but got %#v", "GET", req.Method)
					}
				case userCreate, pixelCreate, webhookCreate, webhookInvoke, graphCreate:
					if req.Method != http.MethodPost {
						t.Fatalf("want %#v, but got %#v", "POST", req.Method)
					}
				case userUpdate, pixelInc, pixelDec, pixelUpdate, graphUpdate:
					if req.Method != http.MethodPut {
						t.Fatalf("want %#v, but got %#v", "PUT", req.Method)
					}

					if cmd == pixelInc || cmd == pixelDec {
						if req.Header.Get(contentLength) != contentZeroLen {
							t.Fatalf("want %#v, but got %#v", contentZeroLen, req.Header.Get(contentLength))
						}
					}
				case userDelete, pixelDelete, webhookDelete:
					if req.Method != http.MethodDelete {
						t.Fatalf("want %#v, but got %#v", "DELETE", req.Method)
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
	case pixelCreate:
		_, err = pixela.CreatePixel(tt.args[0], tt.args[1], tt.args[2], tt.args[3])
	case pixelGet:
		_, err = pixela.GetPixel(tt.args[0], tt.args[1])
	case pixelInc:
		_, err = pixela.IncPixel(tt.args[0])
	case pixelDec:
		_, err = pixela.DecPixel(tt.args[0])
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
	case graphDef:
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
