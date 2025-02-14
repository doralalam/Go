package main

import (
	"fmt"
	"time"
)

// type struct_name struct datatype
type user struct { // CUSTOM DATA TYPE STRUCT
	firstName   string // fields inside the struct
	lastName    string
	birthDate   string
	currentTime time.Time // from time package
}

// This is a method attached to user struct
// unlike regular function, to declare a method, we need not to pass the parameters
// but we should attach the method to a custom data type. here it is user
// (userInstance user) is called as receiver argument
// name of this method is printUserDetails()

// it doesn't matter if we pass pointers or values to this method as it is just to display instead of modifying
func (userInstance user) printUserDetails() { // METHOD
	fmt.Println(userInstance.firstName)
	fmt.Println(userInstance.lastName)
	fmt.Println(userInstance.birthDate)
	fmt.Println(userInstance.currentTime)
}

/*
// to modify the fields inside struct
// this method will create a copy of the original fields and modifies them
// resulting the original fields won't change

// passing values

func (u user) clearUserDetails() {
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
}
*/

// passing pointers to modify the original fields instead of copying

func (u *user) clearUserDetails() { // METHOD
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
}

func main() { // FUNCTION
	userfirstName := getUserDetails("Enter the first name: ")
	//var userlastName string
	userlastName := getUserDetails("Enter the last name: ")
	userbirthDate := getUserDetails("Enter the birth date in DD/MM/YYYY format: ")
	var userInstance user //creating an instance to the struct (user)
	userInstance = user{
		firstName:   userfirstName,
		lastName:    userlastName,
		birthDate:   userbirthDate,
		currentTime: time.Now(),
	}
	// calling the method using the instance
	userInstance.printUserDetails()
	userInstance.clearUserDetails()
	userInstance.printUserDetails()
}

func getUserDetails(text string) string { // FUNCTION
	fmt.Print(text)
	var temp string
	fmt.Scan(&temp)
	return temp
}
