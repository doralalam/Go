package main

import (
	"errors"
	"fmt"
	"time"
)

func newUserConstructor(userfirstName, userlastName, userbirthDate string) (*user, error) { // CONSTRUCTOR
	// CONSTRUCTOR is a special function used to initialize the objects in the start of the program
	// contructor name should start with new as a convention but not mandatory
	if userfirstName == "" || userlastName == "" || userbirthDate == "" {
		// for data validation
		return nil, errors.New("FirstName, LastName and BirthDate are mandatory fields.")
	}
	return &user{
		firstName:   userfirstName,
		lastName:    userlastName,
		birthDate:   userbirthDate,
		currentTime: time.Now(),
	}, nil
}

type user struct { // CUSTOM DATA TYPE STRUCT
	firstName   string
	lastName    string
	birthDate   string
	currentTime time.Time
}

func (userInstance user) printUserDetails() { // METHOD to print values in the fields
	fmt.Println(userInstance.firstName)
	fmt.Println(userInstance.lastName)
	fmt.Println(userInstance.birthDate)
	fmt.Println(userInstance.currentTime)
}

func (u *user) clearUserDetails() { // METHOD to clear the fields
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
}

func main() { // FUNCTION
	userfirstName := getUserDetails("Enter the first name: ")
	userlastName := getUserDetails("Enter the last name: ")
	userbirthDate := getUserDetails("Enter the birth date in DD/MM/YYYY format: ")
	var userInstance *user                                                              //creating an instance to the struct (user) using pointer to avoid copying the data
	userInstance, err := newUserConstructor(userfirstName, userlastName, userbirthDate) // calling constructor
	if err != nil {
		fmt.Println(err)
		return
	}
	// calling the method using the instance
	userInstance.printUserDetails()
	userInstance.clearUserDetails()
	userInstance.printUserDetails()
}

func getUserDetails(text string) string { // FUNCTION to take input values for the fields
	fmt.Print(text)
	var temp string
	fmt.Scanln(&temp) // Scanln ends the scan operation when user hits enter key
	return temp
}
