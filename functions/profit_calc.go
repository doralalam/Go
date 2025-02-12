/*

package main
import "fmt"
func main(){
	// var revenue, expenses, taxRate float64
	var revenue float64 = rev()
	var expenses float64= exp()
	var taxRate float64= tax()
	ebt, profit, ratio := calculations(revenue, expenses, taxRate)
	fmt.Printf("EBT is %.1f\nProfit is %.1f\nRatio is %.1f\n",ebt, profit, ratio)

}

func rev() float64 {
	var revenue float64
	fmt.Println("Enter revenue: ")
	fmt.Scan(&revenue)
	return revenue
}


func exp() float64 {
	var expenses float64
	fmt.Println("Enter expenses: ")
	fmt.Scan(&expenses)
	return expenses
}


func tax() float64 {
	var taxRate float64
	fmt.Println("Enter taxRate: ")
	fmt.Scan(&taxRate)
	return taxRate
}

func calculations(revenue, expenses, taxRate float64)(float64, float64, float64){
	ebt := revenue-expenses
	profit := ebt * (1-taxRate/100)
	ratio := ebt/profit
	return ebt, profit, ratio
}

*/

//This code can also be written with minimal functions like this

package main

import "fmt"

func main() {
	// no need to declare variables whenever we are using colon
	revenue := getInput("Revenue: ")
	expenses := getInput("Expenses: ")
	taxRate := getInput("Tax Rate: ")
	ebt, profit, ratio := calculations(revenue, expenses, taxRate)
	fmt.Printf("EBT is %.1f\nProfit is %.1f\nRatio is %.1f\n", ebt, profit, ratio)

}

func getInput(text string) float64 {
	var inputVar float64
	fmt.Print(text)
	fmt.Scan(&inputVar)
	return inputVar
}

func calculations(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
