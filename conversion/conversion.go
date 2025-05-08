package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(inputStrings []string) ([]float64, error) {
	var floatDatas []float64
	for _, stringVal := range inputStrings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {

			return nil, errors.New("failed to convert string to float")
		}

		floatDatas = append(floatDatas, floatPrice)
	}

	return floatDatas, nil
}
