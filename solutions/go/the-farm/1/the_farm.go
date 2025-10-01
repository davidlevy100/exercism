// Package thefarm computes per-cow fodder using a provided calculator.
package thefarm

import (
	"errors"
	"fmt"
)

// InvalidCowsError is returned for negative or zero cows.
type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// DivideFood calculates per-cow fodder amount. Validates cows > 0.
func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	if err := ValidateNumberOfCows(cows); err != nil {
		return 0, err
	}

	fodder, err := fc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return fodder * factor / float64(cows), nil
}

// ValidateInputAndDivideFood checks input, returns a generic error if invalid.
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if err := ValidateNumberOfCows(cows); err != nil {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, cows)
}

// ValidateNumberOfCows returns an error for negative or zero cows.
func ValidateNumberOfCows(cows int) error {
	switch {
	case cows < 0:
		return &InvalidCowsError{cows: cows, message: "there are no negative cows"}
	case cows == 0:
		return &InvalidCowsError{cows: cows, message: "no cows don't need food"}
	default:
		return nil
	}
}
