package userinput

import (
	"fmt"
	"os"
	|
)

func GetUserInput() (amount float64) {
	for {
		fmt.Println("Welcome to the ATM Machine!")
		fmt.Println("Please enter the amount you wish to deposit (or 0 to exit):")
		fmt.Print("Enter the amount to deposit: ")
		_, err := fmt.Scanf("%f", &amount)
		if err != nil {
			fmt.Println("Error reading user input. Please enter a valid number.")
			continue
		}
		if amount == 0 {
			fmt.Println("Exiting program.")
			os.Exit(0)
		}
		if amount < 0 {
			fmt.Println("Amount must be greater than zero. Please try again.")
			continue
		}
		if amount > 10000 {
			fmt.Println("Amount exceeds the maximum limit of 10000. Please try again.")
			continue
		}
		return amount
	}
}
