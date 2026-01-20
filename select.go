package main

import (
	"log"
	"time"
)

// Generate two go-routine
// One Go-routine produces data to the channel every 1 second
// Another Go-routine produces data to the channel every 5 second

// C1 [1]
// C2 []
// Reading: 1, 5,1,1
func selectGo() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for {
			channel1 <- "Channel 1"
			time.Sleep(1 * time.Second)
		}

	}()

	go func() {
		for {
			channel2 <- "Channel 2"
			time.Sleep(5 * time.Second)
		}
	}()

	for {
		select {
		case msg := <-channel1:
			log.Println(msg)
		case msg := <-channel2:
			log.Println(msg)
		default:
			log.Println("nothing ready")
		}
	}
}
