package main

import "fmt"

// Generics Go 1.18 me introduce hue the.
// Generics ki help se hum ek hi function ya data structure ko different data types ke saath use kar sakte hain, bina duplicate code likhe.

// Use cases of Generics:
// 1. Reusable Data Structures: Generics allow you to create reusable data structures that can work with different types of data. For example, you can create a generic linked list or stack that can store elements of any type.
// 2. Type-Safe Collections: Generics enable you to create type-safe collections, such as slices or maps, that can hold elements of a specific type. This helps catch type-related errors at compile time.
// 3. Generic Functions: You can define generic functions that can operate on different types of data, allowing for code reuse and reducing duplication.
// 4. Improved Code Readability: Generics can improve code readability by making it clear what types are expected and used in a function or data structure.
// 5. Performance Optimization: Generics can help optimize performance by avoiding unnecessary type conversions and allowing the compiler to generate specialized code for specific types.

// Maan lo tumhe do numbers ka maximum nikalna hai.

func Print[T any](value T) { // aab iske liye diffrent different functions likhne ki zarurat nahi hai. Hum ek hi function ko different data types ke saath use kar sakte hain.

	fmt.Println(value)

}

type Number interface {
	int
}

func Prints[T Number](value T) { // aab iske liye diffrent different functions likhne ki zarurat nahi hai. Hum ek hi function ko different data types ke saath use kar sakte hain.

	fmt.Println(value)

}

// ~ ka matlab hai: "Underlying type ye hona chahiye." bolta hai ki sirf exact type hi nahi, balki us type se bane hue custom types bhi allowed hain.

func main() {

	Print(100)

	Print("Aditya")

	Print(true)

	Print(3.14)
	// chal jayega kyunki 10 int hai
	Prints(10)
	// 	Without ~

	// Teacher bolta hai

	// Sirf original naam "int" wale andar aa sakte hain.

	// 	With ~

	// Teacher bolta hai

	// Jiska original base type int hai, sab aa jao

}
