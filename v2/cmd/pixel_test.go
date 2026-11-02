package cmd

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPixelCommands(t *testing.T) {
	ctx := context.Background()

	t.Run("pixel help", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"pixel", "help"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d", code)
		}
		if !strings.Contains(out.String(), "Usage:") {
			t.Errorf("expected usage in out, got: %s", out.String())
		}
	})

	t.Run("pixel unknown command", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"pixel", "unknown"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
		if !strings.Contains(err.String(), "unknown pixel command") {
			t.Errorf("expected error in err, got: %s", err.String())
		}
	})

	t.Run("pixel create argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"pixel", "create", "g1"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("pixel get argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"pixel", "get", "g1"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("pixel update argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"pixel", "update", "g1"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("pixel delete argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"pixel", "delete", "g1"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("pixel increment argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"pixel", "increment"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("pixel decrement argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"pixel", "decrement"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("pixel create success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" || r.URL.Path != "/v1/users/testuser/graphs/test-graph" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "-n", "pixel", "create", "test-graph", "20261102", "5"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"isSuccess":true`) {
			t.Errorf("expected JSON in out, got: %s", out.String())
		}
	})

	t.Run("pixel get success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" || r.URL.Path != "/v1/users/testuser/graphs/test-graph/20261102" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"quantity":"5","optionalData":"{\"key\":\"val\"}"}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "pixel", "get", "test-graph", "20261102"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"quantity":"5"`) {
			t.Errorf("expected pixel JSON in out, got: %s", out.String())
		}
	})

	t.Run("pixel update success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" || r.URL.Path != "/v1/users/testuser/graphs/test-graph/20261102" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "-n", "pixel", "update", "test-graph", "20261102", "10"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"isSuccess":true`) {
			t.Errorf("expected JSON in out, got: %s", out.String())
		}
	})

	t.Run("pixel delete success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "DELETE" || r.URL.Path != "/v1/users/testuser/graphs/test-graph/20261102" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "-n", "pixel", "delete", "test-graph", "20261102"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"isSuccess":true`) {
			t.Errorf("expected JSON in out, got: %s", out.String())
		}
	})

	t.Run("pixel increment success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" || r.URL.Path != "/v1/users/testuser/graphs/test-graph/increment" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "-n", "pixel", "increment", "test-graph"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"isSuccess":true`) {
			t.Errorf("expected JSON in out, got: %s", out.String())
		}
	})

	t.Run("pixel decrement success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" || r.URL.Path != "/v1/users/testuser/graphs/test-graph/decrement" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "-n", "pixel", "decrement", "test-graph"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"isSuccess":true`) {
			t.Errorf("expected JSON in out, got: %s", out.String())
		}
	})
}
