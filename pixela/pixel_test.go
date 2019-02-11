package pixela

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPixela_CreatePixel(t *testing.T) {
	pixelCreateUrl := fmt.Sprintf("%s/v1/users/%s/graphs/%s", baseUrl, username, graphId)

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

				if req.Header.Get(tokenHeader) != token {
					t.Fatalf("want %#v, but got %#v", token, req.Header.Get(tokenHeader))
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
	pixelGetUrl := fmt.Sprintf("%s/v1/users/%s/graphs/%s/%s", baseUrl, username, graphId, dateStr)

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

				if req.Header.Get(tokenHeader) != token {
					t.Fatalf("want %#v, but got %#v", token, req.Header.Get(tokenHeader))
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

func TestPixela_IncPixel(t *testing.T) {
	pixelIncUrl := fmt.Sprintf("%s/v1/users/%s/graphs/%s/increment", baseUrl, username, graphId)

	ivGraphIdErr := errors.New("`pixel inc`: wrong arguments: " + validationErrorMessages["GraphId"])
	respErr := errors.New("`pixel inc`: http request failed.: returns none success status code: 400")

	tests := []struct {
		name       string
		graphId    string
		statusCode int
		response   *bytes.Buffer
		wantErr    error
	}{
		{"normal case", graphId, http.StatusOK, bytes.NewBuffer(scResp), nil},
		{"invalid graphId", "0000", 0, nil, ivGraphIdErr},
		{"return error response", graphId, http.StatusBadRequest, bytes.NewBuffer(errResp), respErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       ioutil.NopCloser(tt.response),
				Header:     make(http.Header),
			}

			c := NewTestClient(func(req *http.Request) *http.Response {
				if req.URL.String() != pixelIncUrl {
					t.Fatalf("want %#v, but got %#v", pixelIncUrl, req.URL.String())
				}

				if req.Header.Get(tokenHeader) != token {
					t.Fatalf("want %#v, but got %#v", token, req.Header.Get(tokenHeader))
				}

				if req.Header.Get(contentLength) != contentZeroLen {
					t.Fatalf("want %#v, but got %#v", contentZeroLen, req.Header.Get(contentLength))
				}

				return resp
			})

			// skip checking instance creation error
			pixela, _ := New(username, token, debug, OptionHTTPClient(c))

			_, err := pixela.IncPixel(tt.graphId)

			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but got %#v", tt.wantErr.Error(), err.Error())
				}
			}
		})
	}
}

func TestPixela_DecPixel(t *testing.T) {
	pixelDecUrl := fmt.Sprintf("%s/v1/users/%s/graphs/%s/decrement", baseUrl, username, graphId)

	ivGraphIdErr := errors.New("`pixel dec`: wrong arguments: " + validationErrorMessages["GraphId"])
	respErr := errors.New("`pixel dec`: http request failed.: returns none success status code: 400")

	tests := []struct {
		name       string
		graphId    string
		statusCode int
		response   *bytes.Buffer
		wantErr    error
	}{
		{"normal case", graphId, http.StatusOK, bytes.NewBuffer(scResp), nil},
		{"invalid graphId", "0000", 0, nil, ivGraphIdErr},
		{"return error response", graphId, http.StatusBadRequest, bytes.NewBuffer(errResp), respErr},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: tt.statusCode,
				Body:       ioutil.NopCloser(tt.response),
				Header:     make(http.Header),
			}

			c := NewTestClient(func(req *http.Request) *http.Response {
				if req.URL.String() != pixelDecUrl {
					t.Fatalf("want %#v, but got %#v", pixelDecUrl, req.URL.String())
				}

				if req.Header.Get(tokenHeader) != token {
					t.Fatalf("want %#v, but got %#v", token, req.Header.Get(tokenHeader))
				}

				if req.Header.Get(contentLength) != contentZeroLen {
					t.Fatalf("want %#v, but got %#v", contentZeroLen, req.Header.Get(contentLength))
				}

				return resp
			})

			// skip checking instance creation error
			pixela, _ := New(username, token, debug, OptionHTTPClient(c))

			_, err := pixela.DecPixel(tt.graphId)

			if err != nil {
				if err.Error() != tt.wantErr.Error() {
					t.Fatalf("want %#v, but got %#v", tt.wantErr.Error(), err.Error())
				}
			}
		})
	}
}
