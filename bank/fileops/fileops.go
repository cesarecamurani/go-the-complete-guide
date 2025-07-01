package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const ReadWritePermissionCode = 0644

func FetchFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	defaultValue := 1000.0

	if err != nil {
		return defaultValue, errors.New("failed to find file")
	}

	text := string(data)
	result, err := strconv.ParseFloat(text, 64)

	if err != nil {
		return defaultValue, errors.New("failed to parse result")
	}

	return result, nil
}

func SaveFloatToFile(float float64, fileName string) {
	err := os.WriteFile(fileName, []byte(fmt.Sprintf("%.2f", float)), ReadWritePermissionCode)
	if err != nil {
		return
	}
}
