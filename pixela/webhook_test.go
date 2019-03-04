package pixela

import (
	"fmt"
	"testing"
)

func TestPixela_CreateWebhook(t *testing.T) {
	webhookCreateUrl := fmt.Sprintf("%s/v1/users/%s/webhooks", baseUrl, username)

	ivGraphIdErr := newCommandError(webhookCreate, "wrong arguments: "+validationErrorMessages["GraphId"])
	ivWebhookTypeErr := newCommandError(webhookCreate, "wrong arguments: "+validationErrorMessages["WebhookType"])
	respErr := newCommandError(webhookCreate, fmt.Sprintf("http request failed: returns none success status code: %d", errStatus))

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{graphId, "increment"}},
		{"invalid graph id", 0, nil, ivGraphIdErr, []string{"0000", "increment"}},
		{"invalid webhook type", 0, nil, ivWebhookTypeErr, []string{graphId, "hoge"}},
		{"invalid status", errStatus, errResp, respErr, []string{graphId, "increment"}},
	}

	subCommandTestHelper(t, webhookCreate, tests, webhookCreateUrl)
}

func TestPixela_GetWebhookDefinitions(t *testing.T) {
	webhookGetUrl := fmt.Sprintf("%s/v1/users/%s/webhooks", baseUrl, username)

	respErr := newCommandError(webhookGet, fmt.Sprintf("http request failed: returns none success status code: %d", errStatus))

	tests := testCases{
		{"normal case", sucStatus, webhookResp, nil, nil},
		{"invalid status", errStatus, errResp, respErr, nil},
	}

	subCommandTestHelper(t, webhookGet, tests, webhookGetUrl)
}

func TestPixela_InvokeWebhooks(t *testing.T) {
	webhookInvokeUrl := fmt.Sprintf("%s/v1/users/%s/webhooks/%s", baseUrl, username, webhookHash)

	respErr := newCommandError(webhookInvoke, fmt.Sprintf("http request failed: returns none success status code: %d", errStatus))

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{webhookHash}},
		{"invalid status", errStatus, errResp, respErr, []string{webhookHash}},
	}

	subCommandTestHelper(t, webhookInvoke, tests, webhookInvokeUrl)
}

func TestPixela_DeleteWebhook(t *testing.T) {
	webhookDeleteUrl := fmt.Sprintf("%s/v1/users/%s/webhooks/%s", baseUrl, username, webhookHash)

	respErr := newCommandError(webhookDelete, fmt.Sprintf("http request failed: returns none success status code: %d", errStatus))

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{webhookHash}},
		{"invalid status", errStatus, errResp, respErr, []string{webhookHash}},
	}

	subCommandTestHelper(t, webhookDelete, tests, webhookDeleteUrl)
}
