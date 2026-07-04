package main

import "fmt"

func greet() {
	fmt.Println("Welcome to Golang")
}

func calculate(a, b int) (int, int) {

	return a + b, a * b

}

func main() {

	greet()

	sum, product := calculate(10, 20)

	fmt.Println("sum", sum)
	fmt.Println("product", product)

}
