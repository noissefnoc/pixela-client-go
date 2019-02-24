package pixela

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPixela_CreateGraph(t *testing.T) {
	graphCreateUrl := fmt.Sprintf("%s/v1/users/%s/graphs", baseUrl, username)

	ivGraphIdErr := errors.New("`graph create`: wrong arguments: " + validationErrorMessages["GraphId"])
	ivNumTypeErr := errors.New("`graph create`: wrong arguments: " + validationErrorMessages["UnitType"])
	ivColorErr := errors.New("`graph create`: wrong arguments: " + validationErrorMessages["Color"])
	ivSelfSufficientErr := errors.New("`graph create`: wrong arguments: " + validationErrorMessages["SelfSufficient"])
	respStatusErr := errors.New("`graph create`: http request failed: returns none success status code: 400")
	respDataErr := errors.New("`graph create`: http request failed: request failed: errorMessage")

	tests := []struct {
		name           string
		graphId        string
		graphName      string
		unit           string
		numType        string
		color          string
		timezone       string
		selfSufficient string
		statusCode     int
		response       *bytes.Buffer
		wantErr        error
	}{
		{"normal case w full option", graphId, graphName, graphUnit, numType, validColor, "Asia/Tokyo", "increment", http.StatusOK, bytes.NewBuffer(scResp), nil},
		{"normal case wo self sufficient", graphId, graphName, graphUnit, numType, validColor, "Asia/Tokyo", "", http.StatusOK, bytes.NewBuffer(scResp), nil},
		{"normal case wo timezone and self sufficient", graphId, graphName, graphUnit, numType, validColor, "", "", http.StatusOK, bytes.NewBuffer(scResp), nil},
		{"invalid graph id", "0000", graphName, graphUnit, numType, validColor, "", "", 0, nil, ivGraphIdErr},
		{"invalid number type", graphId, graphName, graphUnit, "string", validColor, "", "", 0, nil, ivNumTypeErr},
		{"invalid color", graphId, graphName, graphUnit, numType, "invalid color", "", "", 0, nil, ivColorErr},
		{"invalid self sufficient", graphId, graphName, graphUnit, numType, validColor, "", "invalid ss", 0, nil, ivSelfSufficientErr},
		{"invalid response status", graphId, graphName, graphUnit, numType, validColor, "", "", http.StatusBadRequest, nil, respStatusErr},
		{"invalid response data", graphId, graphName, graphUnit, numType, validColor, "", "", http.StatusOK, bytes.NewBuffer(errResp), respDataErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       ioutil.NopCloser(tt.response),
				Header:     make(http.Header),
			}

			c := NewTestClient(func(req *http.Request) *http.Response {
				if req.URL.String() != graphCreateUrl {
					t.Fatalf("want %#v, but got %#v", graphCreateUrl, req.URL.String())
				}

				if req.Header.Get(tokenHeader) != token {
					t.Fatalf("want %#v, but got %#v", token, req.Header.Get(tokenHeader))
				}

				return resp
			})

			// skip checking instance creation error
			pixela, _ := New(username, token, debug, OptionHTTPClient(c))

			_, err := pixela.CreateGraph(tt.graphId, tt.graphName, tt.unit, tt.numType, tt.color, tt.timezone, tt.selfSufficient)

			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but got %#v", tt.wantErr.Error(), err.Error())
				}
			}
		})
	}
}
