package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Run this test with: go test -v -run TestPerformance -- -duration=5 -write=10 -read=1
func TestPerformance(t *testing.T) {
	durationPtr := flag.Int("duration", 5, "Duration in minutes")
	writePtr := flag.Int("write", 10, "Write interval in ms")
	readPtr := flag.Int("read", 1, "Read interval in ms")
	flag.Parse()

	durationMin := *durationPtr
	writeInterval := time.Duration(*writePtr) * time.Millisecond
	readInterval := time.Duration(*readPtr) * time.Millisecond

	baseURL := "http://localhost:8080"
	duration := time.Duration(durationMin) * time.Minute
	endTime := time.Now().Add(duration)

	var wg sync.WaitGroup
	var mu sync.Mutex
	var keys []string

	// Start writer
	go func() {
		for time.Now().Before(endTime) {
			url := "https://example.com"
			payload := []byte(fmt.Sprintf(`{"url":"%s"}`, url))
			resp, err := http.Post(baseURL+"/shorten", "application/json", bytes.NewBuffer(payload))
			if err == nil && resp.StatusCode == http.StatusCreated {
				var res map[string]string
				json.NewDecoder(resp.Body).Decode(&res)
				mu.Lock()
				keys = append(keys, res["short_url"])
				mu.Unlock()
			}
			time.Sleep(writeInterval)
		}
	}()

	// Start reader/load generator
	wg.Add(1)
	go func() {
		defer wg.Done()
		for time.Now().Before(endTime) {
			mu.Lock()
			if len(keys) == 0 {
				mu.Unlock()
				time.Sleep(10 * time.Millisecond)
				continue
			}
			key := keys[len(keys)-1]
			mu.Unlock()

			resp, err := http.Get(baseURL + "/" + key)
			if err == nil {
				assert.NotEqual(t, http.StatusInternalServerError, resp.StatusCode)
			}
			time.Sleep(readInterval)
		}
	}()

	wg.Wait()
	fmt.Printf("Performance test completed over %v (writeInterval: %v, readInterval: %v)\n", duration, writeInterval, readInterval)
}
