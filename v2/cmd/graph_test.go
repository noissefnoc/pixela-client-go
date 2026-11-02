package cmd

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGraphCommands(t *testing.T) {
	ctx := context.Background()

	t.Run("graph help", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"graph", "help"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d", code)
		}
		if !strings.Contains(out.String(), "Usage:") {
			t.Errorf("expected usage in out, got: %s", out.String())
		}
	})

	t.Run("graph unknown command", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"graph", "unknown"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
		if !strings.Contains(err.String(), "unknown graph command") {
			t.Errorf("expected error in err, got: %s", err.String())
		}
	})

	t.Run("graph create argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"graph", "create", "g1", "gname"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
		if !strings.Contains(err.String(), "requires 5 arguments") {
			t.Errorf("expected argument error in err, got: %s", err.String())
		}
	})

	t.Run("graph update argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"graph", "update"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
		if !strings.Contains(err.String(), "requires 1 argument") {
			t.Errorf("expected argument error in err, got: %s", err.String())
		}
	})

	t.Run("graph delete argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"graph", "delete"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("graph get argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"graph", "get", "extra"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("graph svg argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"graph", "svg"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("graph pixels argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"graph", "pixels"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("graph detail argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"graph", "detail"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("graph stat argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"graph", "stat"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("graph stopwatch argument error", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"graph", "stopwatch"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})

	t.Run("graph create success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" || r.URL.Path != "/v1/users/testuser/graphs" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "-n", "graph", "create", "test-graph", "graph-name", "commit", "int", "shibafu"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"isSuccess":true`) {
			t.Errorf("expected JSON in out, got: %s", out.String())
		}
	})

	t.Run("graph update success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "PUT" || r.URL.Path != "/v1/users/testuser/graphs/test-graph" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "-n", "graph", "update", "test-graph", "--name", "new-name"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"isSuccess":true`) {
			t.Errorf("expected JSON in out, got: %s", out.String())
		}
	})

	t.Run("graph delete success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "DELETE" || r.URL.Path != "/v1/users/testuser/graphs/test-graph" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "-n", "graph", "delete", "test-graph"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"isSuccess":true`) {
			t.Errorf("expected JSON in out, got: %s", out.String())
		}
	})

	t.Run("graph get success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" || r.URL.Path != "/v1/users/testuser/graphs" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"graphs":[{"id":"test-graph","name":"graph-name","unit":"commit","type":"int","color":"shibafu","timezone":"Asia/Tokyo","purgeCacheURLs":[]}]}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "graph", "get"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"test-graph"`) {
			t.Errorf("expected graph JSON in out, got: %s", out.String())
		}
	})

	t.Run("graph svg success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" || r.URL.Path != "/v1/users/testuser/graphs/test-graph" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`<svg>pixels</svg>`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "graph", "svg", "test-graph"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), "<svg>pixels</svg>") {
			t.Errorf("expected SVG in out, got: %s", out.String())
		}
	})

	t.Run("graph pixels success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" || r.URL.Path != "/v1/users/testuser/graphs/test-graph/pixels" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"pixels":["20261101","20261102"]}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "graph", "pixels", "test-graph"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), "20261101") {
			t.Errorf("expected pixels JSON in out, got: %s", out.String())
		}
	})

	t.Run("graph detail success", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "graph", "detail", "test-graph"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		expected := "https://pixe.la/v1/users/testuser/graphs/test-graph.html"
		if !strings.Contains(out.String(), expected) {
			t.Errorf("expected %s in out, got: %s", expected, out.String())
		}
	})

	t.Run("graph stat success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "GET" || r.URL.Path != "/v1/users/testuser/graphs/test-graph/stats" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"totalPixelsCount":5,"maxQuantity":10,"minQuantity":1,"totalQuantity":25,"avgQuantity":5.0,"todaysQuantity":3}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "graph", "stat", "test-graph"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"totalPixelsCount":5`) {
			t.Errorf("expected stats JSON in out, got: %s", out.String())
		}
	})

	t.Run("graph stopwatch success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" || r.URL.Path != "/v1/users/testuser/graphs/test-graph/stopwatch" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		}))
		defer server.Close()

		t.Setenv("PIXELA_API_ENDPOINT", server.URL)
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-u", "testuser", "-t", "testtoken", "-n", "graph", "stopwatch", "test-graph"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d (err: %s)", code, err.String())
		}
		if !strings.Contains(out.String(), `"isSuccess":true`) {
			t.Errorf("expected JSON in out, got: %s", out.String())
		}
	})
}
