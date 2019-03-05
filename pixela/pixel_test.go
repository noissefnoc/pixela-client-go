package pixela

import (
	"fmt"
	"testing"
)

func TestPixela_CreatePixel(t *testing.T) {
	pixelCreateUrl := fmt.Sprintf("%s/v1/users/%s/graphs/%s", baseUrl, username, graphId)

	ivGraphIdErr := newCommandError(pixelCreate, "wrong arguments: "+validationErrorMessages["GraphId"])
	ivDateErr := newCommandError(pixelCreate, "wrong arguments: "+validationErrorMessages["Date"])
	ivQuantityErr := newCommandError(pixelCreate, "wrong arguments: "+validationErrorMessages["Quantity"])
	ivOptionalDataErr := newCommandError(pixelCreate, "wrong arguments: "+validationErrorMessages["OptionalData"])
	respDataErr := newCommandError(pixelCreate, "http request failed: post request failed: errorMessage")

	tests := testCases{
		{"normal case wo optionalData", sucStatus, scResp, nil, []string{graphId, dateStr, quantityStr, ""}},
		{"normal case w optionalData", sucStatus, scResp, nil, []string{graphId, dateStr, quantityStr, `{"key": "value"}`}},
		{"invalid graphId", 0, nil, ivGraphIdErr, []string{"0000", dateStr, quantityStr, ""}},
		{"invalid date", 0, nil, ivDateErr, []string{graphId, "000A00", quantityStr, ""}},
		{"invalid quantity", 0, nil, ivQuantityErr, []string{graphId, dateStr, "A", ""}},
		{"invalid optionalData", 0, nil, ivOptionalDataErr, []string{graphId, dateStr, quantityStr, "A"}},
		{"status error", errStatus, errResp, respDataErr, []string{graphId, dateStr, quantityStr, ""}},
	}

	subCommandTestHelper(t, pixelCreate, tests, pixelCreateUrl)
}

func TestPixela_GetPixel(t *testing.T) {
	pixelGetUrl := fmt.Sprintf("%s/v1/users/%s/graphs/%s/%s", baseUrl, username, graphId, dateStr)

	ivGraphIdErr := newCommandError(pixelGet, "wrong arguments: "+validationErrorMessages["GraphId"])
	ivDateErr := newCommandError(pixelGet, "wrong arguments: "+validationErrorMessages["Date"])
	respDataErr := newCommandError(pixelGet, "http request failed: get request failed: errorMessage")

	tests := testCases{
		{"normal case wo optionalData", sucStatus, pixelRespWoOp, nil, []string{graphId, dateStr}},
		{"normal case w optionalData", sucStatus, pixelRespWOp, nil, []string{graphId, dateStr}},
		{"invalid graphId", 0, nil, ivGraphIdErr, []string{"0000", dateStr}},
		{"invalid date", 0, nil, ivDateErr, []string{graphId, "000A00"}},
		{"status error", errStatus, errResp, respDataErr, []string{graphId, dateStr}},
	}

	subCommandTestHelper(t, pixelGet, tests, pixelGetUrl)
}

func TestPixela_IncPixel(t *testing.T) {
	pixelIncUrl := fmt.Sprintf("%s/v1/users/%s/graphs/%s/increment", baseUrl, username, graphId)

	ivGraphIdErr := newCommandError(pixelInc, "wrong arguments: "+validationErrorMessages["GraphId"])
	respDataErr := newCommandError(pixelInc, "http request failed: put request failed: errorMessage")

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{graphId}},
		{"invalid graphId", 0, nil, ivGraphIdErr, []string{"0000"}},
		{"status error", errStatus, errResp, respDataErr, []string{graphId}},
	}

	subCommandTestHelper(t, pixelInc, tests, pixelIncUrl)
}

func TestPixela_DecPixel(t *testing.T) {
	pixelDecUrl := fmt.Sprintf("%s/v1/users/%s/graphs/%s/decrement", baseUrl, username, graphId)

	ivGraphIdErr := newCommandError(pixelDec, "wrong arguments: "+validationErrorMessages["GraphId"])
	respDataErr := newCommandError(pixelDec, "http request failed: put request failed: errorMessage")

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{graphId}},
		{"invalid graphId", 0, nil, ivGraphIdErr, []string{"0000"}},
		{"status error", errStatus, errResp, respDataErr, []string{graphId}},
	}

	subCommandTestHelper(t, pixelDec, tests, pixelDecUrl)
}

func TestPixela_DeletePixel(t *testing.T) {
	pixelDeleteUrl := fmt.Sprintf("%s/v1/users/%s/graphs/%s/%s", baseUrl, username, graphId, dateStr)

	ivGraphIdErr := newCommandError(pixelDelete, "wrong arguments: "+validationErrorMessages["GraphId"])
	ivDateErr := newCommandError(pixelDelete, "wrong arguments: "+validationErrorMessages["Date"])
	respDataErr := newCommandError(pixelDelete, "http request failed: delete request failed: errorMessage")

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{graphId, dateStr}},
		{"invalid graphId", 0, nil, ivGraphIdErr, []string{"0000", dateStr}},
		{"invalid date", 0, nil, ivDateErr, []string{graphId, "000A00"}},
		{"status error", errStatus, errResp, respDataErr, []string{graphId, dateStr}},
	}

	subCommandTestHelper(t, pixelDelete, tests, pixelDeleteUrl)
}

func TestPixela_UpdatePixel(t *testing.T) {
	pixelUpdateUrl := fmt.Sprintf("%s/v1/users/%s/graphs/%s/%s", baseUrl, username, graphId, dateStr)

	ivGraphIdErr := newCommandError(pixelUpdate, "wrong arguments: "+validationErrorMessages["GraphId"])
	ivDateErr := newCommandError(pixelUpdate, "wrong arguments: "+validationErrorMessages["Date"])
	ivQuantityErr := newCommandError(pixelUpdate, "wrong arguments: "+validationErrorMessages["Quantity"])
	ivOptionalDataErr := newCommandError(pixelUpdate, "wrong arguments: "+validationErrorMessages["OptionalData"])
	respDataErr := newCommandError(pixelUpdate, "http request failed: put request failed: errorMessage")

	tests := testCases{
		{"normal case wo optionalData", sucStatus, scResp, nil, []string{graphId, dateStr, quantityStr, ""}},
		{"normal case w optionalData", sucStatus, scResp, nil, []string{graphId, dateStr, quantityStr, `{"key": "value"}`}},
		{"invalid graphId", 0, nil, ivGraphIdErr, []string{"0000", dateStr, quantityStr, ""}},
		{"invalid date", 0, nil, ivDateErr, []string{graphId, "000A00", quantityStr, ""}},
		{"invalid quantity", 0, nil, ivQuantityErr, []string{graphId, dateStr, "A", ""}},
		{"invalid optionalData", 0, nil, ivOptionalDataErr, []string{graphId, dateStr, quantityStr, "A"}},
		{"status error", errStatus, errResp, respDataErr, []string{graphId, dateStr, quantityStr, ""}},
	}

	subCommandTestHelper(t, pixelUpdate, tests, pixelUpdateUrl)
}
