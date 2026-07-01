package main

import "fmt"

// Struct - Go mein "object" ka equivalent
type Person struct {
	Name string
	Age  int
	City string
}

func main() {

	// =====================
	// Basic Data Types
	// =====================
	var i int = 42
	var f float64 = 3.14
	var b bool = true
	var s string = "Go Lang"
	var r rune = 'A'         // Unicode character (int32)
	var by byte = 255        // uint8
	var c complex64 = 3 + 4i // complex number

	fmt.Println("--- Basic Types ---")
	fmt.Println("int:", i)
	fmt.Println("float64:", f)
	fmt.Println("bool:", b)
	fmt.Println("string:", s)
	fmt.Println("rune:", r, "->", string(r))
	fmt.Println("byte:", by)
	fmt.Println("complex:", c)

	// =====================
	// short Hand Syntax
	// =====================

	name := "golang"
	println(name)

	// =====================
	// Array - fixed size
	// =====================
	fmt.Println("\n--- Array (fixed size) ---")
	var arr [3]string = [3]string{"Apple", "Banana", "Mango"}
	fmt.Println("Array:", arr)
	fmt.Println("Array[0]:", arr[0])
	fmt.Println("Array length:", len(arr))

	// =====================
	// Slice - dynamic size (mostly yahi use karte hain)
	// =====================
	fmt.Println("\n--- Slice (dynamic array) ---")
	fruits := []string{"Apple", "Banana", "Mango"}
	fruits = append(fruits, "Grapes") // element add karo
	fmt.Println("Slice:", fruits)
	fmt.Println("Slice[1]:", fruits[1])
	fmt.Println("Slice length:", len(fruits))

	// slice of ints
	nums := []int{10, 20, 30, 40, 50}
	fmt.Println("Sub-slice:", nums[1:4]) // index 1 se 3 tak

	// =====================
	// Map - key:value pairs (JS object jaisa)
	// =====================
	fmt.Println("\n--- Map (key-value) ---")
	person := map[string]string{
		"name": "Rahul",
		"city": "Delhi",
	}
	fmt.Println("Map:", person)
	fmt.Println("Name:", person["name"])

	// map mein value add karo
	person["email"] = "rahul@example.com"
	fmt.Println("After add:", person)

	// map mein key check karo
	val, exists := person["phone"]
	fmt.Println("phone exists?", exists, "| value:", val)

	// =====================
	// Struct - custom object jaisa
	// =====================
	fmt.Println("\n--- Struct (object) ---")
	p1 := Person{
		Name: "Amit",
		Age:  28,
		City: "Mumbai",
	}
	fmt.Println("Struct:", p1)
	fmt.Println("Name:", p1.Name)
	fmt.Println("Age:", p1.Age)

	// struct update
	p1.City = "Pune"
	fmt.Println("Updated city:", p1.City)

	// slice of structs (array of objects jaisa)
	fmt.Println("\n--- Slice of Structs ---")
	people := []Person{
		{Name: "Rohit", Age: 22, City: "Jaipur"},
		{Name: "Priya", Age: 25, City: "Bangalore"},
	}
	for _, p := range people {
		fmt.Printf("%s is %d years old from %s\n", p.Name, p.Age, p.City)
	}

	println("const groupping----")
	println("Server running at", host, "on port", port)
}

// Important rules Go mein:

// := se declare kiya variable use karna zaruri hai, warna compile error aata hai
// var se global level pe bhi declare kar sakte ho, := sirf function ke andar
// const mein type likhna optional hai, Go khud detect kar leta hai
// const ko aap function ke bahar bhi declare kr skte ho but := shorthand ko nhi kr skte ho.

const age = 30

// constant groupping
const (
	port = 5000
	host = "localhost:3000"
)
