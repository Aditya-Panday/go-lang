package main

import "fmt"

// A method is a function that is associated with a specific type (usually a struct).

// Method ek function hota hai jo kisi Struct ya Type ke saath attach hota hai.

// difference between a function and a method:
// 1. Function: A function is a standalone block of code that can be called independently.
// 2. Method: A method is a function that is associated with a specific type (usually a struct) and can be called on instances of that type.

// function example:
// Ye kisi se attached nahi hai.
func add(a, b int) int {
	return a + b
}

// method example:
// Ye User struct ke saath attached hai.
// func (u User) PrintName() {
// 	fmt.Println(u.Name)
// }

// different types of methods in Go:
// 1. Value Receiver Method: A method that operates on a copy of the value it is called on.
// 2. Pointer Receiver Method: A method that operates on the original value it is called on, allowing it to modify the value.

// difference between value receiver and pointer receiver methods:
// 1. Value Receiver Method: It operates on a copy of the value, so any changes made to the value inside the method do not affect the original value.
// 2. Pointer Receiver Method: It operates on the original value, so any changes made to the value inside the method will affect the original value.

// Example of Value Receiver Method:
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// Example of Pointer Receiver Method:
func (c *Circle) SetRadius(r float64) {
	c.radius = r
}

//

type Student struct {
	Name string

	Marks int
}

func (s Student) Print() {

	fmt.Println(s.Name)

}

func (s Student) Grade() {
	if s.Marks >= 90 {
		fmt.Println("Grade A")
	} else {
		fmt.Println("Grade B")
	}
}

func main() {

	s := Student{

		Name: "Aditya",

		Marks: 95,
	}

	s.Print()

	s.Grade()

}
