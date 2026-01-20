package main

import "fmt"

func main() {
	fmt.Println("Starting Concurrent HTTP Requests...")
	httpF()
	fmt.Println()
	fmt.Println("Starting Sequential HTTP Requests...")
	httpS()
}

// Generator -> Sieve
