package cmdmanager

import (
	"encoding/json"
	"fmt"
)

type CMDManager struct {
}

func (cmd CMDManager) ReadFile() ([]string, error) {
	fmt.Print("Please enter the prices.Confirm each price with enter")
	var index int
	var pricelist []string
	for {
		var price string
		fmt.Printf("Price%v : ", index)
		fmt.Scan(&price)
		if price == "0" {
			break
		}
		pricelist = append(pricelist, price)
	}
	return pricelist, nil
}

func (cmd CMDManager) WriteJSONFile(data interface{}) error {
	fmt.Println(json.Marshal(data))
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
