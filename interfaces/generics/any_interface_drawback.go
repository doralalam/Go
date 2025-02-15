// This code is too lengthy even for a small operation like addition

package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)
	result = add(10.5, 12.8)
	fmt.Println(result)
	result = add("hello", "world")
	fmt.Println(result)
}

func add(a, b interface{}) interface{} {
	// this function accepts any input value parameter and returns any output
	// but the lenght of the code is too much as we must handle with each and every test case manually
	aInt, aBool := a.(int) // returns true to aBool if a is int
	bInt, bBool := b.(int) // returns true t bBool if b is int
	if aBool && bBool {    // if both a & b are true, then perfroms addition
		return aInt + bInt
	}
	aFloat, aBool := a.(float64)
	bFloat, bBool := b.(float64)
	if aBool && bBool {
		return aFloat + bFloat
	}
	aString, aBool := a.(string)
	bString, bBool := b.(string)
	if aBool && bBool {
		return aString + bString
	}
	return nil
}
