package main

import (
	"fmt"
	"os"
	"strconv" // package for string conversion
)

const file_name string = "balance.txt" // assign file_name here

func readFromFile() float64 {
	data, _ := os.ReadFile(file_name) //data variable contains the data in byte collection format
	// _ is a special kind of variable that describes the go that you are getting a value
	// but not interested in using it now. (os.ReadFile() returns data and error)
	balanceText := string(data)                       // converting byte collection to string
	balance, _ := strconv.ParseFloat(balanceText, 64) // converting string to float64
	return balance
}

func writeToFile(balance float64) {
	balanceText := fmt.Sprint(balance)                 // convert balance from float64 to string
	os.WriteFile(file_name, []byte(balanceText), 0644) //create the file named balance.txt if the file does not exist
	// []byte() used to convert the string to byte collection
	// 0644 is the file permission
}

func main() {
	accountBalance := readFromFile()
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
