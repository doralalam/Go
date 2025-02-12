package main

import "fmt"

func main() {
	var accountBalance = 10000.0
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
		//if, else if and else block codes must be enclosed in curly braces no matter how many statements are there
		if choice == 1 {
			fmt.Printf("Your account balance is %v", accountBalance)
		} else if choice == 2 {
			fmt.Print("Enter the amount you want to deposit: ")
			var dep float64
			fmt.Scan(&dep)
			if dep > 0 {
				accountBalance += dep
				fmt.Printf("Your updated account balance is %v\n", accountBalance)
			} else {
				fmt.Println("Invalid Transaction")
			}
		} else if choice == 3 {
			fmt.Print("Enter the amount you want to withdraw: ")
			var withdraw float64
			fmt.Scan(&withdraw)
			if withdraw > 0 && withdraw <= accountBalance {
				accountBalance -= withdraw
				fmt.Printf("Your updated account balance is %v\n", accountBalance)
			} else {
				fmt.Println("Invalid Transaction")
			}
		} else {
			break
			//we can use break statement to exit the loop
			//we can use continue statement to skip one iteration
		}
	}
	fmt.Println("Thankyou for choosing GO Bank!")
}
