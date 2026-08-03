package main

// Fan-in: Multiple senders, one receiver.
// Fan-In means collecting results from multiple workers into a single channel.
// Bahut saare workers → Ek result channel.

// Multiple sources → One sink.

import (
	"fmt"
	"time"
)

// Worker function jo numbers ko ek channel mein send karta hai.
func worker(id int, nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return out
}

// Fanin function jo do channels ko ek channel mein merge karta hai.
func fanIn(cs ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, c := range cs {
			for n := range c {
				out <- n
			}
		}
	}()
	return out
}

// main function jo fan-in pattern ko implement karta hai.
func main() {
	w1 := worker(1, 1, 2, 3)
	w2 := worker(2, 4, 5, 6)
// ye line do channels ko ek channel mein merge karti hai.
	in := fanIn(w1, w2)

	// Sabhi numbers ko receive karna.
	for n := range in {
		fmt.Println(n)
	}
}
