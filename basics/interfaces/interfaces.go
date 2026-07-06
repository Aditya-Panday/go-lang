package main

// Methods aur Interfaces milkar Go me Object-Oriented Programming ka foundation banate hain.

// Interfaces
// An interface is a type that defines a set of method signatures. Any type that implements those methods automatically satisfies the interface.
// Simple language
// Interface ek contract hota hai jo batata hai ki kisi type ke paas kaun-kaun se methods hone chahiye.

// Go me interface ko explicitly implement nahi karna padta. Agar struct ke paas required methods hain, to wo automatically interface implement kar deta hai.

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}

func (Dog) Speak() {
	fmt.Println("Woof Woof")
}

type Cat struct{}

func (Cat) Speak() {
	fmt.Println("Meow Meow")
}

func makeSpeak(s Speaker) {

	s.Speak()

}

func main() {

	var s Speaker // Ye actual object nahi hai  Ye sirf kisi bhi type ko hold kar sakta hai jo Speaker interface ko implement karta ho.

	s = Dog{}

	s.Speak()

	s = Cat{}

	s.Speak()

	makeSpeak(Dog{})

	makeSpeak(Cat{})

	// Empty Interface
	// interface{}

	var x interface{} // Ye ek empty interface hai jo kisi bhi type ko hold kar sakta hai.

	x = 10

	x = "Hello"

	x = true

	x = Dog{}

	fmt.Println(x)

}
