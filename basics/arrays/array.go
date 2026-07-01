package main

import "fmt"

func main() {

	// =====================
	// 1. Basic Array - fixed size
	// =====================
	fmt.Println("--- 1. Basic Array ---")
	var arr [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array:", arr)
	fmt.Println("Length:", len(arr))
	fmt.Println("First element:", arr[0])
	fmt.Println("Last element:", arr[4])

	// =====================
	// 2. Short declaration
	// =====================
	fmt.Println("\n--- 2. Short Declaration ---")
	fruits := [3]string{"Apple", "Banana", "Mango"}
	fmt.Println("Fruits:", fruits)

	// ... se size auto calculate
	colors := [...]string{"Red", "Green", "Blue", "Yellow"}
	fmt.Println("Colors:", colors)
	fmt.Println("Colors length:", len(colors))

	// =====================
	// 3. Array update karna
	// =====================
	fmt.Println("\n--- 3. Update Array Element ---")
	fruits[1] = "Grapes"
	fmt.Println("After update:", fruits)

	// =====================
	// 4. Array loop karna
	// =====================
	fmt.Println("\n--- 4. Loop Over Array ---")
	for i, v := range fruits {
		fmt.Printf("Index %d: %s\n", i, v)
	}

	// =====================
	// 5. 2D Array (matrix)
	// =====================
	fmt.Println("\n--- 5. 2D Array ---")
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:", matrix)
	fmt.Println("Row 0:", matrix[0])
	fmt.Println("Element [1][2]:", matrix[1][2])

	// 2D array loop
	fmt.Println("Full matrix:")
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}

	// =====================
	// 6. Slice - dynamic array (mostly yahi use hota hai)
	// =====================
	fmt.Println("\n--- 6. Slice (Dynamic Array) ---")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", nums)
	fmt.Println("Length:", len(nums))
	fmt.Println("Capacity:", cap(nums))

	// append - element add karo
	nums = append(nums, 6, 7)
	fmt.Println("After append:", nums)

	// =====================
	// 7. Slice operations
	// =====================
	fmt.Println("\n--- 7. Slice Operations ---")
	s := []int{10, 20, 30, 40, 50}

	fmt.Println("Original:", s)
	fmt.Println("s[1:3] ->", s[1:3]) // index 1 se 2 tak
	fmt.Println("s[:3]  ->", s[:3])  // shuru se index 2 tak
	fmt.Println("s[2:]  ->", s[2:])  // index 2 se end tak

	// =====================
	// 8. make() se slice banana
	// =====================
	fmt.Println("\n--- 8. make() Slice ---")
	made := make([]int, 5)    // length 5, zero values
	fmt.Println("make slice:", made)

	made2 := make([]int, 3, 10) // length 3, capacity 10
	fmt.Println("make with cap:", made2)
	fmt.Println("len:", len(made2), "cap:", cap(made2))

	// =====================
	// 9. copy() - slice copy karna
	// =====================
	fmt.Println("\n--- 9. Copy Slice ---")
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println("Source:", src)
	fmt.Println("Destination:", dst)

	// dst change karne se src affect nahi hoga
	dst[0] = 999
	fmt.Println("After dst change:")
	fmt.Println("src:", src)
	fmt.Println("dst:", dst)

	// =====================
	// 10. Slice of Structs (array of objects jaisa)
	// =====================
	fmt.Println("\n--- 10. Slice of Structs ---")
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Rahul", 25},
		{"Priya", 22},
		{"Amit", 28},
	}

	for _, p := range people {
		fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
	}
}
