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

// This is a method attached to user struct
// unlike regular function, to declare a method, we need not to pass the parameters
// but we should attach the method to a custom data type. here it is user
// (userInstance user) is called as receiver argument
// name of this method is printUserDetails()

func (userInstance user) printUserDetails() {
	fmt.Println(userInstance.firstName)
	fmt.Println(userInstance.lastName)
	fmt.Println(userInstance.birthDate)
	fmt.Println(userInstance.currentTime)
}

func main() {
	userfirstName := getUserDetails("Enter the first name: ")
	//var userlastName string
	userlastName := getUserDetails("Enter the last name: ")
	userbirthDate := getUserDetails("Enter the birth date in DD/MM/YYYY format: ")
	var userInstance user //creating an instance to the struct (user)
	userInstance = user{firstName: userfirstName, lastName: userlastName, birthDate: userbirthDate, currentTime: time.Now()}
	// calling the method using the instance
	userInstance.printUserDetails()
}

func getUserDetails(text string) string {
	fmt.Print(text)
	var temp string
	fmt.Scan(&temp)
	return temp
}
