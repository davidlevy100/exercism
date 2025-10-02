// Package expenses provides utilities to filter and total expense records.
package expenses

import "errors"

// Record represents an expense record with a day, amount, and category.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod defines a range of days [From, To] inclusive.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns a slice of records for which predicate(r) is true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	result := make([]Record, 0, len(in))
	for _, record := range in {
		if predicate(record) {
			result = append(result, record)
		}
	}
	return result
}

// ByDaysPeriod returns a predicate that reports whether
// a record’s day lies within the given period.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns a predicate that reports whether
// a record’s category matches c.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns the sum of expenses within the given period.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	sum := 0.0
	pred := ByDaysPeriod(p) // build predicate once
	for _, record := range in {
		if pred(record) {
			sum += record.Amount
		}
	}
	return sum
}

// CategoryExpenses returns the sum of expenses for category c
// within the given period. If no records belong to c, an error is returned.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	sum := 0.0
	catPred := ByCategory(c)
	dayPred := ByDaysPeriod(p)
	catMatches := 0

	for _, record := range in {
		if catPred(record) {
			catMatches++
			if dayPred(record) {
				sum += record.Amount
			}
		}
	}

	if catMatches == 0 {
		return 0, errors.New("no category matches")
	}
	return sum, nil
}
