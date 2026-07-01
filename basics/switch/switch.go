package main

import "fmt"

func main() {

	// =====================
	// 1. Basic Switch - value match karo
	// =====================
	fmt.Println("--- 1. Basic Switch ---")
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Aaj Monday hai - Week ka pehla din")
	case "Friday":
		fmt.Println("Aaj Friday hai - Weekend aane wala hai!")
	case "Sunday":
		fmt.Println("Aaj Sunday hai - Rest karo")
	default:
		fmt.Println("Koi aur din hai")
	}

	// =====================
	// 2. Multiple values ek case mein
	// =====================
	fmt.Println("\n--- 2. Multiple Values in One Case ---")
	month := "June"
	switch month {
	case "December", "January", "February":
		fmt.Println("Winter season")
	case "March", "April", "May":
		fmt.Println("Spring season")
	case "June", "July", "August":
		fmt.Println("Summer season")
	default:
		fmt.Println("Autumn season")
	}

	// =====================
	// 3. Switch without expression (if-else jaisa)
	// =====================
	fmt.Println("\n--- 3. Switch Without Expression ---")
	age := 20
	switch {
	case age < 13:
		fmt.Println("Child")
	case age < 18:
		fmt.Println("Teenager")
	case age < 60:
		fmt.Println("Adult")
	default:
		fmt.Println("Senior")
	}

	// =====================
	// 4. Switch with short statement
	// =====================
	fmt.Println("\n--- 4. Switch with Short Statement ---")
	switch score := 85; {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

	// =====================
	// 5. fallthrough - next case bhi execute karo
	// =====================
	fmt.Println("\n--- 5. Fallthrough ---")
	num := 1
	switch num {
	case 1:
		fmt.Println("One")
		fallthrough // next case bhi chalega
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four") // yeh nahi chalega
	}

	// =====================
	// 6. Type Switch - variable ka type check karo
	// =====================
	fmt.Println("\n--- 6. Type Switch ---")
	values := []interface{}{42, "hello", true, 3.14}
	for _, v := range values {
		switch t := v.(type) {
		case int:
			fmt.Printf("%v is an int\n", t)
		case string:
			fmt.Printf("%v is a string\n", t)
		case bool:
			fmt.Printf("%v is a bool\n", t)
		default:
			fmt.Printf("%v is some other type\n", t)
		}
	}
}
