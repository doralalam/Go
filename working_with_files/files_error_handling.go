package main

import (
	"errors" // new package to handle errors
	"fmt"
	"os"
	"strconv" // package for string conversion
)

const file_name string = "balance.txt" // create a file and assign file_name here

func readFromFile() (float64, error) { // error is a new data type available errors package
	data, err := os.ReadFile(file_name) //returns data and error
	if err != nil {                     // nil is a special keyword in go that indicates null
		return 1000, errors.New("Failed to load balance file\nCreate a balance file with name 'balance.txt'")
		//errors.New() returns the error type data string
	}
	balanceText := string(data)                         //byte collection to string
	balance, err := strconv.ParseFloat(balanceText, 64) //string to float and error
	if err != nil {
		return 1000, errors.New("Failed to parse the balance file\nThe file has been corrupted with inappropriate data")
	}
	return balance, nil
}

func writeToFile(balance float64) {
	balanceText := fmt.Sprint(balance)                 // convert balance from float64 to string
	os.WriteFile(file_name, []byte(balanceText), 0644) //create the file named balance.txt if the file does not exist
	// []byte() used to convert the string to byte collection
	// 0644 is the file permission
}

func main() {
	var accountBalance, err = readFromFile() //to capture float64 and errors from readFromFile()
	if err != nil {                          // if error exists, print the error and terminate the program
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("-----------------------")
		panic("Can't continue") // panic() is a function that terminates the program
	}
	fmt.Println("Welcome to GO Bank!")
	// for is the only looping structure in Go
	for {
		fmt.Println("Select what do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice here: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your account balance is %v\n", accountBalance)

		case 2:
			fmt.Print("Enter the amount you want to deposit: ")
			var dep float64
			fmt.Scan(&dep)
			if dep > 0 {
				accountBalance += dep
				fmt.Printf("Your updated account balance is %v\n", accountBalance)
				writeToFile(accountBalance)
			} else {
				fmt.Println("Invalid Transaction")
			}

		case 3:
			fmt.Print("Enter the amount you want to withdraw: ")
			var withdraw float64
			fmt.Scan(&withdraw)
			if withdraw > 0 && withdraw <= accountBalance {
				accountBalance -= withdraw
				fmt.Printf("Your updated account balance is %v\n", accountBalance)
				writeToFile(accountBalance)
			} else {
				fmt.Println("Invalid Transaction")
			}

		default:
			fmt.Println("Thanks for choosing GO bank!")
			return
			//By default, every case in switch have break at the end. So, to come out of the loop, the only possbile option is return
		}
	}
}
