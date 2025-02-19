// Main package having the main function.
// execution starts from here
// create a text file with float value before running

package main

import (
	"fmt"

	"example.com/bankPackage/fileOperations"
)

const file_name string = "balance.txt" // create a file and assign file_name here

func main() {
	var accountBalance, err = fileOperations.ReadFromFile(file_name)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("-----------------------")
		panic("Can't continue")
	}
	fmt.Println("Welcome to GO Bank!")
	for {
		BankingOptions() // this function is declared in another file within the main package
		// we can define the function with in the same package with lowerCase as well as upperCase
		// whereas if the function is from different package, starting letter must be UpperCase
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
				fileOperations.WriteToFile(accountBalance, file_name)
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
				fileOperations.WriteToFile(accountBalance, file_name)
			} else {
				fmt.Println("Invalid Transaction")
			}

		default:
			fmt.Println("Thanks for choosing GO bank!")
			return
		}
	}
}
