package conversion

import (
	"strconv"
)

func Conversion(strslice []string) ([]float64, error) {
	prices := make([]float64, len(strslice))
	for spricesindex, values := range strslice {
		convfloatprice, err := strconv.ParseFloat(values, 64)
		if err != nil {
			panic(err)
		}
		prices[spricesindex] = convfloatprice
	}
	return prices, nil
}
