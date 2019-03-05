package pixela

import (
	"fmt"
	"testing"
)

func TestPixela_CreateGraph(t *testing.T) {
	graphCreateUrl := fmt.Sprintf("%s/v1/users/%s/graphs", baseUrl, username)

	ivGraphIdErr := newCommandError(graphCreate, "wrong arguments: "+validationErrorMessages["GraphId"])
	ivNumTypeErr := newCommandError(graphCreate, "wrong arguments: "+validationErrorMessages["UnitType"])
	ivColorErr := newCommandError(graphCreate, "wrong arguments: "+validationErrorMessages["Color"])
	ivSelfSufficientErr := newCommandError(graphCreate, "wrong arguments: "+validationErrorMessages["SelfSufficient"])
	respDataErr := newCommandError(graphCreate, "http request failed: post request failed: errorMessage")

	tests := testCases{
		{"normal case w full option", sucStatus, scResp, nil, []string{graphId, graphName, graphUnit, numType, validColor, "Asia/Tokyo", "increment"}},
		{"normal case wo self sufficient", sucStatus, scResp, nil, []string{graphId, graphName, graphUnit, numType, validColor, "Asia/Tokyo", ""}},
		{"normal case wo timezone and self sufficient", sucStatus, scResp, nil, []string{graphId, graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid graph id", 0, nil, ivGraphIdErr, []string{"0000", graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid number type", 0, nil, ivNumTypeErr, []string{graphId, graphName, graphUnit, "string", validColor, "", ""}},
		{"invalid color", 0, nil, ivColorErr, []string{graphId, graphName, graphUnit, numType, "invalid color", "", ""}},
		{"invalid self sufficient", 0, nil, ivSelfSufficientErr, []string{graphId, graphName, graphUnit, numType, validColor, "", "invalid ss"}},
		{"invalid response status", errStatus, errResp, respDataErr, []string{graphId, graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid response data", sucStatus, errResp, respDataErr, []string{graphId, graphName, graphUnit, numType, validColor, "", ""}},
	}

	subCommandTestHelper(t, graphCreate, tests, graphCreateUrl)
}

func TestPixela_UpdateGraph(t *testing.T) {
	graphUpdateUrl := fmt.Sprintf("%s/v1/users/%s/graphs/%s", baseUrl, username, graphId)

	ivGraphIdErr := newCommandError(graphUpdate, "wrong arguments: "+validationErrorMessages["GraphId"])
	respDataErr := newCommandError(graphUpdate, "http request failed: put request failed: errorMessage")

	tests := testCases{
		{"normal case w full option", sucStatus, scResp, nil, []string{graphId, graphName, graphUnit, numType, validColor, "Asia/Tokyo", "increment"}},
		{"normal case wo self sufficient", sucStatus, scResp, nil, []string{graphId, graphName, graphUnit, numType, validColor, "Asia/Tokyo", ""}},
		{"normal case wo timezone and self sufficient", sucStatus, scResp, nil, []string{graphId, graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid graph id", 0, nil, ivGraphIdErr, []string{"0000", graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid response status", errStatus, errResp, respDataErr, []string{graphId, graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid response data", sucStatus, errResp, respDataErr, []string{graphId, graphName, graphUnit, numType, validColor, "", ""}},
	}

	subCommandTestHelper(t, graphUpdate, tests, graphUpdateUrl)
}

func TestPixela_DeleteGraph(t *testing.T) {
	graphDeleteUrl := fmt.Sprintf("%s/v1/users/%s/graphs/%s", baseUrl, username, graphId)

	ivGraphIdErr := newCommandError(graphDelete, "wrong arguments: "+validationErrorMessages["GraphId"])
	respDataErr := newCommandError(graphDelete, "http request failed: delete request failed: errorMessage")

	tests := testCases{
		{"normal case w full option", sucStatus, scResp, nil, []string{graphId, graphName, graphUnit, numType, validColor, "Asia/Tokyo", "increment"}},
		{"normal case wo self sufficient", sucStatus, scResp, nil, []string{graphId, graphName, graphUnit, numType, validColor, "Asia/Tokyo", ""}},
		{"normal case wo timezone and self sufficient", sucStatus, scResp, nil, []string{graphId, graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid graph id", 0, nil, ivGraphIdErr, []string{"0000", graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid response status", errStatus, errResp, respDataErr, []string{graphId, graphName, graphUnit, numType, validColor, "", ""}},
		{"invalid response data", sucStatus, errResp, respDataErr, []string{graphId, graphName, graphUnit, numType, validColor, "", ""}},
	}

	subCommandTestHelper(t, graphDelete, tests, graphDeleteUrl)
}

func TestPixela_GetGraphSvg(t *testing.T) {
	graphSvgUrl := fmt.Sprintf("%s/v1/users/%s/graphs/%s", baseUrl, username, graphId)
	respDataErr := newCommandError(graphSvg, "http request failed: get request failed: errorMessage")

	tests := testCases{
		{"normal case w full option", sucStatus, scResp, nil, []string{graphId, dateStr, "short"}},
		{"normal case wo date", sucStatus, scResp, nil, []string{graphId, "", "short"}},
		{"normal case wo mode", sucStatus, scResp, nil, []string{graphId, dateStr, ""}},
		{"invalid response status", errStatus, errResp, respDataErr, []string{graphId, dateStr, ""}},
		{"invalid response data", sucStatus, errResp, respDataErr, []string{graphId, dateStr, ""}},
	}

	subCommandTestHelper(t, graphSvg, tests, graphSvgUrl)
}

func TestPixela_GetGraphDefinition(t *testing.T) {
	graphDefUrl := fmt.Sprintf("%s/v1/users/%s/graphs", baseUrl, username)

	respDataErr := newCommandError(graphDef, "http request failed: get request failed: errorMessage")

	tests := testCases{
		{"normal case w full option", sucStatus, graphDefResp, nil, nil},
		{"normal case wo mode", sucStatus, graphDefResp, nil, nil},
		{"invalid response status", errStatus, errResp, respDataErr, nil},
		{"invalid response data", sucStatus, errResp, respDataErr, nil},
	}

	subCommandTestHelper(t, graphDef, tests, graphDefUrl)
}
