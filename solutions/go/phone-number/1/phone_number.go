// Package phonenumber normalizes and formats NANP phone numbers.
package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

var phonePattern = regexp.MustCompile(`^(?:\+?1)?\D*([2-9]\d{2})\D*([2-9]\d{2})\D*(\d{4})\D*$`)

// Parts holds parsed phone number segments.
type Parts struct {
	Area, Exchange, Line string
}

// parse extracts parts from a valid NANP number.
func parse(s string) (*Parts, error) {
	m := phonePattern.FindStringSubmatch(s)
	if m == nil {
		return nil, errors.New("invalid phone number")
	}
	return &Parts{m[1], m[2], m[3]}, nil
}

// Number returns digits only (area+exchange+line).
func Number(s string) (string, error) {
	p, err := parse(s)
	if err != nil {
		return "", err
	}
	return p.Area + p.Exchange + p.Line, nil
}

// AreaCode returns the 3-digit area code.
func AreaCode(s string) (string, error) {
	p, err := parse(s)
	if err != nil {
		return "", err
	}
	return p.Area, nil
}

// Format returns the number as "(AAA) BBB-CCCC".
func Format(s string) (string, error) {
	p, err := parse(s)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", p.Area, p.Exchange, p.Line), nil
}
