package main

import (
	"fmt"
	"time"
)

// Channels

// Ab hum Go ke sabse important topic par aa gaye hain.

// Agar kisi interview me Go ki strength puchi jaye, to answer hoga:

// Goroutines + Channels

// Go ka famous slogan bhi hai:

// "Don't communicate by sharing memory, share memory by communicating."

// Iska matlab hai:

// Variables ko multiple goroutines ke beech directly share mat karo. Data ko Channels ke through pass karo.

// Defination
// A Channel is a communication pipe that allows goroutines to send and receive data safely.

// Socho tumhare paas ek delivery pipe hai.

// Worker 1  ─────📦─────► Worker 2
// Worker 1 package bhejta hai.

// Worker 2 receive karta hai.

// Ye package hi data hai.

// Channel = Pipe

// Why Channels?

// Without channel

// Goroutine A

// ↓

// Shared Variable

// ↓

// Goroutine B

// Problem:

// Race Condition
// Data Corruption
// Synchronization Issues

// Channel Syntax

// Channel banana

// ch := make(chan int)

// send
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing number", num)
		time.Sleep(time.Second)
	}

}
// receive
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

func main() {

// create a channel
// messageChain := make(chan string)

// // send a message to channel
// messageChain <- "ping" // sending is a blocking operation.

// // receive a message from channel // reciveing is also a blocking operation
// m:= <-messageChain
// fmt.Println("m",m)


	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.Intn(100)
	// }

	

	result := make(chan int)
	go sum(result, 4, 5)
	res := <-result // blocking

	fmt.Println(res)



}

