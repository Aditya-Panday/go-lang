package main

// errgroup runs multiple goroutines and stops all if one returns an error.

// Easy language

// Suppose

// 3 APIs

// User API

// Order API

// Payment API

// Agar

// Payment API fail

// To

// Baaki sab cancel.

// Ye errgroup karta hai.

import (
	"context"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		time.Sleep(1 * time.Second)
		return nil
	})
	g.Go(func() error {
		time.Sleep(2 * time.Second)
		return nil
	})
	g.Wait()
}