package pixela

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestClient_CreateUser(t *testing.T) {
	client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/v1/users" {
			t.Errorf("expected /v1/users, got %s", r.URL.Path)
		}

		var payload createUserPayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatalf("failed to decode body: %v", err)
		}
		if payload.Username != "testuser" || payload.Token != "testtoken" {
			t.Errorf("unexpected payload: %+v", payload)
		}
		if payload.AgreeTermsOfService != "yes" || payload.NotMinor != "yes" {
			t.Errorf("unexpected agreements: %+v", payload)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
	})
	defer server.Close()

	res, err := client.CreateUser(context.Background(), CreateUserParams{
		AgreeTermsOfService: true,
		NotMinor:            true,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !res.IsSuccess || res.Message != "Success." {
		t.Errorf("unexpected result: %+v", res)
	}
}

func TestClient_UpdateUser(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("expected PUT, got %s", r.Method)
			}
			if r.URL.Path != "/v1/users/testuser" {
				t.Errorf("expected /v1/users/testuser, got %s", r.URL.Path)
			}
			if r.Header.Get("X-USER-TOKEN") != "testtoken" {
				t.Errorf("unexpected token header: %s", r.Header.Get("X-USER-TOKEN"))
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		})
		defer server.Close()

		res, err := client.UpdateUser(context.Background(), UpdateUserParams{
			NewToken: "newtoken1234",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !res.IsSuccess {
			t.Errorf("unexpected result: %+v", res)
		}
	})

	t.Run("invalid new token", func(t *testing.T) {
		client, server := newTestClient(t, nil)
		defer server.Close()

		_, err := client.UpdateUser(context.Background(), UpdateUserParams{
			NewToken: "short",
		})
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	})
}

func TestClient_DeleteUser(t *testing.T) {
	client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.URL.Path != "/v1/users/testuser" {
			t.Errorf("expected /v1/users/testuser, got %s", r.URL.Path)
		}
		if r.Header.Get("X-USER-TOKEN") != "testtoken" {
			t.Errorf("unexpected token header: %s", r.Header.Get("X-USER-TOKEN"))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
	})
	defer server.Close()

	res, err := client.DeleteUser(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !res.IsSuccess {
		t.Errorf("unexpected result: %+v", res)
	}
}
