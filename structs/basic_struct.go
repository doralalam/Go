package main

import (
	"fmt"
	"time"
)

// custom data type
// type struct_name struct datatype
type user struct {
	firstName   string // fields inside the struct
	lastName    string
	birthDate   string
	currentTime time.Time // from time package
}

func main() {
	userfirstName := getUserDetails("Enter the first name: ")
	//var userlastName string
	userlastName := getUserDetails("Enter the last name: ")
	userbirthDate := getUserDetails("Enter the birth date in DD/MM/YYYY format: ")
	var userInstance user //creating an instance to the struct (user)
	userInstance = user{firstName: userfirstName, lastName: userlastName, birthDate: userbirthDate, currentTime: time.Now()}
	// in the above method of instantiation, order of the fields doesn't matter as they are in form of key value paris
	// if any value is not assigned, then by default if will have null values
	// struct can be assigned like this also
	// userInstnace = user{userfirstName, userlastName, userbirthDate, time.Now()}
	// in the above method, order of the parameters should be same as in the struct declaration
	printUserDetails(userInstance)
}

func printUserDetails(userInstance user) { // passing the parameter userInstance of user type (struct)
	// advantage of the struct is this. Unlike, regular functions, where all individual parameters need to be passed
	// in struct, we can user only struct parameter (custom data type)
	// reduces the chances of mistake
	fmt.Println(userInstance.firstName)
	fmt.Println(userInstance.lastName)
	fmt.Println(userInstance.birthDate)
	fmt.Println(userInstance.currentTime)
}

func getUserDetails(text string) string {
	fmt.Print(text)
	var temp string
	fmt.Scan(&temp)
	return temp
}
