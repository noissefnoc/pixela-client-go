package cmd

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestUserCommands(t *testing.T) {
	ctx := context.Background()

	t.Run("user help", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"user", "help"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d", code)
		}
		if !strings.Contains(out.String(), "Usage:") {
			t.Errorf("expected usage in out, got: %s", out.String())
		}
	})

	t.Run("user unknown command", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"user", "unknown"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
		if !strings.Contains(err.String(), "unknown user command") {
			t.Errorf("expected error in err, got: %s", err.String())
		}
	})

	t.Run("user create argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"user", "create", "onlyone"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
		if !strings.Contains(err.String(), "requires 2 arguments") {
			t.Errorf("expected argument error in err, got: %s", err.String())
		}
	})

	t.Run("user create success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" || r.URL.Path != "/v1/users" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		tmpDir := t.TempDir()
		cfgPath := filepath.Join(tmpDir, ".pixela.yaml")

		var out, err bytes.Buffer
		code := Run(ctx, []string{"--verbose", "--config", cfgPath, "user", "create", "testuser", "testtoken123"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"isSuccess":true`) {
			t.Errorf("expected JSON output in out, got: %s", out.String())
		}

		// Verify saved config file
		data, readErr := os.ReadFile(cfgPath)
		if readErr != nil {
			t.Fatalf("failed to read config file: %v", readErr)
		}
		if !strings.Contains(string(data), "username: testuser") || !strings.Contains(string(data), "token: testtoken123") {
			t.Errorf("unexpected config file content: %s", string(data))
		}
	})

	t.Run("user update success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" || r.URL.Path != "/v1/users/testuser" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			if r.Header.Get("X-USER-TOKEN") != "oldtoken123" {
				t.Errorf("unexpected token: %s", r.Header.Get("X-USER-TOKEN"))
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "oldtoken123", "-n", "user", "update", "newtoken456"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"isSuccess":true`) {
			t.Errorf("expected JSON in out, got: %s", out.String())
		}
	})

	t.Run("user delete success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "DELETE" || r.URL.Path != "/v1/users/testuser" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-n", "user", "delete", "testuser", "testtoken123"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"isSuccess":true`) {
			t.Errorf("expected JSON in out, got: %s", out.String())
		}
	})
}
