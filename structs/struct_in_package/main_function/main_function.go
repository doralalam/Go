package main

import (
	"fmt"

	"example.com/struct_package/user" // imported package
)

func main() { // FUNCTION

	userfirstName := getUserDetails("Enter the first name: ")
	userlastName := getUserDetails("Enter the last name: ")
	userbirthDate := getUserDetails("Enter the birth date in DD/MM/YYYY format: ")

	var userInstance *user.User                                               //creating an instance to the struct (user) using pointer to avoid copying the data
	userInstance, err := user.New(userfirstName, userlastName, userbirthDate) // calling constructor in another package
	if err != nil {
		fmt.Println(err)
		return
	}
	adminInstance := user.NewAdmin("lalam.com", "1234")
	// calling the method using the instance
	userInstance.PrintUserDetails()
	userInstance.ClearUserDetails()
	userInstance.PrintUserDetails()

	adminInstance.UserInfo.PrintUserDetails()
	adminInstance.UserInfo.ClearUserDetails()
	adminInstance.UserInfo.PrintUserDetails()
}

func getUserDetails(text string) string { // FUNCTION to take input values for the fields
	fmt.Print(text)
	var temp string
	fmt.Scanln(&temp) // Scanln ends the scan operation when user hits enter key
	return temp
}
