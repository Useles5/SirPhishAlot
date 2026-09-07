package main

import (
	"io"
	"log"
	"net/http"

	"github.com/Useles5/sirphishalot/internal/parser"
)

func main() {
	http.HandleFunc("/ingest", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.Header().Set("Allow", "POST")
			http.Error(w, "Failed to read body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// extract the domains
		_, err = parser.ExtractDomains(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// send OK header
		w.WriteHeader(http.StatusOK)
	})

	log.Println("SirPhishAlot Ingestion Node running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}
