// used to record data from certstream for stress testing the project

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

const (
	targetCount = 50000
	wsURL       = "ws://localhost:9000/"
	outputFile  = "data.jsonl"
)

func main() {
	log.Printf("Starting recording data ...")
	log.Printf("URL: %s, Count: %v", wsURL, targetCount)

	// create file
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer file.Close()

	// add 64KB buffer to minimize syscalls
	writeBuffer := bufio.NewWriterSize(file, 64*1024)
	defer writeBuffer.Flush()

	// connect to local rust certstream server
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		log.Fatalf("failed to connect to %s: %v\nMake sure teh Docker container is running", wsURL, err)
	}
	defer conn.Close()

	// graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nReceived an interrupt, flushing captured data to disk")
		writeBuffer.Flush()
		file.Close()
		os.Exit(0)
	}()

	log.Println("Connected to docker container, starting recording process")
	startTime := time.Now()
	captured := 0

	for captured < targetCount {

		// read message
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatalf("failed to read websocket message: %v", err)
		}

		// write message
		if _, err = writeBuffer.Write(message); err != nil {
			log.Fatalf("failed to write to file: %v", err)
		}
		// write '\n' to add newline
		if err := writeBuffer.WriteByte('\n'); err != nil {
			log.Fatalf("failed to write newline character to file: %v", err)
		}

		// increment count
		captured++

		if captured%2500 == 0 || captured == targetCount {
			elapsed := time.Since(startTime).Seconds()
			rate := float64(captured) / elapsed
			percent := (float64(captured) / float64(targetCount)) * 100
			fmt.Printf("[%5.1f%%] %d / %d saved (%.1f certs/sec)\n", percent, captured, targetCount, rate)
		}
	}

	// flush to disk
	writeBuffer.Flush()
	totalElapsed := time.Since(startTime)
	fileInfo, _ := file.Stat()
	// Convert raw bytes to Megabytes (1024 bytes * 1024 KB = 1 MB)
	sizeMB := float64(fileInfo.Size()) / (1024 * 1024)

	fmt.Printf("\nDone! Successfully recorded %d certificates in %s\n", captured, totalElapsed.Round(time.Second))
	fmt.Printf("File: %s (%.2f MB)\n", outputFile, sizeMB)

}
