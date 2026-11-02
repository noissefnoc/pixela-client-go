package cmd

import (
	"bytes"
	"context"
	"path/filepath"
	"strings"
	"testing"
)

func TestRootRun(t *testing.T) {
	ctx := context.Background()

	t.Run("no args shows usage", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d", code)
		}
		if !strings.Contains(out.String(), "Usage:") {
			t.Errorf("expected usage in out, got: %s", out.String())
		}
	})

	t.Run("-h shows usage", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-h"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d", code)
		}
		if !strings.Contains(out.String(), "Usage:") {
			t.Errorf("expected usage in out, got: %s", out.String())
		}
	})

	t.Run("--help shows usage", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"--help"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d", code)
		}
		if !strings.Contains(out.String(), "Usage:") {
			t.Errorf("expected usage in out, got: %s", out.String())
		}
	})

	t.Run("help command shows usage", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"help"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d", code)
		}
		if !strings.Contains(out.String(), "Usage:") {
			t.Errorf("expected usage in out, got: %s", out.String())
		}
	})

	t.Run("-v shows version", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"-v"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d", code)
		}
		if !strings.Contains(out.String(), "pixela version") {
			t.Errorf("expected version in out, got: %s", out.String())
		}
	})

	t.Run("version command shows version", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"version"}, nil, &out, &err)
		if code != 0 {
			t.Fatalf("expected 0, got %d", code)
		}
		if !strings.Contains(out.String(), "pixela version") {
			t.Errorf("expected version in out, got: %s", out.String())
		}
	})

	t.Run("unknown command returns 1", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"unknowncmd"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
		if !strings.Contains(err.String(), "unknown command") {
			t.Errorf("expected error message in err, got: %s", err.String())
		}
	})

	t.Run("missing authentication returns error", func(t *testing.T) {
		t.Setenv("PIXELA_USERNAME", "")
		t.Setenv("PIXELA_TOKEN", "")
		tmpDir := t.TempDir()
		emptyCfg := filepath.Join(tmpDir, "empty.yaml")

		var out, err bytes.Buffer
		code := Run(ctx, []string{"--config", emptyCfg, "graph", "get"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
		if !strings.Contains(err.String(), "authentication error") {
			t.Errorf("expected authentication error in err, got: %s", err.String())
		}
	})

	t.Run("invalid flag returns 1", func(t *testing.T) {
		var out, err bytes.Buffer
		code := Run(ctx, []string{"--invalidflag"}, nil, &out, &err)
		if code != 1 {
			t.Fatalf("expected 1, got %d", code)
		}
	})
}
