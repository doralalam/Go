package main

import "fmt"

func main() {
	prices := []float64{12.43, 32.23, 34.23}
	fmt.Println("Prices: ", prices)

	discountedPrices := []float64{23.32, 32.44, 54.53, 46.26}
	// prices = append(prices, discountedPrices)
	// we can't append like this because prices is a list and discountedPrices is also a list
	// we can't append a list to another list
	// to achieve this, we must unpack the list using ...
	prices = append(prices, discountedPrices...)
	// discountedPrices... indicates the go to unpack the list and then append to the destination
	fmt.Println("Updated Prices:", prices)

}
