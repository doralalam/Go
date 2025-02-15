// This program is a good example to demonstrate the len() and cap() functions

package main

import "fmt"

/*

func main() {
	var productPrices [4]float64 = [4]float64{10.5, 12.6, 23.5, 54.3}
	fmt.Println("Product Prices:", productPrices)
	featuredPrices := productPrices[:1] // all the elements in the right side from 1 were ommitted
	fmt.Println("Featured Prices:", featuredPrices)
	highlightedPrices := featuredPrices[0:] // all the elements from start to end were considered
	fmt.Println("Highlighted Prices:", highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
	// lenght will be 1 since highlightedPrices have only one element
	// capacity will be 4 since highlightedPrices was formed from slicing featuredPrices which again formed from
	// slicing productPrices. And in this entire slicing phase, no values were ommitted on the starting side (left side)
	// go internally remembers all the values sliced on ending side (left side) but not on starting side (left side)
}

*/

func main() {
	var productPrices [4]float64 = [4]float64{10.5, 12.6, 23.5, 54.3}
	fmt.Println("Product Prices:", productPrices)
	fmt.Println("2nd element in Product Prices:", productPrices[1]) // accessing the 2nd element
	productPrices[3] = 100.0                                        // resetting the 4th element
	fmt.Println("After resetting 4th element in Product Prices: ", productPrices)
	featuredPrices := productPrices[1:] // 0th element was ommitted in left
	fmt.Println("Featured Prices:", featuredPrices)
	highlightedPrices := featuredPrices[0:] // all the elements from start to end were considered
	fmt.Println("Highlighted Prices:", highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
	// lenght will be 3 since highlightedPrices have all elements from featuredPrices
	// capacity will be 3 since highlightedPrices was formed from slicing featuredPrices which again formed from
	// slicing productPrices by ommitting 0th element on left side.
	// go internally remembers all the values sliced on ending side (left side) but not on starting side (left side)
}
