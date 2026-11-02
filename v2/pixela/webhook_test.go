package pixela

import (
	"context"
	"net/http"
	"testing"
)

func TestClient_Webhook(t *testing.T) {
	t.Run("CreateWebhook", func(t *testing.T) {
		client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/v1/users/testuser/webhooks" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true,"webhookHash":"test-hash"}`))
		})
		defer server.Close()

		res, err := client.CreateWebhook(context.Background(), CreateWebhookParams{
			GraphID: "test-graph",
			Type:    "increment",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !res.IsSuccess || res.WebhookHash != "test-hash" {
			t.Errorf("unexpected result: %+v", res)
		}
	})

	t.Run("GetWebhooks", func(t *testing.T) {
		client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"webhooks":[{"webhookHash":"test-hash","graphID":"test-graph","type":"increment"}]}`))
		})
		defer server.Close()

		defs, err := client.GetWebhooks(context.Background())
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if len(defs.Webhooks) != 1 || defs.Webhooks[0].WebhookHash != "test-hash" {
			t.Errorf("unexpected webhooks: %+v", defs)
		}
	})

	t.Run("InvokeWebhook", func(t *testing.T) {
		client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/v1/users/testuser/webhooks/test-hash" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}
			if r.Header.Get("X-USER-TOKEN") != "" {
				t.Errorf("invoke webhook should not send X-USER-TOKEN")
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		})
		defer server.Close()

		res, err := client.InvokeWebhook(context.Background(), "test-hash")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !res.IsSuccess {
			t.Errorf("unexpected result: %+v", res)
		}
	})

	t.Run("DeleteWebhook", func(t *testing.T) {
		client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/v1/users/testuser/webhooks/test-hash" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}
			if r.Header.Get("X-USER-TOKEN") != "testtoken" {
				t.Errorf("expected token header")
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		})
		defer server.Close()

		res, err := client.DeleteWebhook(context.Background(), "test-hash")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !res.IsSuccess {
			t.Errorf("unexpected result: %+v", res)
		}
	})

	t.Run("validation errors", func(t *testing.T) {
		client, server := newTestClient(t, nil)
		defer server.Close()

		if _, err := client.CreateWebhook(context.Background(), CreateWebhookParams{GraphID: "Invalid"}); err == nil {
			t.Error("expected error for invalid graphID")
		}
		if _, err := client.CreateWebhook(context.Background(), CreateWebhookParams{GraphID: "test-graph", Type: "invalid"}); err == nil {
			t.Error("expected error for invalid type")
		}
	})
}
