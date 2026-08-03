package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Atomic operations perform simple operations safely without using a mutex.
// Chhote operations ke liye lightweight lock.

// Without Atomic

// counter++

// Race.

// Atomic

// atomic.AddInt64(&counter,1)

// Safe.

// Kab Use Kare?

// Sirf

// Counter
// Flag
// Boolean
// Small integer

/*


Atomic vs Mutex

Atomic										Mutex
Fast										Thoda slow
Sirf simple operations						Complex logic
Counter										Data structures


Mutex vs Atomic

Mutex: Ek time mein sirf 1 goroutine.

Atomic: Ek time mein 1 operation.

Mathematical operations pe best.

*/

func main() {

	var counter int64

	// sync.WaitGroup ka use ek se zyada goroutines ko sync karne ke liye kiya jata hai.
	var wg sync.WaitGroup

	// 1000 goroutines run karega.
	for i := 0; i < 1000; i++ {

		wg.Add(1)

		go func() {

			defer wg.Done()

			// atomic.AddInt64(pointer, value)

			atomic.AddInt64(&counter, 1)

		}()

	}

	// wg.Wait() -> Ye goroutines ko wait karega ki khatam ho jaye.
	wg.Wait()

	// fmt.Println("Counter:", counter) -> Ye counter ko print karta hai.
	fmt.Println("Counter:", counter)

}
