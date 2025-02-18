package cmdmanager

import "fmt"

type CMDmanager struct {
}

func (cmd CMDmanager) ReadingFromFile() ([]string, error) {
	prices := []string{}
	fmt.Println("Enter the prices. Enter 0 to exit")
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

func (cmd CMDmanager) WriteJSON(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDmanager {
	return CMDmanager{}
}
