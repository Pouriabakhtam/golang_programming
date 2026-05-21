package main

import (
	"bank/atm"
	"bank/userinput"
	"fmt"
)

const bank_info = "bank_info.txt"

func main() {
	amount := userinput.GetUserInput()
	fmt.Println("current balance:", atm.Balance(bank_info))
	atm.Deposit(amount, bank_info)
}
