//example for avoiding unecessary value copies

package main

import "fmt"

func main() {
	age := 32
	var agePointer *int // declaring a pointer with *int
	agePointer = &age   // assigning a pointer with the address of the variable
	// agePointer := &age
	fmt.Println("Age: ", age)
	fmt.Println("Adult Age: ", getAdultYears(agePointer)) // passing a pointer to a function as a parameter
	// fmt.Println("Adult Age: ", getAdultYears(&age))
}

func getAdultYears(age *int) int { // function accepting the pointer as a parameter
	return *age - 18 // dereferencing a pointer to perform arithmetic operation
}
