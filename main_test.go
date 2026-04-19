package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHomeReturnsApplicationMetadata(t *testing.T) {
	t.Setenv("APP_NAME", "Go API PUCPR DevOps")
	t.Setenv("APP_VERSION", "1.0.0")

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)

	newServer().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	var payload map[string]string
	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatalf("unable to decode payload: %v", err)
	}

	if payload["status"] != "ok" {
		t.Fatalf("expected status ok, got %q", payload["status"])
	}

	if payload["project"] != "Go API PUCPR DevOps" {
		t.Fatalf("expected project Go API PUCPR DevOps, got %q", payload["project"])
	}
}

func TestHealthEndpoint(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/health", nil)

	newServer().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	if got := response.Body.String(); got != "{\"status\":\"healthy\"}\n" {
		t.Fatalf("unexpected body: %q", got)
	}
}

func TestAboutEndpoint(t *testing.T) {
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/about", nil)

	newServer().ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	var payload map[string]string
	if err := json.Unmarshal(response.Body.Bytes(), &payload); err != nil {
		t.Fatalf("unable to decode payload: %v", err)
	}

	if payload["stack"] != "Go" {
		t.Fatalf("expected Go stack, got %q", payload["stack"])
	}

	if payload["container"] != "Docker" {
		t.Fatalf("expected Docker container, got %q", payload["container"])
	}
}

func TestGetEnvFallsBackWhenMissing(t *testing.T) {
	const key = "UNSET_ENV_FOR_TEST"
	_ = os.Unsetenv(key)

	if got := getEnv(key, "fallback"); got != "fallback" {
		t.Fatalf("expected fallback value, got %q", got)
	}
}
