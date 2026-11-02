package cmd

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWebhookCommands(t *testing.T) {
	ctx := context.Background()

	t.Run("webhook help", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"webhook", "help"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d", code)
		}
		if !strings.Contains(out.String(), "Usage:") {
			t.Errorf("expected usage in out, got: %s", out.String())
		}
	})

	t.Run("webhook unknown command", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"webhook", "unknown"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
		if !strings.Contains(err.String(), "unknown webhook command") {
			t.Errorf("expected error in err, got: %s", err.String())
		}
	})

	t.Run("webhook create argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"webhook", "create", "g1"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("webhook get argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"webhook", "get", "extra"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("webhook invoke argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"webhook", "invoke"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("webhook delete argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"webhook", "delete"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("webhook create success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" || r.URL.Path != "/v1/users/testuser/webhooks" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true,"webhookHash":"hash123"}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "-n", "webhook", "create", "test-graph", "increment"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"webhookHash":"hash123"`) {
			t.Errorf("expected webhookHash in out, got: %s", out.String())
		}
	})

	t.Run("webhook get success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" || r.URL.Path != "/v1/users/testuser/webhooks" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"webhooks":[{"webhookHash":"hash123","graphID":"test-graph","type":"increment"}]}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "webhook", "get"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"hash123"`) {
			t.Errorf("expected webhooks JSON in out, got: %s", out.String())
		}
	})

	t.Run("webhook invoke success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" || r.URL.Path != "/v1/users/testuser/webhooks/hash123" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "-n", "webhook", "invoke", "hash123"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"isSuccess":true`) {
			t.Errorf("expected JSON in out, got: %s", out.String())
		}
	})

	t.Run("webhook delete success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "DELETE" || r.URL.Path != "/v1/users/testuser/webhooks/hash123" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "-n", "webhook", "delete", "hash123"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"isSuccess":true`) {
			t.Errorf("expected JSON in out, got: %s", out.String())
		}
	})
}
