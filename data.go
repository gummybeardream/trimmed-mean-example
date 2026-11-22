package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readValueColumn(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("open %s: %w", filename, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", filename, err)
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("%s must contain a header and data", filename)
	}

	if len(records[0]) != 1 || records[0][0] != "value" {
		return nil, fmt.Errorf("%s must have one column named value", filename)
	}

	values := make([]string, 0, len(records)-1)

	for index, record := range records[1:] {
		rowNumber := index + 2

		if len(record) != 1 {
			return nil, fmt.Errorf(
				"%s row %d must contain one value",
				filename,
				rowNumber,
			)
		}

		value := strings.TrimSpace(record[0])
		if value == "" {
			return nil, fmt.Errorf(
				"%s row %d contains an empty value",
				filename,
				rowNumber,
			)
		}

		values = append(values, value)
	}

	return values, nil
}

func readIntegers(filename string) ([]int, error) {
	textValues, err := readValueColumn(filename)
	if err != nil {
		return nil, err
	}

	values := make([]int, 0, len(textValues))

	for index, text := range textValues {
		value, err := strconv.Atoi(text)
		if err != nil {
			return nil, fmt.Errorf(
				"%s row %d is not a valid integer: %w",
				filename,
				index+2,
				err,
			)
		}

		values = append(values, value)
	}

	return values, nil
}

func readFloats(filename string) ([]float64, error) {
	textValues, err := readValueColumn(filename)
	if err != nil {
		return nil, err
	}

	values := make([]float64, 0, len(textValues))

	for index, text := range textValues {
		value, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return nil, fmt.Errorf(
				"%s row %d is not a valid floating-point number: %w",
				filename,
				index+2,
				err,
			)
		}

		values = append(values, value)
	}

	return values, nil
}
