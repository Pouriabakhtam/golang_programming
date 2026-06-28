package main

import (
	"fmt"

	"tax_calc.ca/filemanager"
	"tax_calc.ca/prices"
)

func main() {
	taxRate := []float64{0.0, 0.5, 0.7}
	donechannel := make([]chan bool, len(taxRate))
	for index, taxvalues := range taxRate {
		donechannel[index] = make(chan bool)
		fm := filemanager.New("price.txt", fmt.Sprintf("result_%0f.json", taxvalues*100))
		//cmd := cmdmanager.New()
		testtax := prices.New(fm, taxvalues)
		go testtax.Process(donechannel[index])
	}

	for _, done := range donechannel {
		<-done
	}
}
