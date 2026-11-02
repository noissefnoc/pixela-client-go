package pixela

import (
	"context"
	"net/http"
	"testing"
)

func TestClient_Pixel(t *testing.T) {
	t.Run("CreatePixel", func(t *testing.T) {
		client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/v1/users/testuser/graphs/test-graph" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		})
		defer server.Close()

		res, err := client.CreatePixel(context.Background(), "test-graph", CreatePixelParams{
			Date:     "20261102",
			Quantity: "5",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !res.IsSuccess {
			t.Errorf("unexpected result: %+v", res)
		}
	})

	t.Run("GetPixel", func(t *testing.T) {
		client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/v1/users/testuser/graphs/test-graph/20261102" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"quantity":"5","optionalData":"{\"key\":\"val\"}"}`))
		})
		defer server.Close()

		pixel, err := client.GetPixel(context.Background(), "test-graph", "20261102")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if pixel.Quantity != "5" || pixel.OptionalData != `{"key":"val"}` {
			t.Errorf("unexpected pixel: %+v", pixel)
		}
	})

	t.Run("UpdatePixel", func(t *testing.T) {
		client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("expected PUT, got %s", r.Method)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		})
		defer server.Close()

		res, err := client.UpdatePixel(context.Background(), "test-graph", "20261102", UpdatePixelParams{
			Quantity: "10",
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !res.IsSuccess {
			t.Errorf("unexpected result: %+v", res)
		}
	})

	t.Run("IncrementPixel", func(t *testing.T) {
		client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("expected PUT, got %s", r.Method)
			}
			if r.URL.Path != "/v1/users/testuser/graphs/test-graph/increment" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		})
		defer server.Close()

		res, err := client.IncrementPixel(context.Background(), "test-graph")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !res.IsSuccess {
			t.Errorf("unexpected result: %+v", res)
		}
	})

	t.Run("DecrementPixel", func(t *testing.T) {
		client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPut {
				t.Errorf("expected PUT, got %s", r.Method)
			}
			if r.URL.Path != "/v1/users/testuser/graphs/test-graph/decrement" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		})
		defer server.Close()

		res, err := client.DecrementPixel(context.Background(), "test-graph")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !res.IsSuccess {
			t.Errorf("unexpected result: %+v", res)
		}
	})

	t.Run("DeletePixel", func(t *testing.T) {
		client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodDelete {
				t.Errorf("expected DELETE, got %s", r.Method)
			}
			if r.URL.Path != "/v1/users/testuser/graphs/test-graph/20261102" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		})
		defer server.Close()

		res, err := client.DeletePixel(context.Background(), "test-graph", "20261102")
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

		if _, err := client.CreatePixel(context.Background(), "Invalid", CreatePixelParams{}); err == nil {
			t.Error("expected error for invalid graphID")
		}
		if _, err := client.CreatePixel(context.Background(), "test-graph", CreatePixelParams{Date: "invalid"}); err == nil {
			t.Error("expected error for invalid date")
		}
		if _, err := client.CreatePixel(context.Background(), "test-graph", CreatePixelParams{Date: "20261102", Quantity: "invalid"}); err == nil {
			t.Error("expected error for invalid quantity")
		}
		if _, err := client.CreatePixel(context.Background(), "test-graph", CreatePixelParams{Date: "20261102", Quantity: "10", OptionalData: "not-json"}); err == nil {
			t.Error("expected error for invalid optionalData")
		}

		if _, err := client.GetPixel(context.Background(), "Invalid", "20261102"); err == nil {
			t.Error("expected error for invalid graphID")
		}
		if _, err := client.GetPixel(context.Background(), "test-graph", "invalid"); err == nil {
			t.Error("expected error for invalid date")
		}

		if _, err := client.UpdatePixel(context.Background(), "Invalid", "20261102", UpdatePixelParams{}); err == nil {
			t.Error("expected error for invalid graphID")
		}
		if _, err := client.UpdatePixel(context.Background(), "test-graph", "invalid", UpdatePixelParams{}); err == nil {
			t.Error("expected error for invalid date")
		}
		if _, err := client.UpdatePixel(context.Background(), "test-graph", "20261102", UpdatePixelParams{Quantity: "invalid"}); err == nil {
			t.Error("expected error for invalid quantity")
		}
		if _, err := client.UpdatePixel(context.Background(), "test-graph", "20261102", UpdatePixelParams{Quantity: "1", OptionalData: "not-json"}); err == nil {
			t.Error("expected error for invalid optionalData")
		}

		if _, err := client.IncrementPixel(context.Background(), "Invalid"); err == nil {
			t.Error("expected error for invalid graphID")
		}
		if _, err := client.DecrementPixel(context.Background(), "Invalid"); err == nil {
			t.Error("expected error for invalid graphID")
		}
		if _, err := client.DeletePixel(context.Background(), "Invalid", "20261102"); err == nil {
			t.Error("expected error for invalid graphID")
		}
		if _, err := client.DeletePixel(context.Background(), "test-graph", "invalid"); err == nil {
			t.Error("expected error for invalid date")
		}
	})
}
