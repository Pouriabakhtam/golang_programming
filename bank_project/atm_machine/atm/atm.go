package atm

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func Deposit(amount float64, bank_info string) {
	if amount <= 0 {
		panic("Amount must be greater than zero")
	} else {
		fmt.Printf("Depositing amount: %.2f\n", amount)
		file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		entry := fmt.Sprintf("%v - user has deposited %.2f to the account\n", time.Now(), amount)
		_, err = file.WriteString(entry)
		if err != nil {
			panic(err)
		}
		// os.WriteFile("log.txt", []byte(fmt.Sprintf("user has already deposited %v to the account", amount)), 0644)
	}
}

func Balance(bank_info string) float64 {
	data, err := os.ReadFile(bank_info)
	if err != nil {
		panic("Error reading bank info or there is no account")
	} else {
		balance, err := strconv.ParseFloat(string(data), 64)
		if err != nil {
			panic("Error parsing balance")
		}
		return balance
	}
}
