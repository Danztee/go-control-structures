package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading balance file:", err)
		return 0, errors.New("failed to find balance file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 0, errors.New("failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueTxt := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueTxt), 0o644)
}
