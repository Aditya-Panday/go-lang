package main

// Worker Pool is a pattern where multiple workers consume jobs from a shared channel.
// It is used to process multiple jobs in parallel.
// Ek queue hoti hai jisme jobs aati hain aur multiple workers us queue se jobs utha kar process karte hain.

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second) // simulate work
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= 5; a++ {
		fmt.Println("result", <-results)
	}
}