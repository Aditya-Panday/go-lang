package main

// timeout pattern is used to implement timeout for a channel operation.
// It uses select with time.After.

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "From Channel"
	}()

	select {

	case msg := <-ch:
		fmt.Println("msg",msg)

	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}

}
