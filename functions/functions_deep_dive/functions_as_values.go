// to demonstrate functions as values

package main

import "fmt"

// custom type for function for better readbility
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	// pass double function as value to transformNumbers function
	doubledNumbers := transformNumbers(&numbers, double)
	fmt.Println("Doubled Number slice: ", doubledNumbers)
	// pass triple function as value to transformNumbers function
	tripledNumbers := transformNumbers(&numbers, triple)
	fmt.Println("Tripled Number slice: ", tripledNumbers)

}

// transformNumbers function with numbers of slice type and multiple of custom transformFn type as values / parameters with return type as int slice
func transformNumbers(numbers *[]int, multiple transformFn) []int {
	// temporary int slice
	tNumbers := []int{}
	for _, val := range *numbers {
		// appending
		tNumbers = append(tNumbers, multiple(val))
	}
	return tNumbers

}

// double function
func double(value int) int {
	return value * 2
}

// triple function
func triple(value int) int {
	return value * 3
}
