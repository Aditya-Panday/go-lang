package main

import (
	"fmt"
	"slices"
)

func main() {

	// 1. Slice banana - 3 tarike
	fmt.Println("--- 1. Slice Banana ---")
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Literal:", s1)

	s2 := make([]int, 5)
	fmt.Println("make():", s2)

	s3 := make([]int, 3, 10)
	fmt.Printf("make(len,cap): %v | len=%d cap=%d\n", s3, len(s3), cap(s3))

	// 2. append() - elements add karo
	fmt.Println("\n--- 2. append() ---")
	nums := []int{1, 2, 3}
	nums = append(nums, 4)
	nums = append(nums, 5, 6, 7)
	fmt.Println("After append:", nums)

	extra := []int{8, 9, 10}
	nums = append(nums, extra...)
	fmt.Println("After append slice:", nums)

	// 3. len() and cap()
	fmt.Println("\n--- 3. len() and cap() ---")
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	// 4. Sub-slicing
	fmt.Println("\n--- 4. Sub-Slicing ---")
	a := []int{10, 20, 30, 40, 50}
	fmt.Println("Original:  ", a)
	fmt.Println("a[1:3] -> ", a[1:3])
	fmt.Println("a[:3]  -> ", a[:3])
	fmt.Println("a[2:]  -> ", a[2:])

	// 5. copy() - deep copy
	fmt.Println("\n--- 5. copy() ---")
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copy(dst, src)
	dst[0] = 999
	fmt.Println("src (unchanged):", src)
	fmt.Println("dst (changed):  ", dst)

	// 6. slices.Contains()
	fmt.Println("\n--- 6. slices.Contains() ---")
	fruits := []string{"Apple", "Banana", "Mango"}
	fmt.Println("Contains Apple: ", slices.Contains(fruits, "Apple"))
	fmt.Println("Contains Grapes:", slices.Contains(fruits, "Grapes"))

	// 7. slices.Index()
	fmt.Println("\n--- 7. slices.Index() ---")
	fmt.Println("Index of Banana:", slices.Index(fruits, "Banana"))
	fmt.Println("Index of Grapes:", slices.Index(fruits, "Grapes"))

	// 8. slices.Reverse()
	fmt.Println("\n--- 8. slices.Reverse() ---")
	r := []int{1, 2, 3, 4, 5}
	slices.Reverse(r)
	fmt.Println("Reversed:", r)

	// 9. slices.Sort()
	fmt.Println("\n--- 9. slices.Sort() ---")
	unsorted := []int{5, 2, 8, 1, 9, 3}
	slices.Sort(unsorted)
	fmt.Println("Sorted ints:   ", unsorted)

	strs := []string{"Banana", "Apple", "Mango"}
	slices.Sort(strs)
	fmt.Println("Sorted strings:", strs)

	// 10. slices.Max() and slices.Min()
	fmt.Println("\n--- 10. slices.Max() and slices.Min() ---")
	numbers := []int{3, 1, 9, 2, 7, 5}
	fmt.Println("Max:", slices.Max(numbers))
	fmt.Println("Min:", slices.Min(numbers))

	// 11. Delete element (middle se)
	fmt.Println("\n--- 11. Delete Element ---")
	d := []string{"A", "B", "C", "D", "E"}
	d = append(d[:2], d[3:]...)
	fmt.Println("After delete index 2:", d)

	// 12. Insert element (middle mein)
	fmt.Println("\n--- 12. Insert Element ---")
	ins := []int{1, 2, 4, 5}
	ins = append(ins[:2], append([]int{3}, ins[2:]...)...)
	fmt.Println("After insert 3 at index 2:", ins)

	// 13. 2D Slice
	fmt.Println("\n--- 13. 2D Slice ---")
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, row := range matrix {
		fmt.Println(row)
	}
}
