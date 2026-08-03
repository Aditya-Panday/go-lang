package main

// Context carries deadlines, cancellation signals, and request-scoped values across API boundaries.
// Context ek signal hai jo bolta hai ki "ab kaam band karo".

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	go longRunningTask(ctx)

	// Wait for the task to finish or timeout
	<-ctx.Done()

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Task timed out")
	}
}

func longRunningTask(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second): // Task takes 5 seconds
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Println("Task cancelled")
	}
}


func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped\n", id)
			return
		default:
			fmt.Printf("Worker %d working...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}
