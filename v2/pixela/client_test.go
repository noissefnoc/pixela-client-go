package pixela

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newTestClient(t *testing.T, handler http.HandlerFunc) (*Client, *httptest.Server) {
	t.Helper()
	server := httptest.NewServer(handler)
	client, err := New("testuser", "testtoken", WithBaseURL(server.URL), WithHTTPClient(server.Client()))
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	return client, server
}

func TestNewClient(t *testing.T) {
	t.Run("valid credentials", func(t *testing.T) {
		c, err := New("validuser", "validtoken123", WithDebug(true))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if c.Username() != "validuser" {
			t.Errorf("expected username validuser, got %s", c.Username())
		}
		if c.BaseURL() != defaultBaseURL {
			t.Errorf("expected default base URL, got %s", c.BaseURL())
		}
	})

	t.Run("invalid username", func(t *testing.T) {
		_, err := New("Invalid_User", "validtoken123")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})

	t.Run("invalid token", func(t *testing.T) {
		_, err := New("validuser", "short")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestAPIError(t *testing.T) {
	client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"This user already exists.","isSuccess":false}`))
	})
	defer server.Close()

	_, err := client.CreateUser(context.Background(), CreateUserParams{
		AgreeTermsOfService: true,
		NotMinor:            true,
	})
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T: %v", err, err)
	}
	if apiErr.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", apiErr.StatusCode)
	}
	if apiErr.Message != "This user already exists." {
		t.Errorf("unexpected message: %s", apiErr.Message)
	}
	if apiErr.Error() == "" {
		t.Error("expected non-empty Error() string")
	}
}
