package main

// What if goroutines?
// Goroutine is a lightweight concurrent function managed by the Go runtime.
// Goroutine ek function hota hai jo background me independently execute hota hai. jo function execute krna hai uske starting me go likh do bus.

// Normal Function vs Goroutine

// Normal Function 							Goroutine
// Ek ke baad ek execute					Saath-saath execute ho sakta hai
// Blocking									Non-blocking
// Current thread par chalta hai 			Go runtime manage karta hai
// Heavy									Very lightweight

// Step By Step Guide to Create a Goroutine
// 1. Create a function
// 2. Call the function with the keyword 'go' in front of it

// main() start
// go hello()
// Go runtime ek nayi goroutine create karta hai
// main() aage continue karta hai
// hello() parallel me run hota hai
// Sleep ki wajah se main wait karta hai

// Agar Sleep hata dein?

// Possible Output: Kuch bhi nahi.

// Kyun?

// main function jaldi khatam ho jata hai aur program exit ho jata hai, goroutine ko run karne ka time nahi milta.

// Main function exit ho jaye to?

// Saari goroutines terminate ho jaati hain.

// Synchronization ke liye kya use karte hain?

// sync.WaitGroup, Channels, Mutexes.

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()               // Calling the function as a goroutine
	time.Sleep(1 * time.Second) // Sleep for 1 second to allow goroutine to finish
}

func sayHello() {
	fmt.Println("Hello from Goroutine!")
}
