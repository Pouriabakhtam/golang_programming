package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)
	agepointer := &age

	// fmt.Println("Age address:", &age)
	// fmt.Println("Age pointer address:", agepointer)
	// fmt.Println("Age value using pointer:", *agepointer)

	// *agepointer = 35
	// fmt.Println("updated age pointer address:", agepointer)
	// fmt.Println("updated age value using pointer:", *agepointer)

	println(checkage(agepointer))
}

func checkage(agepointer *int) bool {
	if *agepointer >= 18 {
		return true
	} else {
		return false
	}
}
