package main

// Enum -----> An enum (enumeration) is a type that represents a fixed set of constant values.    // Enum ka use tab karte hain jab kisi variable ki sirf kuch fixed values hi valid hon.
// Example:-
// Order Status

// What is iota?
// Automatic incrementing constant generator.

// Can we make methods on enums in Go?
// Yes, we can define methods on enums in Go by defining a custom type and associating methods with that type.

// Pending
// Processing
// Shipped
// Delivered
// Cancelled

// IMportant points about enums in Go:
// 1. Go does not have a built-in enum type like some other programming languages.
// Lekin Go me typed constants (const) + iota ki help se enums banaye jaate hain.
// 2. Instead, Go uses constants and iota to create enumerated values.
// 3. The iota identifier is used to simplify the creation of incrementing numbers for constants.
// 4. Enums in Go are typically defined using a custom type and a set of constants.

import "fmt"

// Example of defining an enum using iota:
type Status int

const (
	Pending Status = iota
	// For Custom Starting Value) ----> 	Pending = iota + 1
	Processing
	Shipped
	Delivered
)

// Bit Flags)

// Ye networking aur permissions me bahut use hota hai.
const (
	Read = 1 << iota
	Write
	Execute
)

func main() {
	// Status Enum
	fmt.Println(Pending)
	fmt.Println(Processing)
	fmt.Println(Shipped)
	fmt.Println(Delivered)

	// Bit Flags
	fmt.Println(Read)
	fmt.Println(Write)
	fmt.Println(Execute)

}
