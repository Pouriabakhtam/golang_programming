package main

import (
	"fmt"
	"math"
)

func main() {
	// here I did define variables in different way, you can also define them in one line as shown below
	// investmentAmount := 10000.0
	// expectedReturnRate := 5.5
	// years := 10.0
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64
	const inflattionRate = 2.0
	fmt.Print("Enter the initial investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+(expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+(inflattionRate/100), years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
