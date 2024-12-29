package functions

import (
	"errors"
	"fmt"
)

func validateNumber(number *float64) (bool, error) {
	if *number <= 0 {
		return false, fmt.Errorf("%v is unexpected, your number must be greater than 0", *number)
	}
	return true, nil
}

func limitNumber(number *float64, threshold *float64) (bool, error) {
	if *number <= 0 || *number > *threshold{
		return false, fmt.Errorf("%v is unexpected, your number must be greater than 0 and less than or equal %v", *number, *threshold)
	}
	return true, nil
}

func customErrorText(text string) error {
	return errors.New(text)
}