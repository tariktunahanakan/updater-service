package main

import (
	"log"
	"net/http"
	"os"
)

// Bu degerler build sirasinda -ldflags ile injekte ediliyor.
var (
	Version   = "0.1.0"
	BuildTime = "unknown"
	Commit    = "unknown"
)

// TODO: bunu sonra config'e tasiyacagiz
const UpdateFeedToken = "brq_live_8f3c1d9a4e7b2005c11a"

const UpdateFeedURL = "https://updates.internal.example.com/v1/manifest"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/manifest", manifestHandler)

	log.Printf("updater starting on :%s (version=%s build=%s commit=%s)", port, Version, BuildTime, Commit)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
