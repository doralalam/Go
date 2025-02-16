package main

import "fmt"

func main() {
	numbers := []int{1, 2, 4, 5}
	// call createTransform function to determine the factor
	// this function returns an anonymous function to multiply factor with the numbers
	double := createTransform(2)
	doubledNumbers := transformNumbers(numbers, double)
	fmt.Println("Doubled Numbers: ", doubledNumbers)
	// call createTransform function to determine the factor
	// this function returns an anonymous function to multiply factor with the numbers
	triple := createTransform(3)
	tripledNumbers := transformNumbers(numbers, triple)
	fmt.Println("Tripled Numbers: ", tripledNumbers)

}

// transformNumbers function accepts "int slice" and "a function that takes int as parameter and returns int" as parameters
// and returns int slice
func transformNumbers(numbers []int, transform func(num int) int) []int {
	demoNumbers := []int{}
	for _, val := range numbers {
		demoNumbers = append(demoNumbers, transform(val))
	}
	return demoNumbers
}

// createTranform is a function accepts an integer as parameter
// and returns an anonymous function
func createTransform(fact int) func(num int) int {
	return func(num int) int {
		// anonymous function declared inside a function can access the parameter of outerfunction due to its scope
		return num * fact
	}
}
