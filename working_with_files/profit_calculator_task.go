package main

import (
	"fmt"
	"os"
)

func main() {
	// no need to declare variables whenever we are using colon
	revenue := getInput("Revenue")
	expenses := getInput("Expenses")
	taxRate := getInput("Tax Rate")
	ebt, profit, ratio := calculations(revenue, expenses, taxRate)
	textToFile := fmt.Sprintf("EBT is %.1f\nProfit is %.1f\nRatio is %.1f\n", ebt, profit, ratio)
	writeFile(textToFile)
}

func writeFile(textToFile string) { //write to file
	os.WriteFile("profit_cal_results.txt", []byte(textToFile), 0644)
}

func getInput(text string) float64 {
	var inputVar float64
	fmt.Print(text, ": ")
	fmt.Scan(&inputVar)
	if inputVar > 0 {
		return inputVar
	} else {
		dummy := fmt.Sprintf("Invalid input for %s.\n%s value should be greater than 0", text, text)
		panic(dummy)
	}
}

func calculations(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
