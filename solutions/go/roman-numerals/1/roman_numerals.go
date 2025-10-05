// Package romannumerals provides conversion between integers and Roman numerals.
package romannumerals

import "errors"

// ToRomanNumeral converts an integer between 1 and 3000 to a Roman numeral string.
func ToRomanNumeral(num int) (string, error) {
	if num < 1 || num > 3999 {
		return "", errors.New("input out of range")
	}

	values := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4,
		1,
	}
	symbols := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I",
	}

	var roman string
	for i := range values {
		for num >= values[i] {
			roman += symbols[i]
			num -= values[i]
		}
	}

	return roman, nil
}
