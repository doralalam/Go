package main

import "fmt"

func main() {
	// call a function
	result := factorial(6)
	fmt.Println("Factorial is", result)
}

func factorial(number int) int {
	// recusion continues until one of the instance of the function returns a value
	if number == 1 {
		return 1
	}
	// call function itself
	return number * factorial(number-1)
}
