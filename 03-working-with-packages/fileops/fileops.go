package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return defaultValue, errors.New("Failed to read the file.")
	}
	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("Failed to convert value.")
	}
	return value, nil
}

func WriteOutputToFile(value float64, fileName string) {
	outputText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(outputText), 0644)
}
