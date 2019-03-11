package pixela

import (
	"fmt"
	"testing"
)

func TestPixela_CreateWebhook(t *testing.T) {
	webhookCreateURL := fmt.Sprintf("%s/v1/users/%s/webhooks", baseURL, username)

	ivGraphIDErr := newCommandError(webhookCreate, "wrong arguments: "+validationErrorMessages["GraphID"])
	ivWebhookTypeErr := newCommandError(webhookCreate, "wrong arguments: "+validationErrorMessages["WebhookType"])
	respDataErr := newCommandError(webhookCreate, "http request failed: post request failed: errorMessage")

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{graphID, "increment"}},
		{"invalid graph id", 0, nil, ivGraphIDErr, []string{"0000", "increment"}},
		{"invalid webhook type", 0, nil, ivWebhookTypeErr, []string{graphID, "hoge"}},
		{"invalid status", errStatus, errResp, respDataErr, []string{graphID, "increment"}},
	}

	subCommandTestHelper(t, webhookCreate, tests, webhookCreateURL)
}

func TestPixela_GetWebhookDefinitions(t *testing.T) {
	webhookGetURL := fmt.Sprintf("%s/v1/users/%s/webhooks", baseURL, username)

	respDataErr := newCommandError(webhookGet, "http request failed: get request failed: errorMessage")

	tests := testCases{
		{"normal case", sucStatus, webhookResp, nil, nil},
		{"invalid status", errStatus, errResp, respDataErr, nil},
	}

	subCommandTestHelper(t, webhookGet, tests, webhookGetURL)
}

func TestPixela_InvokeWebhooks(t *testing.T) {
	webhookInvokeURL := fmt.Sprintf("%s/v1/users/%s/webhooks/%s", baseURL, username, webhookHash)

	respDataErr := newCommandError(webhookInvoke, "http request failed: post request failed: errorMessage")

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{webhookHash}},
		{"invalid status", errStatus, errResp, respDataErr, []string{webhookHash}},
	}

	subCommandTestHelper(t, webhookInvoke, tests, webhookInvokeURL)
}

func TestPixela_DeleteWebhook(t *testing.T) {
	webhookDeleteURL := fmt.Sprintf("%s/v1/users/%s/webhooks/%s", baseURL, username, webhookHash)

	respDataErr := newCommandError(webhookDelete, "http request failed: delete request failed: errorMessage")

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{webhookHash}},
		{"invalid status", errStatus, errResp, respDataErr, []string{webhookHash}},
	}

	subCommandTestHelper(t, webhookDelete, tests, webhookDeleteURL)
}
