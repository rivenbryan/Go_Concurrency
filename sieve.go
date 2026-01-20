package main

import "fmt"

// Prime Sieve Exercise
// The purpose is to find Prime Numbers from [2, limit]
// We have a generator that generates number from [2, limit]

// A number is prime IIF can be divisible by itself.

// The first number is a prime (2) -> Create a rule: If the number can be divided by 2, it is not a prime number
// The next number is a prime (3) -> Create a rule: If the number can be divided by 2, it is not a prime number
// The next number is a not a prime (4) <-- Why? it can be divided by 2.
// The next number is a prime (5) <-- it can NOT be divided by 2,3,
// The next number is a not a prime (6) <-- it can be divided by 2
// The next number is not a prime (9) <-- it can be divided by 3
// The next number is a prime (11) <-- it can NOT be divided by 2,3,5,7

// Generates or Sends data to the channel to main
func generator(limit int, ch chan<- int) {
	for i := 2; i < limit; i++ { // Current: 3
		ch <- i // For each number, write it to the channel
	}
	close(ch)
}

// Generator -> c0 -> filter(2) -> c1 -> filter(3) -> c2 -> sieve
//
// c2 -> filter(5) -> c3 -> sieve
//
// The parameters, filter(2) indicates checking if the number can be divided by 2
// Receives the data to the channel
func filter(number int, currentChannel <-chan int, nextChannel chan<- int) {
	// Filter(2) ->
	for i := range currentChannel { // For all the numbers in the source channel, we check if it is a prime number
		if i%number != 0 { // If the src % 2 != 0, it is potentially a prime number
			nextChannel <- i // We have to pass to the next filter to check
		}
	}
	close(nextChannel)
}

func sieve(limit int) {
	currentChannel := make(chan int)

	// This creates a channel for generator
	go generator(limit, currentChannel)

	for {

		number, ok := <-currentChannel // Number : 2
		if !ok {
			break
		}
		nextChannel := make(chan int)
		// Create filters only for prime number
		go filter(number, currentChannel, nextChannel)

		currentChannel = nextChannel
		// 2
		fmt.Printf("Prime: %d\n", number)
	}
}
