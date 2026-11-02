package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

// mockServer represents a stateful in-memory mock Pixela server for CLI integration tests.
type mockServer struct {
	mu       sync.Mutex
	users    map[string]string            // username -> token
	graphs   map[string]map[string]string // username -> graphID -> name
	pixels   map[string]map[string]string // username/graphID -> date -> quantity
	webhooks map[string]map[string]string // username -> hash -> type
}

func newMockServer() *mockServer {
	return &mockServer{
		users:    make(map[string]string),
		graphs:   make(map[string]map[string]string),
		pixels:   make(map[string]map[string]string),
		webhooks: make(map[string]map[string]string),
	}
}

func (s *mockServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	path := r.URL.Path

	// Helper for JSON response
	respond := func(status int, msg string, success bool, extra map[string]any) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		res := map[string]any{
			"message":   msg,
			"isSuccess": success,
		}
		for k, v := range extra {
			res[k] = v
		}
		json.NewEncoder(w).Encode(res)
	}

	// 1. User APIs
	if path == "/v1/users" && r.Method == "POST" {
		var payload struct {
			Username string `json:"username"`
			Token    string `json:"token"`
		}
		json.NewDecoder(r.Body).Decode(&payload)
		s.users[payload.Username] = payload.Token
		respond(http.StatusOK, "Success.", true, nil)
		return
	}

	if strings.HasPrefix(path, "/v1/users/") {
		parts := strings.Split(strings.Trim(path, "/"), "/")
		username := parts[2]

		// User update / delete
		if len(parts) == 3 {
			if r.Method == "PUT" {
				var payload struct {
					NewToken string `json:"newToken"`
				}
				json.NewDecoder(r.Body).Decode(&payload)
				s.users[username] = payload.NewToken
				respond(http.StatusOK, "Success.", true, nil)
				return
			}
			if r.Method == "DELETE" {
				delete(s.users, username)
				respond(http.StatusOK, "Success.", true, nil)
				return
			}
		}

		// Graph APIs
		if len(parts) >= 4 && parts[3] == "graphs" {
			if len(parts) == 4 {
				if r.Method == "POST" {
					var payload struct {
						ID   string `json:"id"`
						Name string `json:"name"`
					}
					json.NewDecoder(r.Body).Decode(&payload)
					if s.graphs[username] == nil {
						s.graphs[username] = make(map[string]string)
					}
					s.graphs[username][payload.ID] = payload.Name
					respond(http.StatusOK, "Success.", true, nil)
					return
				}
				if r.Method == "GET" {
					w.Header().Set("Content-Type", "application/json")
					var list []map[string]string
					for id, name := range s.graphs[username] {
						list = append(list, map[string]string{"id": id, "name": name})
					}
					json.NewEncoder(w).Encode(map[string]any{"graphs": list})
					return
				}
			}

			graphID := parts[4]
			key := fmt.Sprintf("%s/%s", username, graphID)

			if len(parts) == 5 {
				if r.Method == "PUT" {
					respond(http.StatusOK, "Success.", true, nil)
					return
				}
				if r.Method == "DELETE" {
					delete(s.graphs[username], graphID)
					respond(http.StatusOK, "Success.", true, nil)
					return
				}
				if r.Method == "GET" {
					w.Header().Set("Content-Type", "image/svg+xml")
					w.Write([]byte(`<svg>pixels</svg>`))
					return
				}
				if r.Method == "POST" {
					// Pixel create
					var payload struct {
						Date     string `json:"date"`
						Quantity string `json:"quantity"`
					}
					json.NewDecoder(r.Body).Decode(&payload)
					if s.pixels[key] == nil {
						s.pixels[key] = make(map[string]string)
					}
					s.pixels[key][payload.Date] = payload.Quantity
					respond(http.StatusOK, "Success.", true, nil)
					return
				}
			}

			// Sub-resources under graphs
			if len(parts) == 6 {
				sub := parts[5]
				switch sub {
				case "increment":
					respond(http.StatusOK, "Success.", true, nil)
					return
				case "decrement":
					respond(http.StatusOK, "Success.", true, nil)
					return
				case "stats":
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(map[string]any{
						"totalPixelsCount": len(s.pixels[key]),
					})
					return
				case "pixels":
					w.Header().Set("Content-Type", "application/json")
					var dates []string
					for d := range s.pixels[key] {
						dates = append(dates, d)
					}
					json.NewEncoder(w).Encode(map[string]any{"pixels": dates})
					return
				default:
					// Date specific pixel (e.g. /v1/users/<user>/graphs/<graph>/<date>)
					date := sub
					if r.Method == "GET" {
						w.Header().Set("Content-Type", "application/json")
						qty := s.pixels[key][date]
						json.NewEncoder(w).Encode(map[string]string{"quantity": qty})
						return
					}
					if r.Method == "PUT" {
						var payload struct {
							Quantity string `json:"quantity"`
						}
						json.NewDecoder(r.Body).Decode(&payload)
						s.pixels[key][date] = payload.Quantity
						respond(http.StatusOK, "Success.", true, nil)
						return
					}
					if r.Method == "DELETE" {
						delete(s.pixels[key], date)
						respond(http.StatusOK, "Success.", true, nil)
						return
					}
				}
			}
		}

		// Webhook APIs
		if len(parts) >= 4 && parts[3] == "webhooks" {
			if len(parts) == 4 {
				if r.Method == "POST" {
					var payload struct {
						GraphID string `json:"graphID"`
						Type    string `json:"type"`
					}
					json.NewDecoder(r.Body).Decode(&payload)
					hash := "mock-hash-12345"
					if s.webhooks[username] == nil {
						s.webhooks[username] = make(map[string]string)
					}
					s.webhooks[username][hash] = payload.Type
					respond(http.StatusOK, "Success.", true, map[string]any{"webhookHash": hash})
					return
				}
				if r.Method == "GET" {
					w.Header().Set("Content-Type", "application/json")
					var list []map[string]string
					for h, t := range s.webhooks[username] {
						list = append(list, map[string]string{"webhookHash": h, "type": t})
					}
					json.NewEncoder(w).Encode(map[string]any{"webhooks": list})
					return
				}
			}
			if len(parts) == 5 {
				hash := parts[4]
				if r.Method == "POST" {
					respond(http.StatusOK, "Success.", true, nil)
					return
				}
				if r.Method == "DELETE" {
					delete(s.webhooks[username], hash)
					respond(http.StatusOK, "Success.", true, nil)
					return
				}
			}
		}
	}

	http.NotFound(w, r)
}

