package user

import (
	"errors"
	"fmt"
	"time"
)

// we can create custom data types using type
// example type str string
// gives an alias name str to string which is an inbuilt data type

type User struct { // CUSTOM DATA TYPE STRUCT
	firstName   string
	lastName    string
	birthDate   string
	currentTime time.Time
}

// struct embedding
type Admin struct {
	email    string
	password string
	UserInfo User // userInfo is the name of the field and User is the custom datatype
	// if we don't assign a field name, then this will work as anonymous (just mention the custom datatype User)
}

func NewAdmin(userEmail, userPassword string) Admin {
	return Admin{
		email:    userEmail,
		password: userPassword,
		UserInfo: User{
			firstName:   "ADMIN",
			lastName:    "ADMIN",
			birthDate:   "----",
			currentTime: time.Now(),
		},
	}
}

func New(userfirstName, userlastName, userbirthDate string) (*User, error) { // CONSTRUCTOR
	if userfirstName == "" || userlastName == "" || userbirthDate == "" {
		// for data validation
		return nil, errors.New("FirstName, LastName and BirthDate are mandatory fields")
	}
	return &User{
		firstName:   userfirstName,
		lastName:    userlastName,
		birthDate:   userbirthDate,
		currentTime: time.Now(),
	}, nil
}

func (userInstance User) PrintUserDetails() { // METHOD to print values in the fields
	fmt.Println(userInstance.firstName)
	fmt.Println(userInstance.lastName)
	fmt.Println(userInstance.birthDate)
	fmt.Println(userInstance.currentTime)
}

func (u *User) ClearUserDetails() { // METHOD to clear the fields
	u.firstName = ""
	u.lastName = ""
	u.birthDate = ""
}
