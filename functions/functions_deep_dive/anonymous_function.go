package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	// pass an anonymous function as value in transformNumbers function
	// it is better to use anonymous function whenever a function is called for only once to save program memory
	// we call a function as an anonymous if it doesn't have a name to call it specifically
	transformedNumbers := transformNumbers(numbers, func(number int) int {
		return number * 2
	})
	fmt.Println("Transformed Number slice: ", transformedNumbers)

}

func transformNumbers(numbers []int, transform func(int) int) []int {
	demoNumbers := []int{}
	for _, value := range numbers {
		// append the result of anonymous function
		demoNumbers = append(demoNumbers, transform(value))
	}
	return demoNumbers
}
