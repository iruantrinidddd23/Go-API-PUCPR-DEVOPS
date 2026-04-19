package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type app struct {
	name    string
	version string
}

func newServer() http.Handler {
	application := app{
		name:    getEnv("APP_NAME", "Go API PUCPR DevOps"),
		version: getEnv("APP_VERSION", "1.0.0"),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", application.handleHome)
	mux.HandleFunc("/health", application.handleHealth)
	mux.HandleFunc("/about", application.handleAbout)

	return mux
}

func (a app) handleHome(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"message": "API Go da disciplina de DevOps em execucao.",
		"status":  "ok",
		"project": a.name,
		"version": a.version,
	})
}

func (a app) handleHealth(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "healthy"})
}

func (a app) handleAbout(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{
		"stack":     "Go",
		"ci":        "GitHub Actions",
		"container": "Docker",
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, `{"status":"error"}`, http.StatusInternalServerError)
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func main() {
	port := getEnv("PORT", "8000")

	log.Printf("starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, newServer()); err != nil {
		log.Fatal(err)
	}
}
