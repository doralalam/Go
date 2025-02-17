package conversion

import (
	"errors"
	"strconv"
)

// string to float conversion
func StringsToFloat(strings []string) ([]float64, error) {
	floatValues := make([]float64, len(strings))
	for lineIndex, lineValue := range strings {
		// to convert data from string to float64
		prasedValue, err := strconv.ParseFloat(lineValue, 64)
		if err != nil {
			return nil, errors.New("Conversion Failed")
		}
		floatValues[lineIndex] = prasedValue
	}
	return floatValues, nil
}
