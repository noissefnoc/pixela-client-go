package pixela

import (
	"fmt"
	"testing"
)

func TestPixela_CreateWebhook(t *testing.T) {
	webhookCreateUrl := fmt.Sprintf("%s/v1/users/%s/webhooks", baseUrl, username)

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{graphId, "increment"}},
	}

	subCommandTestHelper(t, webhookCreate, tests, webhookCreateUrl)
}

func TestPixela_GetWebhookDefinitions(t *testing.T) {
	webhookGetUrl := fmt.Sprintf("%s/v1/users/%s/webhooks", baseUrl, username)

	tests := testCases{
		{"normal case", sucStatus, webhookResp, nil, nil},
	}

	subCommandTestHelper(t, webhookGet, tests, webhookGetUrl)
}

func TestPixela_InvokeWebhooks(t *testing.T) {
	webhookInvokeUrl := fmt.Sprintf("%s/v1/users/%s/webhooks/%s", baseUrl, username, webhookHash)

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, []string{webhookHash}},
	}

	subCommandTestHelper(t, webhookInvoke, tests, webhookInvokeUrl)
}

func TestPixela_DeleteWebhook(t *testing.T) {
	webhookDeleteUrl := fmt.Sprintf("%s/v1/users/%s/webhooks/%s", baseUrl, username, webhookHash)

	tests := testCases{
		{"normal case", sucStatus, scResp, nil, nil},
	}

	subCommandTestHelper(t, webhookDelete, tests, webhookDeleteUrl)
}
