package pixela

import (
	"fmt"
	"testing"
)

func TestPixela_CreatePixel(t *testing.T) {
	pixelCreateURL := fmt.Sprintf("%s/v1/users/%s/graphs/%s", baseURL, username, graphID)

	ivGraphIDErr := newCommandError(pixelCreate, "wrong arguments: "+validationErrorMessages["GraphID"])
	ivDateErr := newCommandError(pixelCreate, "wrong arguments: "+validationErrorMessages["Date"])
	ivQuantityErr := newCommandError(pixelCreate, "wrong arguments: "+validationErrorMessages["Quantity"])
	ivOptionalDataErr := newCommandError(pixelCreate, "wrong arguments: "+validationErrorMessages["OptionalData"])
	respDataErr := newCommandError(pixelCreate, "http request failed: post request failed: errorMessage")

	tests := testCases{
		{"normal case wo optionalData", sucStatus, scResp, nil, []string{graphID, dateStr, quantityStr, ""}},
		{"normal case w optionalData", sucStatus, scResp, nil, []string{graphID, dateStr, quantityStr, `{"key": "value"}`}},
		{"invalid graphID", 0, nil, ivGraphIDErr, []string{"0000", dateStr, quantityStr, ""}},
		{"invalid date", 0, nil, ivDateErr, []string{graphID, "000A00", quantityStr, ""}},
		{"invalid quantity", 0, nil, ivQuantityErr, []string{graphID, dateStr, "A", ""}},
		{"invalid optionalData", 0, nil, ivOptionalDataErr, []string{graphID, dateStr, quantityStr, "A"}},
		{"status error", errStatus, errResp, respDataErr, []string{graphID, dateStr, quantityStr, ""}},
	}

	subCommandTestHelper(t, pixelCreate, tests, pixelCreateURL)
}

func TestPixela_GetPixel(t *testing.T) {
	pixelGetURL := fmt.Sprintf("%s/v1/users/%s/graphs/%s/%s", baseURL, username, graphID, dateStr)

	ivGraphIDErr := newCommandError(pixelGet, "wrong arguments: "+validationErrorMessages["GraphID"])
	ivDateErr := newCommandError(pixelGet, "wrong arguments: "+validationErrorMessages["Date"])
	respDataErr := newCommandError(pixelGet, "http request failed: get request failed: errorMessage")

	tests := testCases{
		{"normal case wo optionalData", sucStatus, pixelRespWoOp, nil, []string{graphID, dateStr}},
		{"normal case w optionalData", sucStatus, pixelRespWOp, nil, []string{graphID, dateStr}},
		{"invalid graphID", 0, nil, ivGraphIDErr, []string{"0000", dateStr}},
		{"invalid date", 0, nil, ivDateErr, []string{graphID, "000A00"}},
		{"status error", errStatus, errResp, respDataErr, []string{graphID, dateStr}},
	}

	subCommandTestHelper(t, pixelGet, tests, pixelGetURL)
}

func TestPixela_IncPixel(t *testing.T) {
	pixelIncURL := fmt.Sprintf("%s/v1/users/%s/graphs/%s/increment", baseURL, username, graphID)

	ivGraphIDErr := newCommandError(pixelInc, "wrong arguments: "+validationErrorMessages["GraphID"])
	respDataErr := newCommandError(pixelInc, "http request failed: put request failed: errorMessage")

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{graphID}},
		{"invalid graphID", 0, nil, ivGraphIDErr, []string{"0000"}},
		{"status error", errStatus, errResp, respDataErr, []string{graphID}},
	}

	subCommandTestHelper(t, pixelInc, tests, pixelIncURL)
}

func TestPixela_DecPixel(t *testing.T) {
	pixelDecURL := fmt.Sprintf("%s/v1/users/%s/graphs/%s/decrement", baseURL, username, graphID)

	ivGraphIDErr := newCommandError(pixelDec, "wrong arguments: "+validationErrorMessages["GraphID"])
	respDataErr := newCommandError(pixelDec, "http request failed: put request failed: errorMessage")

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{graphID}},
		{"invalid graphID", 0, nil, ivGraphIDErr, []string{"0000"}},
		{"status error", errStatus, errResp, respDataErr, []string{graphID}},
	}

	subCommandTestHelper(t, pixelDec, tests, pixelDecURL)
}

func TestPixela_DeletePixel(t *testing.T) {
	pixelDeleteURL := fmt.Sprintf("%s/v1/users/%s/graphs/%s/%s", baseURL, username, graphID, dateStr)

	ivGraphIDErr := newCommandError(pixelDelete, "wrong arguments: "+validationErrorMessages["GraphID"])
	ivDateErr := newCommandError(pixelDelete, "wrong arguments: "+validationErrorMessages["Date"])
	respDataErr := newCommandError(pixelDelete, "http request failed: delete request failed: errorMessage")

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{graphID, dateStr}},
		{"invalid graphID", 0, nil, ivGraphIDErr, []string{"0000", dateStr}},
		{"invalid date", 0, nil, ivDateErr, []string{graphID, "000A00"}},
		{"status error", errStatus, errResp, respDataErr, []string{graphID, dateStr}},
	}

	subCommandTestHelper(t, pixelDelete, tests, pixelDeleteURL)
}

func TestPixela_UpdatePixel(t *testing.T) {
	pixelUpdateURL := fmt.Sprintf("%s/v1/users/%s/graphs/%s/%s", baseURL, username, graphID, dateStr)

	ivGraphIDErr := newCommandError(pixelUpdate, "wrong arguments: "+validationErrorMessages["GraphID"])
	ivDateErr := newCommandError(pixelUpdate, "wrong arguments: "+validationErrorMessages["Date"])
	ivQuantityErr := newCommandError(pixelUpdate, "wrong arguments: "+validationErrorMessages["Quantity"])
	ivOptionalDataErr := newCommandError(pixelUpdate, "wrong arguments: "+validationErrorMessages["OptionalData"])
	respDataErr := newCommandError(pixelUpdate, "http request failed: put request failed: errorMessage")

	tests := testCases{
		{"normal case wo optionalData", sucStatus, scResp, nil, []string{graphID, dateStr, quantityStr, ""}},
		{"normal case w optionalData", sucStatus, scResp, nil, []string{graphID, dateStr, quantityStr, `{"key": "value"}`}},
		{"invalid graphID", 0, nil, ivGraphIDErr, []string{"0000", dateStr, quantityStr, ""}},
		{"invalid date", 0, nil, ivDateErr, []string{graphID, "000A00", quantityStr, ""}},
		{"invalid quantity", 0, nil, ivQuantityErr, []string{graphID, dateStr, "A", ""}},
		{"invalid optionalData", 0, nil, ivOptionalDataErr, []string{graphID, dateStr, quantityStr, "A"}},
		{"status error", errStatus, errResp, respDataErr, []string{graphID, dateStr, quantityStr, ""}},
	}

	subCommandTestHelper(t, pixelUpdate, tests, pixelUpdateURL)
}
