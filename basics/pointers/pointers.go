package main

// A pointer is a variable that stores the memory address of another variable.

// Pointer ek variable hota hai jo kisi doosre variable ki value nahi, balki uska memory address store karta hai.

// first learn what is memory address and how to get it in golang.

import "fmt"

func swap(a, b *int) {

	*a, *b = *b, *a

}

func main() {
	var a int = 10
	var b *int = &a // b is a pointer to an integer, and it stores the memory address of variable a & provides memmory address.

	fmt.Println("Value of a:", a)                       // prints the value of a
	fmt.Println("Memory address of a:", &a)             // prints the memory address of a
	fmt.Println("Value of b (memory address of a):", b) // prints the memory address stored in b
	fmt.Println("Value pointed to by b:", *b)           // dereferencing b to get the value of a

	// Learn Dereference  (Pointer se original value kaise nikale?)
	name := "Aditya"

	ptr := &name

	fmt.Println(*ptr)

	// Pointer se Value Change Karna
	*ptr = "Rahul"

	fmt.Println("name", name)

	// swap program
	x := 10
	y := 20

	swap(&x, &y)

	fmt.Println(x, y)

	// New keyword in golang
	// New keyword ek built-in function hai jo ek naya variable create karta hai aur uska pointer return karta hai.
	ptra := new(int)

	fmt.Println(*ptra)

	*ptra = 100

	fmt.Println(*ptra)

}

// Pointer pass karne ka benefit?

// Memory bachti hai (large data copy nahi hota)
// Original value modify kar sakte ho
// Performance improve hoti hai

// &  → Address nikaalo

// *  → Address se value nikalo
