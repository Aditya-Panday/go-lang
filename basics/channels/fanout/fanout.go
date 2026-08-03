package main

// Fan-out: One sender, multiple receivers.
// Fan-Out means distributing work from one source to multiple workers.
// Ek source → Bahut saare workers.

import (
	"fmt"
	"sync"
	"time"
)

// Producer function jo numbers ko ek channel mein send karta hai.
func produce(nums ...int) <-chan int {
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

// Fanout function jo ek channel se numbers ko do channels mein send karta hai.
func fanOut(in <-chan int) (<-chan int, <-chan int) {
	out1 := make(chan int)
	out2 := make(chan int)

	go func() {
		defer close(out1)
		defer close(out2)
		for n := range in {
			out1 <- n
			out2 <- n
		}
	}()
	return out1, out2
}



// main function jo fan-out pattern ko implement karta hai.
func main() {
	in := produce(1, 2, 3, 4, 5)
// ye line ek hi channel se do
	out1, out2 := fanOut(in)
// sync.WaitGroup ka use ek se zyada goroutines ko sync karne ke liye kiya jata hai.
	var wg sync.WaitGroup
// wg.Add(2) -> Ye goroutine ko batata hai ki do goroutines hain jo wait karengi.
	wg.Add(2)
// Pehli goroutine jo out1 channel se numbers ko receive karti hai.
	go func() {
		// defer wg.Done() -> Ye Goroutine khatam hone par wg ko notify karta hai.
		defer wg.Done()
		for n := range out1 {
			fmt.Println("Out1:", n)
		}
	}()

	// Dusri goroutine jo out2 channel se numbers ko receive karti hai.
	go func() {
		defer wg.Done()
		for n := range out2 {
			fmt.Println("Out2:", n)
		}
	}()

	wg.Wait()
}

