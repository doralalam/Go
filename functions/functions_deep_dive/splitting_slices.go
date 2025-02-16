// demonstration of variadic function

package main

import "fmt"

func main() {
	// variadic function is also useful whenever we don't have a slice but multiple values to pass to a funciton
	sum := sumup(1, 2, 3, 4, 5)
	numbers := []int{1, 10, 15}
	// splitting the slice into parameter values using ellipses (...)
	anotherSum := sumup(numbers...)
	fmt.Println("Sum of the values (except 1st value):", sum)
	fmt.Println("Another sum value:", anotherSum)

}

// from the parameters, 1st value is taken as start and rest of the paramaters we internally made into slice
func sumup(numbers ...int) int {
	// ... is called as ellipsis and used to accept 'n' number of parameters from a user
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
