package main

// Variadic Function wo function hota hai jo fixed number ki jagah variable number of arguments accept kar sakta hai.

import "fmt"

func sum(numbers ...int) {

	fmt.Println(numbers)

}
func log(messages ...string) {

	for _, msg := range messages {

		fmt.Println(msg)

	}

}

func main() {

	sum(10)
	sum(10, 20)
	sum(10, 20, 30)
	sum(10, 20, 30, 40)

	log("Server Started")

	log(
		"Connected Database",
		"User Login",
		"Payment Success",
	)

}

// Important Points
// Point						Explanation
// ...type						Variable number of arguments accept karta hai
// Inside function				Variadic parameter ek slice hota hai
// Position						Variadic parameter hamesha last me hota hai
// Slice pass karna				slice... likhna padta hai

// Easy Trick to Remember

// Socho ... ka matlab hai:

// "Jitni values bhejni hain bhej do, Go un sabko automatically ek slice me convert kar dega."
