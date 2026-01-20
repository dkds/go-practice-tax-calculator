package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for i, str := range strings {
		floatVal, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, errors.New("Failed to convert the string to float: " + str)
		}

		floats[i] = floatVal
	}
	return floats, nil
}
