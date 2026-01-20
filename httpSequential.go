package main

import (
	"log"
	"net/http"
	"time"
)

// A function to send a HTTP request and send the result to a channel
func getSeq(url string) result {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		return result{url, err, 0}
	} else {
		t := time.Since(start)
		resp.Body.Close()
		return result{url, nil, t}
	}

}

func httpS() {
	start := time.Now()
	list := []string{"https://www.google.com", "https://www.facebook.com", "https://www.youtube.com"}

	// Reading from the channel
	for i := range list {
		r := getSeq(list[i])
		if r.err != nil {
			log.Println(r.err)
		} else {
			log.Printf("%s took %s\n", r.url, r.latency)
		}
	}
	log.Printf("Sequential execution took: %s\n", time.Since(start))
}
