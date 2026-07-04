package main

import "fmt"

// Closure ek anonymous function Closure wo function hai jo apne bahar ke variables ko "yaad" rakhta hai aur unhe baad me bhi use kar sakta hai.
func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}

}

func main() {
	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())
	name := "Aditya"

	show := func() {
		fmt.Println(name)
	}

	name = "Rahul"

	show()
}

// Closure variable ki copy nahi rakhta.

// Wo variable ka reference rakhta hai.

// Isliye latest value print hui.
