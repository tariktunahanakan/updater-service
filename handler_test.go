package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestBuildManifest(t *testing.T) {
	want := `{"kernel":"6.1.42","firewall":"3.2.1","vpn":"2.8.0","webui":"4.0.5"}`
	got := buildManifest()
	if got != want {
		t.Errorf("manifest mismatch\n got: %s\nwant: %s", got, want)
	}
}

func TestHealthzIsFast(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	start := time.Now()
	healthzHandler(rec, req)
	elapsed := time.Since(start)

	if elapsed > 50*time.Millisecond {
		t.Errorf("healthz too slow: %v", elapsed)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}
}

func TestVersionHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/version", nil)
	rec := httptest.NewRecorder()
	versionHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", rec.Code)
	}
	if rec.Header().Get("Content-Type") != "application/json" {
		t.Errorf("wrong content type")
	}
}
