/*
	This code is to demonstrate the advantage of make function.
	In general, we declare a dynamic slice like this
		userNames := []string{}
	Whenever the above line gets executed, an empty array with pre-allocated memory will be created and when we append new value to the array,
	each time, the old slice will get dictched and then new slice with that length will be created to form an array.
	This process continues as many times as we append new values.
	However, using the make(), we can initially create a slice of expected length and capacity, so that until capacity
	got filled, the creation of new slice ditching the old slice will not take place leaving compiler less work.
	Creation of new slice only takes place only when the capacity of the created slice got filled.

	Syntax to use make() in arrays
		make([]<slice_type>, <length>, <capacity>)
	Example:
		userNames := make([]string, 2, 5)
		->	this line pre-allocates the memory required for slice with 2 indexed and a max of 5 indexes


	This same function can also be used with maps
	Syntax to use make() with maps
		make(map[string]string, <length>)
	Example:
		courseRatings := make(map[string]float64, 2)
		->	this line pre-allocates the memory for 2 map structures
*/

package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)
	// as we mentioned length as 2, we can assign first 2 values as below
	userNames[0] = "Ram"
	// since, we haven't assigned 2nd value, it will be left as empty
	fmt.Println("Usernames:", userNames)
	userNames = append(userNames, "Laxman")
	userNames = append(userNames, "Hanuman")
	userNames = append(userNames, "Sugreev")
	// now the slice gets full and new slice will be created dynamically due to the usage of make()
	userNames = append(userNames, "Angadh")
	fmt.Println("Usernames:", userNames)

}
