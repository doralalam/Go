package main

import "fmt"

func main() {
	prices := []float64{10.89, 7.43}
	fmt.Println("Prices:", prices)
	fmt.Println("Lenght of Prices array:", len(prices))
	fmt.Println("Capacity of Prices array:", cap(prices))
	prices = append(prices, 21.9)
	fmt.Println("Updated Prices after adding new element:", prices)
	fmt.Println("Lenght of Prices array:", len(prices))
	fmt.Println("Capacity of Prices array:", cap(prices))
	// Go often doubles the capacity of a slice when it needs to grow.
	// So, from an initial capacity of 2, Go increased the capacity to 4
	// to accommodate the new element and have room for future growth.
	prices = prices[1:]
	fmt.Println("Updated Prices after deleting the 1st element:", prices)
	fmt.Println("Lenght of Prices array:", len(prices))
	fmt.Println("Capacity of Prices array:", cap(prices))

}
