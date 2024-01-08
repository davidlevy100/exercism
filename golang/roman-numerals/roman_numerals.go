// Package romannumerals offers functions for handling roman numerals
package romannumerals

import "errors"

// ToRomanNumeral takes an integer between 1 and 3000 and returns a roman numeral as a string
func ToRomanNumeral(num int) (string, error) {

	thousands := map[int]string{
		3: "MMM",
		2: "MM",
		1: "M",
	}

	hundreds := map[int]string{
		9: "CM",
		8: "DCCC",
		7: "DCC",
		6: "DC",
		5: "D",
		4: "CD",
		3: "CCC",
		2: "CC",
		1: "C",
	}

	tens := map[int]string{
		9: "XC",
		8: "LXXX",
		7: "LXX",
		6: "LX",
		5: "L",
		4: "XL",
		3: "XXX",
		2: "XX",
		1: "X",
	}

	ones := map[int]string{
		9: "IX",
		8: "VIII",
		7: "VII",
		6: "VI",
		5: "V",
		4: "IV",
		3: "III",
		2: "II",
		1: "I",
	}

	var roman string

	if num < 1 || num > 3000 {
		return roman, errors.New("input out of range")
	}

	//thousands
	roman += thousands[num/1000]
	num = num % 1000

	//hundreds
	roman += hundreds[num/100]
	num = num % 100

	//tens
	roman += tens[num/10]
	num = num % 10

	//ones
	roman += ones[num]

	return roman, nil

}
