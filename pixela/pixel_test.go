package pixela

import (
	"bytes"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPixela_CreatePixel(t *testing.T) {
	ivGraphIdErr := errors.New("`pixel create`: wrong arguments: " + validationErrorMessages["GraphId"])
	ivDateErr := errors.New("`pixel create`: wrong arguments: " + validationErrorMessages["Date"])
	ivQuantityErr := errors.New("`pixel create`: wrong arguments: " + validationErrorMessages["Quantity"])
	ivOptionalDataErr := errors.New("`pixel create`: wrong arguments: " + validationErrorMessages["OptionalData"])
	respErr := errors.New("`pixel create`: http request failed.: returns none success status code: 400")

	tests := []struct {
		name         string
		graphId      string
		date         string
		quantity     string
		optionalData string
		statusCode   int
		response     *bytes.Buffer
		wantErr      error
	}{
		{"normal case wo optionalData", graphId, dateStr, quantityStr, "", http.StatusOK, bytes.NewBuffer(scResp), nil},
		{"normal case w optionalData", graphId, dateStr, quantityStr, `{"key": "value"}`, http.StatusOK, bytes.NewBuffer(scResp), nil},
		{"invalid graphId", "0000", dateStr, quantityStr, "", 0, nil, ivGraphIdErr},
		{"invalid date", graphId, "000A00", quantityStr, "", 0, nil, ivDateErr},
		{"invalid quantity", graphId, dateStr, "A", "", 0, nil, ivQuantityErr},
		{"invalid optionalData", graphId, dateStr, quantityStr, "A", 0, nil, ivOptionalDataErr},
		{"return error response", graphId, dateStr, quantityStr, "", http.StatusBadRequest, bytes.NewBuffer(errResp), respErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       ioutil.NopCloser(tt.response),
				Header:     make(http.Header),
			}

			c := NewTestClient(func(req *http.Request) *http.Response {
				if req.URL.String() != pixelCreateUrl {
					t.Fatalf("want %#v, but got %#v", pixelCreateUrl, req.URL.String())
				}

				return resp
			})

			// skip checking instance creation error
			pixela, _ := New(username, token, debug, OptionHTTPClient(c))

			_, err := pixela.CreatePixel(tt.graphId, tt.date, tt.quantity, tt.optionalData)

			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but got %#v", tt.wantErr.Error(), err.Error())
				}
			}
		})
	}
}

func TestPixela_GetPixel(t *testing.T) {
	ivGraphIdErr := errors.New("`pixel get`: wrong arguments: " + validationErrorMessages["GraphId"])
	ivDateErr := errors.New("`pixel get`: wrong arguments: " + validationErrorMessages["Date"])
	respErr := errors.New("`pixel get`: http request failed.: returns none success status code: 400")

	tests := []struct {
		name       string
		graphId    string
		date       string
		statusCode int
		response   *bytes.Buffer
		wantErr    error
	}{
		{"normal case wo optionalData", graphId, dateStr, http.StatusOK, bytes.NewBuffer(pixelRespWoOp), nil},
		{"normal case w optionalData", graphId, dateStr, http.StatusOK, bytes.NewBuffer(pixelRespWOp), nil},
		{"invalid graphId", "0000", dateStr, 0, nil, ivGraphIdErr},
		{"invalid date", graphId, "000A00", 0, nil, ivDateErr},
		{"return error response", graphId, dateStr, http.StatusBadRequest, bytes.NewBuffer(errResp), respErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       ioutil.NopCloser(tt.response),
				Header:     make(http.Header),
			}

			c := NewTestClient(func(req *http.Request) *http.Response {
				if req.URL.String() != pixelGetUrl {
					t.Fatalf("want %#v, but got %#v", pixelGetUrl, req.URL.String())
				}

				return resp
			})

			// skip checking instance creation error
			pixela, _ := New(username, token, debug, OptionHTTPClient(c))

			_, err := pixela.GetPixel(tt.graphId, tt.date)

			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but got %#v", tt.wantErr.Error(), err.Error())
				}
			}
		})
	}
}
