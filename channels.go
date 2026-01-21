package main

import "time"

type T struct {
	i byte
	b bool
}

// Producer
func send(i int, ch chan<- *T) {
	t := &T{i: byte(i)}
	ch <- t

	t.b = true // UNSAFE

}

func ch1() {
	vs := make([]T, 5)     // [0,0,0,0,0]
	ch := make(chan *T, 5) // []

	for i := range vs {
		go send(i, ch) // Start 5 GoRoutine
	}

	time.Sleep(1 * time.Second) // Wait for all goroutines to finish

	for i := range vs {
		vs[i] = *<-ch
	}
	// Consumer
	for _, v := range vs {
		println(v.i, v.b)
	}

}
