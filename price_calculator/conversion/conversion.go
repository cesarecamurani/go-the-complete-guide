package conversion

import (
	"fmt"
	"strconv"
)

const bitSize = 64

func StringsToFloat(strings []string) ([]float64, error) {
	var floats []float64

	for _, line := range strings {
		floatPrice, parseErr := strconv.ParseFloat(line, bitSize)

		if parseErr != nil {
			return nil, fmt.Errorf("failed to parse string to float: %v", parseErr)
		}

		floats = append(floats, floatPrice)
	}

	return floats, nil
}
