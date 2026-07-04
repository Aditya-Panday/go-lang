package main

import (
	"fmt"
	"maps"
	"strings"
)

// 1. Difference between Slice and Map?
// Slice							Map
// Index based						Key based
// Ordered							Unordered
// Duplicate values 				allowed	Keys must be unique
// Access by index					Access by key
// Good for lists					Good for fast lookup
// maps -> hash, 					object, dict
func main() {
	// creating map

	k := make(map[string]string)

	// setting an element
	k["name"] = "golang"
	k["area"] = "backend"

	// get an element
	fmt.Println(k["name"], k["area"])
	// IMP: if key does not exists in the map then it returns zero value

	m := make(map[string]int)
	m["age"] = 30
	m["price"] = 50
	fmt.Println(m["phone"])
	fmt.Println(len(m))

	delete(m, "price")
	clear(m)

	fmt.Println("first m", m)
	fmt.Println("first m2", m)

	d := map[string]int{"price": 40, "phones": 3}

	v, ok := d["phones"]
	fmt.Println(v)
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 8}
	fmt.Println(maps.Equal(m1, m2))
	// user data
	users := map[string]map[string]string{

		"user1": {
			"name": "Aditya",
			"city": "Delhi",
		},

		"user2": {
			"name": "Rahul",
			"city": "Noida",
		},
	}

	fmt.Println(users["user1"]["name"])

	// data access
	students := map[string][]string{

		"Aditya": {"Go", "Node", "React"},
		"Rahul":  {"Java", "Spring"},
	}

	fmt.Println(students["Aditya"])

	// advanced
	text := "go is fast go is powerful go"

	words := strings.Fields(text)
	// fmt.Println("words", words)

	counter := make(map[string]int)

	for _, word := range words {

		counter[word]++

	}

	fmt.Println(counter)

}
