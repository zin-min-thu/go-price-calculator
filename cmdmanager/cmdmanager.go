package cmdmanager

import (
	"fmt"
)

type CMDManager struct {
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("please enter your prices, confirm every price with enter")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManager) WriteResult(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}
