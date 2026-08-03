package main

// Deadlock occurs when goroutines wait forever for each other.

// Sab wait kar rahe hain aur koi aage nahi badh raha.

/*

Deadlock Kab Hota Hai?

Jab 2+ goroutines ek dusre ka lock release hone ka wait karti hain.

Example:
Goroutine 1 lock A
Goroutine 2 lock B

Goroutine 1 wait kare B ke liye
Goroutine 2 wait kare A ke liye

Solution:

Order Maintain karo.

Hamesha A → B nahi toh B → A
Kabbi bhi A → B aur B → A nahi.

*/

// Example

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var m1 sync.Mutex
	var m2 sync.Mutex

	go func() {
		m1.Lock()
		time.Sleep(1 * time.Second)
		m2.Lock()
		fmt.Println("Goroutine 1")
		m2.Unlock()
		m1.Unlock()
	}()

	go func() {
		m2.Lock()
		time.Sleep(1 * time.Second)
		m1.Lock()
		fmt.Println("Goroutine 2")
		m1.Unlock()
		m2.Unlock()
	}()

}

