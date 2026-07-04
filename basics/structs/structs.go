package main

import "fmt"

// Structs Go ke sabse important concepts me se ek hain, kyunki real-world backend development me lagbhag har jagah Structs use hote hain.

// API Requests/Responses
// Database Models
// JSON
// Authentication
// Microservices
// ORM Libraries
// gRPC

// Struct is a user-defined data type that groups multiple related fields of different data types into a single unit.
// Struct ek custom data type hai jo related data ko ek object ki tarah store karta hai.

// Real Life Example

// Maan lo ek Student hai.

// Uske paas:

// Name
// Age
// City
// Marks

// Agar Struct na ho to

// name := "Aditya"
// age := 25
// city := "Delhi"
// marks := 95

// Ye sab alag-alag variables hain.

// Struct me

// Student
// │
// ├── Name
// ├── Age
// ├── City
// └── Marks

type Student struct {
	Name string
	Age  int
	City string
}

type User struct {
	Name string `json:"name"`

	Age int `json:"age"`
}

// Student

// ↓

// Custom Data Type

func updateAge(s *Student) {

	s.Age = 30

}

func main() {

	var s Student
	fmt.Println("s--->", s.Name)

	s.Name = "Aditya"
	s.Age = 25
	s.City = "Delhi"
	// s.Group = "A"

	fmt.Println(s)

	// (Struct Literal)
	s2 := Student{
		Name: "Aditya",
		Age:  25,
		City: "Delhi",
	}

	fmt.Println(s2)

	// Struct with pointer
	st := Student{
		Name: "Aditya",
		Age:  25,
	}
	// fmt.Println("st--->", st)

	updateAge(&st)

	fmt.Println("st--->", st)

	// Anonymous Struct
	user := struct {
		Name string
		Age  int
	}{

		Name: "Aditya",
		Age:  25,
	}

	fmt.Println(user)

	// Struct Tags (Very Important)
	// Ye tags batate hain ki field JSON, database, XML, validation, etc. me kis naam se map hogi.

}

// What is struct.
// Struct is a custom data type used to group related fields into a single unit.

// Struct aur Map  difference?
// Struct → fixed fields.
// Map → dynamic key-value.

// Struct ko function me pointer se kyun pass karte hain?
// Copy avoid hoti hai.
// Performance improve hoti hai.
// Original object modify kar sakte hain.

// Struct Embedding
// Struct embedding ek aisa feature hai jisme ek struct ke andar dusre struct ko embed kar sakte hain.

// Isse hum code reuse kar sakte hain aur ek struct ke andar dusre struct ke fields aur methods ko access kar sakte hain.

// Example:

// type Address struct {
// 	City  string
// 	State string
// }

// type User struct {
// 	Name    string
// 	Address Address
// }

// user := User{
// 	Name: "Aditya",
// 	Address: Address{
// 		City:  "Delhi",
// 		State: "Delhi",
// 	},
// }

// fmt.Println(user.Address.City)

// Method Overiding
// Go me method overriding ka concept nahi hai, lekin hum struct embedding ke through similar behavior achieve kar sakte hain.

// Example:

// type Animal struct {}

// func (a Animal) Speak() {
// 	fmt.Println("Animal speaks")
// }

// type Dog struct {
// 	Animal
// }

// func (d Dog) Speak() {
// 	fmt.Println("Dog barks")
// }

// dog := Dog{}
// dog.Speak() // Output: Dog barks
// dog.Animal.Speak() // Output: Animal speaks
