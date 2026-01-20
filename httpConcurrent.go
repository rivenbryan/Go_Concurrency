package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

// A function to send a HTTP request and send the result to a channel
func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}

}

func httpF() {
	start := time.Now()
	results := make(chan result)
	list := []string{"https://www.google.com", "https://www.facebook.com", "https://www.youtube.com"}

	// Writing to the channel
	for _, url := range list {
		go get(url, results) // Creates a goroutine for each URL
	}

	// Reading from the channel
	for range list {
		r := <-results
		if r.err != nil {
			log.Println(r.err)
		} else {
			log.Printf("%s took %s\n", r.url, r.latency)
		}
	}
	log.Printf("Concurrent execution took: %s\n", time.Since(start))
}
