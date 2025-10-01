// Package sorting provides examples of working with interfaces,
// type assertions, and string/number conversions. It defines
// simple types (NumberBox and FancyNumberBox) and functions
// that return human-readable descriptions of their contents.
package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber returns a string describing the given float64.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

// NumberBox is a type that can return an integer value.
type NumberBox interface {
	Number() int
}

// DescribeNumberBox returns a string describing the number
// stored in a NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

// FancyNumber wraps a string representation of a number.
type FancyNumber struct {
	n string
}

// Value returns the string value of the FancyNumber.
func (i FancyNumber) Value() string {
	return i.n
}

// FancyNumberBox is a type that can return a string value.
type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber returns the integer value of a FancyNumber.
// It returns 0 if the FancyNumberBox is not a FancyNumber or if
// the string cannot be converted to an integer.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if val, ok := fnb.(FancyNumber); ok {
		num, err := strconv.Atoi(val.n)
		if err == nil {
			return num
		}
	}
	return 0
}

// DescribeFancyNumberBox returns a string describing the number
// stored in a FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything returns a string describing the given value.
// It handles ints, float64s, NumberBoxes, and FancyNumberBoxes.
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
	case int:
		return DescribeNumber(float64(v))
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	default:
		return "Return to sender"
	}
}
