package main

import "fmt"

func main() {

	// =====================
	// 1. Basic for loop (C-style) - JavaScript for jaisa
	// =====================
	fmt.Println("--- 1. Basic For Loop ---")
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	// =====================
	// 2. While loop jaisa (sirf condition)
	// =====================
	fmt.Println("\n--- 2. While Loop Style ---")
	n := 1
	for n <= 5 {
		fmt.Println("n =", n)
		n++
	}

	// =====================
	// 3. Infinite loop (break se bahar nikalte hain)
	// =====================
	fmt.Println("\n--- 3. Infinite Loop with break ---")
	count := 0
	for {
		fmt.Println("count =", count)
		count++
		if count == 3 {
			break // loop se bahar niklo
		}
	}

	// =====================
	// 4. for range - Slice/Array pe loop (JS forEach jaisa)
	// =====================
	fmt.Println("\n--- 4. Range Loop on Slice ---")
	fruits := []string{"Apple", "Banana", "Mango", "Grapes"}
	for index, value := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// sirf value chahiye, index nahi
	fmt.Println("\n--- Range: only values ---")
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// =====================
	// 5. for range - Map pe loop
	// =====================
	fmt.Println("\n--- 5. Range Loop on Map ---")
	person := map[string]string{
		"name": "Rahul",
		"city": "Delhi",
		"job":  "Developer",
	}
	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// =====================
	// 6. for range - String pe loop (character by character)
	// =====================
	fmt.Println("\n--- 6. Range Loop on String ---")
	for index, char := range "Hello" {
		fmt.Printf("Index: %d, Char: %c\n", index, char)
	}

	// =====================
	// 7. continue - current iteration skip karo
	// =====================
	fmt.Println("\n--- 7. Continue (skip even numbers) ---")
	for i := 1; i <= 8; i++ {
		if i%2 == 0 {
			continue // even numbers skip karo
		}
		fmt.Println("Odd:", i)
	}

	// =====================
	// 8. Nested loops
	// =====================
	fmt.Println("\n--- 8. Nested Loops ---")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
	}
}
