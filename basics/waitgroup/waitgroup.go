package main

import (
	"fmt"
	"sync"
)

// WaitGroup
// A WaitGroup is used to wait for a collection of goroutines to finish before the main function exits.

// WaitGroup ek counter ki tarah kaam karta hai jo batata hai kitni goroutines abhi complete honi baaki hain. Main function tab tak wait karta hai jab tak counter 0 na ho jaye.

// func printHello() {
// 	fmt.Println("Hello")
// }

// func main() {

// 	go printHello()

// 	fmt.Println("Main Finished")

// }

// Possible output ----> Main Finished
// Kyun?

// main()

// ↓

// go printHello()

// ↓

// Main function khatam

// ↓

// Program Exit

// ↓

// Goroutine ko chance hi nahi mila

// WaitGroup use karne ke steps:

// 1. Import sync package
// 2. Create a WaitGroup variable
// 3. Add the number of goroutines to WaitGroup (Add)
// 4. Call WaitGroup.Done() at the end of each goroutine
// 5. Call WaitGroup.Wait() in the main function

// func printHy(wg *sync.WaitGroup) {

// 	defer wg.Done() // WaitGroup ka counter 1 se decrease krta hai

// 	fmt.Println("Hello")

// }

// func main() {

// 	var wg sync.WaitGroup

// 	wg.Add(1)

// 	go printHy(&wg)

// 	wg.Wait()

// 	fmt.Println("Main Finished")

// }


func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("Worker", id, "Finished")

}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {

		wg.Add(1)

		go worker(i, &wg)

	}

	wg.Wait()

	fmt.Println("All Workers Completed")

}
