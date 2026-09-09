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

		domainsList, err := parser.ExtractDomains(body)
		if err != nil {
			http.Error(w, "Failed to extract domains", http.StatusBadRequest)
		}
		// initialize scanner
		scanner := parser.NewDomainScanner(domainsList)
		count := 0

		for {
			domain, hasNext := scanner.Next()
			if !hasNext {
				break // no more domain
			}

			// TODO: Send to Redpanda/NATS
			log.Printf("Found domain: %s", domain)
			count++
		}

		log.Printf("Successfully ingested %d domains", count)

		// send OK header
		w.WriteHeader(http.StatusOK)
	})

	log.Println("SirPhishAlot Ingestion Node running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server crashed: %v", err)
	}
}
