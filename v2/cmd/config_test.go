package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfigFile(t *testing.T) {
	tmpDir := t.TempDir()
	cfgPath := filepath.Join(tmpDir, ".pixela.yaml")

	content := "# Comment\nusername: testuser\ntoken: \"mytoken\"\n"
	if err := os.WriteFile(cfgPath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	u, tok, err := loadConfigFile(cfgPath)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u != "testuser" {
		t.Errorf("got %q, want %q", u, "testuser")
	}
	if tok != "mytoken" {
		t.Errorf("got %q, want %q", tok, "mytoken")
	}

	// Non-existent file should succeed with empty values
	u, tok, err = loadConfigFile(filepath.Join(tmpDir, "nonexistent.yaml"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u != "" || tok != "" {
		t.Errorf("expected empty strings, got user=%q tok=%q", u, tok)
	}
}

func TestSaveConfigFile(t *testing.T) {
	tmpDir := t.TempDir()
	cfgPath := filepath.Join(tmpDir, ".pixela.yaml")
	var buf bytes.Buffer

	// First save should succeed
	if err := saveConfigFile(cfgPath, "user1", "tok1", &buf); err != nil {
		t.Fatalf("saveConfigFile failed: %v", err)
	}

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		t.Fatalf("failed to read saved file: %v", err)
	}
	expected := "username: user1\ntoken: tok1\n"
	if string(data) != expected {
		t.Errorf("got %q, want %q", string(data), expected)
	}

	// Second save should fail because file already exists
	buf.Reset()
	if err := saveConfigFile(cfgPath, "user2", "tok2", &buf); err == nil {
		t.Errorf("expected error when saving to existing file")
	}
	if !bytes.Contains(buf.Bytes(), []byte("already a config file")) {
		t.Errorf("expected warning message, got: %s", buf.String())
	}
}