// TestCLI_E2ELifecycle tests the entire lifecycle of Pixela resources through the CLI.
func TestCLI_E2ELifecycle(t *testing.T) {
	mock := newMockServer()
	server := httptest.NewServer(mock)
	defer server.Close()

	t.Setenv("PIXELA_API_ENDPOINT", server.URL)

	ctx := context.Background()
	tmpDir := t.TempDir()
	cfgPath := filepath.Join(tmpDir, ".pixela.yaml")

	runCLI := func(args ...string) (int, string, string) {
		var out, err bytes.Buffer
		baseArgs := append([]string{"--config", cfgPath}, args...)
		code := Run(ctx, baseArgs, nil, &out, &err)
		return code, out.String(), err.String()
	}

	// 1. Create user
	code, out, errStr := runCLI("--verbose", "user", "create", "alice", "secr3ttoken")
	if code != 0 {
		t.Fatalf("user create failed: code=%d err=%s", code, errStr)
	}
	if !strings.Contains(out, `"isSuccess":true`) {
		t.Errorf("expected success message in user create, got: %s", out)
	}

	// 2. Create graph using credentials saved in ~/.pixela.yaml
	code, _, errStr = runCLI("graph", "create", "commit-graph", "Commits", "commit", "int", "shibafu")
	if code != 0 {
		t.Fatalf("graph create failed: code=%d err=%s", code, errStr)
	}

	// 3. Get graph definitions
	code, out, errStr = runCLI("graph", "get")
	if code != 0 {
		t.Fatalf("graph get failed: code=%d err=%s", code, errStr)
	}
	if !strings.Contains(out, "commit-graph") {
		t.Errorf("expected commit-graph in definitions, got: %s", out)
	}

	// 4. Record a pixel
	code, _, errStr = runCLI("pixel", "create", "commit-graph", "20261102", "5")
	if code != 0 {
		t.Fatalf("pixel create failed: code=%d err=%s", code, errStr)
	}

	// 5. Get pixel value
	code, out, errStr = runCLI("pixel", "get", "commit-graph", "20261102")
	if code != 0 {
		t.Fatalf("pixel get failed: code=%d err=%s", code, errStr)
	}
	if !strings.Contains(out, `"quantity":"5"`) {
		t.Errorf("expected quantity 5, got: %s", out)
	}

	// 6. Update pixel value
	code, _, errStr = runCLI("pixel", "update", "commit-graph", "20261102", "7")
	if code != 0 {
		t.Fatalf("pixel update failed: code=%d err=%s", code, errStr)
	}

	// 7. Increment pixel
	code, _, errStr = runCLI("pixel", "increment", "commit-graph")
	if code != 0 {
		t.Fatalf("pixel increment failed: code=%d err=%s", code, errStr)
	}

	// 8. Create webhook
	code, out, errStr = runCLI("--verbose", "webhook", "create", "commit-graph", "increment")
	if code != 0 {
		t.Fatalf("webhook create failed: code=%d err=%s", code, errStr)
	}
	if !strings.Contains(out, "mock-hash-12345") {
		t.Errorf("expected webhook hash, got: %s", out)
	}

	// 9. Invoke webhook
	code, _, errStr = runCLI("webhook", "invoke", "mock-hash-12345")
	if code != 0 {
		t.Fatalf("webhook invoke failed: code=%d err=%s", code, errStr)
	}

	// 10. Delete webhook
	code, _, errStr = runCLI("webhook", "delete", "mock-hash-12345")
	if code != 0 {
		t.Fatalf("webhook delete failed: code=%d err=%s", code, errStr)
	}

	// 11. Delete pixel
	code, _, errStr = runCLI("pixel", "delete", "commit-graph", "20261102")
	if code != 0 {
		t.Fatalf("pixel delete failed: code=%d err=%s", code, errStr)
	}

	// 12. Delete graph
	code, _, errStr = runCLI("graph", "delete", "commit-graph")
	if code != 0 {
		t.Fatalf("graph delete failed: code=%d err=%s", code, errStr)
	}

	// 13. Delete user
	code, _, errStr = runCLI("user", "delete", "alice", "secr3ttoken")
	if code != 0 {
		t.Fatalf("user delete failed: code=%d err=%s", code, errStr)
	}
}
