// This code with minimal changes can be applicable to arrays, slices and maps

package main

import "fmt"

func main() {
	websites := map[string]string{}
	var count int
	fmt.Print("Enter number of websites you would like to add: ")
	fmt.Scan(&count)
	var key, val string
	// to read input from the user
	for i := 0; i < count; i++ {
		fmt.Print("Enter your website name: ")
		fmt.Scan(&key)
		fmt.Print("Enter your website address: ")
		fmt.Scan(&val)
		websites[key] = val
	}
	// to display the output in console
	for nKey, nval := range websites {
		fmt.Printf("%v : %v\n", nKey, nval)
	}
}
