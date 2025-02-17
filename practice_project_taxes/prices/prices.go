package prices

import (
	"fmt"

	"example.com/tax_project/conversion"
	"example.com/tax_project/filemanager"
)

// struct type
type TaxStructure struct {
	IOManager        filemanager.Filemanager `json:"-"` //to neglect the label in json
	TaxRate          float64                 `json:"tax_rate"`
	InputPrices      []float64               `json:"input_prices"`
	TaxIncludedPrice map[string]string       `json:"tax_included_prices"`
}

// constructor
func NewTaxObject(fm filemanager.Filemanager, tax float64) *TaxStructure {
	return &TaxStructure{IOManager: fm,
		InputPrices: []float64{},
		TaxRate:     tax,
	}
}

// method for calculation
func (job *TaxStructure) Process() {
	// create a dynamic map using make()
	job.LoadData()
	result := make(map[string]string)
	for _, prices := range job.InputPrices {
		taxedPrice := prices * (1 + (job.TaxRate / 100))
		result[fmt.Sprintf("%.1f", prices)] = fmt.Sprintf("%.1f", taxedPrice)
	}
	// call a function from another packagae to write the data in JSON format
	job.TaxIncludedPrice = result
	job.IOManager.WriteJSON(job)
	//filemanager.WriteJSON(fmt.Sprintf("result_%v.json", job.TaxRate), job)

}

// method for loading the data
func (job *TaxStructure) LoadData() {
	// call a function from another package to open the file to read
	stringValues, err := job.IOManager.ReadingFromFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	// call a function to convert from string to float64
	floatValues, err := conversion.StringsToFloat(stringValues)
	if err != nil {
		fmt.Println(err)
		return
	}
	job.InputPrices = floatValues
}
