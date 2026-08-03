package main

import (
	"fmt"
	"sync"
)

// A Mutex (Mutual Exclusion) is used to ensure that only one goroutine can access a shared resource at a time.
// Mutex ek lock hai. Agar ek goroutine kisi data ko use kar rahi hai, to doosri goroutine ko wait karna padega.

// Maan lo

// counter := 0

// 100 goroutines

// counter++

// kar rahi hain.

// Tum expect karoge

// 100

// Lekin output ho sakta hai

// 78

// 82

// 95

// 99

// Kyun?

// counter++ Actually Kya Hota Hai?

// Hum sochte hain

// counter++

// Ek operation hai.

// Reality:

// Step 1 → Read counter

// Step 2 → +1

// Step 3 → Write back

// Yani 3 alag operations.

// Suppose

// Counter = 5

// Worker A

// Read → 5

// Worker B

// Read → 5

// Worker A

// Write → 6

// Worker B

// Write → 6

// Expected

// 7

// Actual

// 6

// Ek increment lose ho gaya.

// Isi ko kehte hain

// Race Condition

// Race Condition Definition

// A race condition occurs when multiple goroutines access and modify the same data simultaneously, leading to unpredictable results.

// Easy language

// Jab do ya zyada goroutines ek hi variable ko ek saath modify karti hain.

// Mutex Solution
// Worker A

// ↓

// LOCK

// ↓

// counter++

// ↓

// UNLOCK

// ↓

// Worker B

// ↓

// LOCK

// ↓

// counter++

// ↓

// UNLOCK

// Ab koi race nahi.



var counter int

var mu sync.Mutex

func increment(wg *sync.WaitGroup) {

	defer wg.Done()

	mu.Lock()

	counter++

	mu.Unlock()

}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {

		wg.Add(1)

		go increment(&wg)

	}

	wg.Wait()

	fmt.Println(counter)

}

// Lock & Unlock

// Ye do methods yaad rakhna.

// mu.Lock()

// Matlab

// Gate Close
// mu.Unlock()

// Matlab

// Gate Open
// Visualization


