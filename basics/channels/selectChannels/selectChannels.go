package main

import (
	"fmt"
	"time"
)

// select waits on multiple channel operations and executes the one that becomes ready first.
// select ek saath multiple channels ko listen karta hai. Jo channel pehle ready hota hai, uska case execute ho jata hai.

// Without select

// Suppose do channels hain.

// fmt.Println(<-ch1)
// fmt.Println(<-ch2)

// Problem:

// Agar ch1 ready nahi hai, to program block ho jayega.

// Chahe ch2 ready ho.


func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "From Channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "From Channel 2"
	}()

	select {

	case msg := <-ch1:
		fmt.Println("msg",msg)

	case msg := <-ch2:
		fmt.Println("msg",msg)
	}

}