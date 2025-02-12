//example for Data Mutation

package main

import "fmt"

func main() {
	age := 32
	agePointer := &age
	fmt.Println("Age: ", age)
	getAdultYears(agePointer) //function is passed with pointer as a paramater
	fmt.Println("Adult Age: ", age)
}

func getAdultYears(age *int) { // function accepting the pointer as a parameter and have no return type
	*age = *age - 18 // *age is deferenced -> age - 18 -> *age again referenced -> *age = age-18 value is stored
	// we haven't used any return here
	// Directly, the result was written into the origin variable only
}

/*

fmt.Scan(&age) is also one of the example for data mutation
Here, we are passing (address of the age variable) pointer to scan function
Then whatever user enters, was overwritten on the the address specified or pointer

*/
