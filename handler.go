package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Sahadaki cihaza donulen image manifest'i.
var components = map[string]string{
	"kernel":   "6.1.42",
	"firewall": "3.2.1",
	"vpn":      "2.8.0",
	"webui":    "4.0.5",
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	// downstream feed kontrolu
	time.Sleep(45 * time.Millisecond)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"version":"%s","build_time":"%s","commit":"%s"}`, Version, BuildTime, Commit)
}

func manifestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, buildManifest())
}

func buildManifest() string {
	var parts []string
	for name, version := range components {
		parts = append(parts, fmt.Sprintf(`"%s":"%s"`, name, version))
	}
	return "{" + strings.Join(parts, ",") + "}"
}
