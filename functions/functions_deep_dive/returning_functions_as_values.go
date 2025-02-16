// to demonstrate returning functions as values

package main

import "fmt"

// custom type for function for better readbility
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	anotherNumbers := []int{3, 4, 5, 6}
	// pass double function as value to transformNumbers function
	doubledNumbers := transformNumbers(&numbers, double)
	fmt.Println("Doubled Number slice: ", doubledNumbers)
	// pass triple function as value to transformNumbers function
	tripledNumbers := transformNumbers(&numbers, triple)
	fmt.Println("Tripled Number slice: ", tripledNumbers)
	// transformNumbers function receiving double function as value from getFunction
	// i.e., getFunction returns double function as value
	newDoubledNumbers := transformNumbers(&numbers, getFunction(numbers))
	fmt.Println("Newly Doubled Numbers: ", newDoubledNumbers)
	// transformNumbers function receiving triple function as value from getFunction
	// i.e., getFunction returnstriple function as value
	newTripledNumbers := transformNumbers(&anotherNumbers, getFunction(anotherNumbers))
	fmt.Println("Newly Tripled another numbers: ", newTripledNumbers)

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

// getFunction returns new function of custom type transformFn
func getFunction(numbers []int) transformFn {
	if (numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

// double function
func double(value int) int {
	return value * 2
}

// triple function
func triple(value int) int {
	return value * 3
}
