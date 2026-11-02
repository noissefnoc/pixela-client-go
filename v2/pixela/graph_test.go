package pixela

import (
	"context"
	"net/http"
	"testing"
)

func TestClient_CreateGraph(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if r.URL.Path != "/v1/users/testuser/graphs" {
				t.Errorf("expected /v1/users/testuser/graphs, got %s", r.URL.Path)
			}
			if r.Header.Get("X-USER-TOKEN") != "testtoken" {
				t.Errorf("unexpected token header")
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
		})
		defer server.Close()

		res, err := client.CreateGraph(context.Background(), CreateGraphParams{
			ID:             "test-graph",
			Name:           "Test Graph",
			Unit:           "commit",
			Type:           "int",
			Color:          "shibafu",
			SelfSufficient: "increment",
		})
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

		if _, err := client.CreateGraph(context.Background(), CreateGraphParams{ID: "Invalid"}); err == nil {
			t.Error("expected error for invalid ID")
		}
		if _, err := client.CreateGraph(context.Background(), CreateGraphParams{ID: "valid", Type: "invalid"}); err == nil {
			t.Error("expected error for invalid Type")
		}
		if _, err := client.CreateGraph(context.Background(), CreateGraphParams{ID: "valid", Type: "int", Color: "invalid"}); err == nil {
			t.Error("expected error for invalid Color")
		}
		if _, err := client.CreateGraph(context.Background(), CreateGraphParams{ID: "valid", Type: "int", Color: "shibafu", SelfSufficient: "invalid"}); err == nil {
			t.Error("expected error for invalid SelfSufficient")
		}
	})
}

func TestClient_GetGraph(t *testing.T) {
	expectedSVG := `<svg>sample graph</svg>`
	client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if r.URL.Path != "/v1/users/testuser/graphs/test-graph" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.URL.Query().Get("date") != "20261102" || r.URL.Query().Get("mode") != "badge" || r.URL.Query().Get("appearance") != "dark" {
			t.Errorf("unexpected query: %s", r.URL.RawQuery)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expectedSVG))
	})
	defer server.Close()

	svgBytes, err := client.GetGraph(context.Background(), "test-graph", GetGraphOptions{
		Date:       "20261102",
		Mode:       "badge",
		Appearance: "dark",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(svgBytes) != expectedSVG {
		t.Errorf("unexpected SVG: %s", string(svgBytes))
	}

	// invalid graphID
	if _, err := client.GetGraph(context.Background(), "Invalid"); err == nil {
		t.Error("expected error for invalid graphID")
	}
}

func TestClient_GetGraphDefinitions(t *testing.T) {
	client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"graphs":[{"id":"test-graph","name":"Test Graph","unit":"commit","type":"int","color":"shibafu","timezone":"Asia/Tokyo"}]}`))
	})
	defer server.Close()

	defs, err := client.GetGraphDefinitions(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(defs.Graphs) != 1 || defs.Graphs[0].ID != "test-graph" {
		t.Errorf("unexpected definitions: %+v", defs)
	}
}

func TestClient_GetGraphDetail(t *testing.T) {
	client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/users/testuser/graphs/test-graph/graph-def" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"id":"test-graph","name":"Test Graph","unit":"commit","type":"int","color":"shibafu"}`))
	})
	defer server.Close()

	g, err := client.GetGraphDetail(context.Background(), "test-graph")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if g.ID != "test-graph" {
		t.Errorf("unexpected graph: %+v", g)
	}

	if _, err := client.GetGraphDetail(context.Background(), "Invalid"); err == nil {
		t.Error("expected error for invalid graphID")
	}
}

func TestClient_GetGraphStats(t *testing.T) {
	client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"totalPixelsCount":10,"maxQuantity":15,"minQuantity":1,"totalQuantity":50,"avgQuantity":5.0,"todaysQuantity":2}`))
	})
	defer server.Close()

	stats, err := client.GetGraphStats(context.Background(), "test-graph")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if stats.TotalPixelsCount != 10 || stats.AvgQuantity != 5.0 {
		t.Errorf("unexpected stats: %+v", stats)
	}

	if _, err := client.GetGraphStats(context.Background(), "Invalid"); err == nil {
		t.Error("expected error for invalid graphID")
	}
}

func TestClient_GetPixelDates(t *testing.T) {
	client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("from") != "20261101" || r.URL.Query().Get("to") != "20261102" {
			t.Errorf("unexpected query: %s", r.URL.RawQuery)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"pixels":["20261101","20261102"]}`))
	})
	defer server.Close()

	dates, err := client.GetPixelDates(context.Background(), "test-graph", "20261101", "20261102")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(dates.Pixels) != 2 || dates.Pixels[0] != "20261101" {
		t.Errorf("unexpected dates: %+v", dates)
	}

	// invalid params
	if _, err := client.GetPixelDates(context.Background(), "Invalid", "", ""); err == nil {
		t.Error("expected error for invalid graphID")
	}
	if _, err := client.GetPixelDates(context.Background(), "test-graph", "invalid", ""); err == nil {
		t.Error("expected error for invalid from")
	}
	if _, err := client.GetPixelDates(context.Background(), "test-graph", "", "invalid"); err == nil {
		t.Error("expected error for invalid to")
	}
}

func TestClient_UpdateGraph(t *testing.T) {
	client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
	})
	defer server.Close()

	res, err := client.UpdateGraph(context.Background(), "test-graph", UpdateGraphParams{
		Name:           "Updated Graph",
		Color:          "momiji",
		SelfSufficient: "decrement",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !res.IsSuccess {
		t.Errorf("unexpected result: %+v", res)
	}

	// invalid params
	if _, err := client.UpdateGraph(context.Background(), "Invalid", UpdateGraphParams{}); err == nil {
		t.Error("expected error for invalid graphID")
	}
	if _, err := client.UpdateGraph(context.Background(), "test-graph", UpdateGraphParams{Color: "invalid"}); err == nil {
		t.Error("expected error for invalid color")
	}
	if _, err := client.UpdateGraph(context.Background(), "test-graph", UpdateGraphParams{SelfSufficient: "invalid"}); err == nil {
		t.Error("expected error for invalid selfSufficient")
	}
}

func TestClient_DeleteGraph(t *testing.T) {
	client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
	})
	defer server.Close()

	res, err := client.DeleteGraph(context.Background(), "test-graph")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !res.IsSuccess {
		t.Errorf("unexpected result: %+v", res)
	}

	if _, err := client.DeleteGraph(context.Background(), "Invalid"); err == nil {
		t.Error("expected error for invalid graphID")
	}
}

func TestClient_StopGraphStopwatch(t *testing.T) {
	client, server := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if r.URL.Path != "/v1/users/testuser/graphs/test-graph/stopwatch" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Success.","isSuccess":true}`))
	})
	defer server.Close()

	res, err := client.StopGraphStopwatch(context.Background(), "test-graph")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !res.IsSuccess {
		t.Errorf("unexpected result: %+v", res)
	}

	if _, err := client.StopGraphStopwatch(context.Background(), "Invalid"); err == nil {
		t.Error("expected error for invalid graphID")
	}
}
