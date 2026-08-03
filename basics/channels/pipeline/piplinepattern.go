package main

// Pipeline is a chain of stages where the output of one stage becomes the input of the next stage.
// Pipeline ek function hai jo ek channel se input leta hai aur dusre channel mein output send karta hai.
// Isme stages hote hain jo data ko process karte hain.

import (
	"fmt"
)

// Pipeline pattern is used to process data in a sequence of stages.
// Pipeline pattern ek saath multiple stages mein data ko process karne ke liye use hota hai.
// Stages ek channel se data receive karte hain aur dusre channel mein send karte hain.

// Stage 1: Generate numbers.
// Generator function jo numbers ko ek channel mein send karta hai.
func generate(done <-chan struct{}, nums ...int) <-chan int {
// Generator function do arguments leta hai:
// 1. done <-chan struct{}: Cancellation channel.
// 2. nums ...int: Numbers jo channel mein send karne hain.
// ye channel ko open karta hai aur usme numbers send karta hai.
	out := make(chan int)
	// ye goroutine channel mein numbers ko send karti hai.
	go func() {
		// defer close(out) -> Ye goroutine khatam hone par channel ko close karta hai.
		defer close(out)
		for _, n := range nums {
			// select case: Cancellation channel ko listen karta hai.
			select {
			case <-done:
				return
			case out <- n:
			}
		}
	}()
	return out
}

// Stage 2: Square numbers.
func square(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-done:
				return
			case out <- n * n:
			}
		}
	}()
	return out
}

// main function jo pipeline pattern ko implement karta hai.
func main() {
	done := make(chan struct{})
	defer close(done)

	in := generate(done, 1, 2, 3, 4, 5)
	sq := square(done, in)

	for n := range sq {
		fmt.Println(n)
	}
}
