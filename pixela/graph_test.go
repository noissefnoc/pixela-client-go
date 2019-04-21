package pixela

import (
	"fmt"
	"testing"
)

func TestPixela_CreateGraph(t *testing.T) {
	graphCreateURL := fmt.Sprintf("%s/v1/users/%s/graphs", baseURL, username)

	ivGraphIDErr := newCommandError(graphCreate, "wrong arguments: "+validationErrorMessages["GraphID"])
	ivNumTypeErr := newCommandError(graphCreate, "wrong arguments: "+validationErrorMessages["UnitType"])
	ivColorErr := newCommandError(graphCreate, "wrong arguments: "+validationErrorMessages["Color"])
	ivSelfSufficientErr := newCommandError(graphCreate, "wrong arguments: "+validationErrorMessages["SelfSufficient"])
	respDataErr := newCommandError(graphCreate, "http request failed: post request failed: errorMessage")

	tests := testCases{
		{"normal case w full option", sucStatus, scResp, nil, []string{graphID, graphName, graphUnit, numType, validColor, "Asia/Tokyo", "increment"}},
		{"normal case wo self sufficient", sucStatus, scResp, nil, []string{graphID, graphName, graphUnit, numType, validColor, "Asia/Tokyo", ""}},
		{"normal case wo timezone and self sufficient", sucStatus, scResp, nil, []string{graphID, graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid graph id", 0, nil, ivGraphIDErr, []string{"0000", graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid number type", 0, nil, ivNumTypeErr, []string{graphID, graphName, graphUnit, "string", validColor, "", ""}},
		{"invalid color", 0, nil, ivColorErr, []string{graphID, graphName, graphUnit, numType, "invalid color", "", ""}},
		{"invalid self sufficient", 0, nil, ivSelfSufficientErr, []string{graphID, graphName, graphUnit, numType, validColor, "", "invalid ss"}},
		{"invalid response status", errStatus, errResp, respDataErr, []string{graphID, graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid response data", sucStatus, errResp, respDataErr, []string{graphID, graphName, graphUnit, numType, validColor, "", ""}},
	}

	subCommandTestHelper(t, graphCreate, tests, graphCreateURL)
}

func TestPixela_UpdateGraph(t *testing.T) {
	graphUpdateURL := fmt.Sprintf("%s/v1/users/%s/graphs/%s", baseURL, username, graphID)

	ivGraphIDErr := newCommandError(graphUpdate, "wrong arguments: "+validationErrorMessages["GraphID"])
	respDataErr := newCommandError(graphUpdate, "http request failed: put request failed: errorMessage")

	tests := testCases{
		{"normal case w full option", sucStatus, scResp, nil, []string{graphID, graphName, graphUnit, numType, validColor, "Asia/Tokyo", "increment"}},
		{"normal case wo self sufficient", sucStatus, scResp, nil, []string{graphID, graphName, graphUnit, numType, validColor, "Asia/Tokyo", ""}},
		{"normal case wo timezone and self sufficient", sucStatus, scResp, nil, []string{graphID, graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid graph id", 0, nil, ivGraphIDErr, []string{"0000", graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid response status", errStatus, errResp, respDataErr, []string{graphID, graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid response data", sucStatus, errResp, respDataErr, []string{graphID, graphName, graphUnit, numType, validColor, "", ""}},
	}

	subCommandTestHelper(t, graphUpdate, tests, graphUpdateURL)
}

func TestPixela_DeleteGraph(t *testing.T) {
	graphDeleteURL := fmt.Sprintf("%s/v1/users/%s/graphs/%s", baseURL, username, graphID)

	ivGraphIDErr := newCommandError(graphDelete, "wrong arguments: "+validationErrorMessages["GraphID"])
	respDataErr := newCommandError(graphDelete, "http request failed: delete request failed: errorMessage")

	tests := testCases{
		{"normal case w full option", sucStatus, scResp, nil, []string{graphID, graphName, graphUnit, numType, validColor, "Asia/Tokyo", "increment"}},
		{"normal case wo self sufficient", sucStatus, scResp, nil, []string{graphID, graphName, graphUnit, numType, validColor, "Asia/Tokyo", ""}},
		{"normal case wo timezone and self sufficient", sucStatus, scResp, nil, []string{graphID, graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid graph id", 0, nil, ivGraphIDErr, []string{"0000", graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid response status", errStatus, errResp, respDataErr, []string{graphID, graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid response data", sucStatus, errResp, respDataErr, []string{graphID, graphName, graphUnit, numType, validColor, "", ""}},
	}

	subCommandTestHelper(t, graphDelete, tests, graphDeleteURL)
}

func TestPixela_GetGraphSvg(t *testing.T) {
	graphSvgURL := fmt.Sprintf("%s/v1/users/%s/graphs/%s", baseURL, username, graphID)
	respDataErr := newCommandError(graphSvg, "http request failed: get request failed: errorMessage")

	tests := testCases{
		{"normal case w full option", sucStatus, scResp, nil, []string{graphID, dateStr, "short"}},
		{"normal case wo date", sucStatus, scResp, nil, []string{graphID, "", "short"}},
		{"normal case wo mode", sucStatus, scResp, nil, []string{graphID, dateStr, ""}},
		{"invalid response status", errStatus, errResp, respDataErr, []string{graphID, dateStr, ""}},
		{"invalid response data", sucStatus, errResp, respDataErr, []string{graphID, dateStr, ""}},
	}

	subCommandTestHelper(t, graphSvg, tests, graphSvgURL)
}

func TestPixela_GetGraphDefinition(t *testing.T) {
	graphDefURL := fmt.Sprintf("%s/v1/users/%s/graphs", baseURL, username)

	respDataErr := newCommandError(graphGet, "http request failed: get request failed: errorMessage")

	tests := testCases{
		{"normal case w full option", sucStatus, graphDefResp, nil, nil},
		{"normal case wo mode", sucStatus, graphDefResp, nil, nil},
		{"invalid response status", errStatus, errResp, respDataErr, nil},
		{"invalid response data", sucStatus, errResp, respDataErr, nil},
	}

	subCommandTestHelper(t, graphGet, tests, graphDefURL)
}

func TestPixela_GetGraphPixelsDateList(t *testing.T) {
	graphGetPixelsDateURL := fmt.Sprintf("%s/v1/users/%s/graphs/%s/pixels", baseURL, username, graphID)

	respDataErr := newCommandError(graphPixels, "http request failed: get request failed: errorMessage")

	tests := testCases{
		{"normal case w full option", sucStatus, graphPixelsResp, nil, []string{graphID, "20190101", "20190102"}},
		{"normal case wo from", sucStatus, graphPixelsResp, nil, []string{graphID, "", "20190102"}},
		{"normal case wo to", sucStatus, graphPixelsResp, nil, []string{graphID, "20190101", ""}},
		{"normal case wo from and to", sucStatus, graphPixelsResp, nil, []string{graphID, "", ""}},
		{"invalid response status", errStatus, errResp, respDataErr, []string{graphID, "", ""}},
		{"invalid response data", sucStatus, errResp, respDataErr, []string{graphID, "", ""}},
	}

	subCommandTestHelper(t, graphPixels, tests, graphGetPixelsDateURL)
}

func TestPixela_GetGraphDetailURL(t *testing.T) {
	want := fmt.Sprintf("%s/v1/users/%s/graphs/%s.html", baseURL, username, graphID)

	// skip checking instance creation error
	pixela, _ := New(username, token, false)
	got := pixela.GetGraphDetailURL(graphID)

	if got != want {
		t.Fatalf("want %#v, but got %#v", want, got)
	}
}

func TestPixela_GetGraphStat(t *testing.T) {
	graphStatURL := fmt.Sprintf("%s/v1/users/%s/graphs/%s/stats", baseURL, username, graphID)

	respDataErr := newCommandError(graphStat, "http request failed: get request failed: errorMessage")

	tests := testCases{
		{"normal case", sucStatus, graphStatResp, nil, []string{graphID}},
		{"invalid response status", errStatus, errResp, respDataErr, []string{graphID}},
		{"invalid response data", sucStatus, errResp, respDataErr, []string{graphID}},
	}

	subCommandTestHelper(t, graphStat, tests, graphStatURL)
}