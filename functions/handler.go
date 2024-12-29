package functions

import (
	"errors"
	"fmt"
)

func validateNumber(number float64) (bool, error) {
	if number <= 0 {
		return false, fmt.Errorf("your number must be greater than 0, %f is unexpected", number)
	}
	return true, nil
}

func customErrorText(text string) error {
	return errors.New(text)
}