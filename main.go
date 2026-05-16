package main

import (
	"application_pouri/calculator"
	"application_pouri/rand_numb"
	"fmt"
)

func main() {
	fmt.Println(rand_numb.Random_n(120))
	for {
		fmt.Println("Welcome to the Calculator!")
		fmt.Println("Please select an operation:")
		fmt.Println("+ -> Addition")
		fmt.Println("- -> Subtraction")
		fmt.Println("* -> Multiplication")
		fmt.Println("/ -> Division")
		fmt.Println("q -> Exit")

		fmt.Print("Enter your choice: ")
		var choice string
		fmt.Scan(&choice)

		if choice == "q" {
			fmt.Println("Exiting the calculator. Goodbye!")
			break
		} else {
			fmt.Print("Enter the first number: ")
			var num1 float64
			fmt.Scan(&num1)
			fmt.Print("Enter the second number: ")
			var num2 float64
			fmt.Scan(&num2)
			if choice == "+" {
				result := calculator.Add(num1, num2)
				fmt.Println("The result of addition is:", result)
			} else if choice == "-" {
				result := calculator.Subtract(num1, num2)
				fmt.Println("The result of subtraction is:", result)
			} else if choice == "*" {
				result := calculator.Multiply(num1, num2)
				fmt.Println("The result of multiplication is:", result)
			} else if choice == "/" {
				result, err := calculator.Divide(num1, num2)
				if err != nil {
					fmt.Println("Error:", err)
				} else {
					fmt.Println("The result of division is:", result)
				}
			}
		}
	}
}
