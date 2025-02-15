// This code resolves the drawback of complexity mentioned in any_interface_drawback.go file

package main

import "fmt"

func main() {
	result1 := add(1, 2) // each time the function is called, result is stored in new variable
	fmt.Println(result1)
	result2 := add(10.5, 12.8)
	fmt.Println(result2)
	result3 := add("hello", "world")
	fmt.Println(result3)
}

func add[T int | float64 | string](a, b T) T {
	// This function uses generics instead of any
	// [T] is the place holder and inside the square brackets, we can mention the expected data types
	// this T indicates the possibile data types we mentioned. So we can use T instead of repeating those types again & again
	// i.e., (a,b T) indicates parameter could be of any type in T and return type could also be T (example: int or float64 or string)
	return a + b
}
