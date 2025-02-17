package main

import (
	"fmt"

	"example.com/tax_project/filemanager"
	"example.com/tax_project/prices"
)

func main() {
	taxRates := []float64{5, 10, 18}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%v.json", taxRate))
		newObject := prices.NewTaxObject(fm, taxRate)
		newObject.Process()
	}
}
