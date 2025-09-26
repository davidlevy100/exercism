// Package weather provides simple utilities for reporting weather conditions.
package weather

// CurrentCondition is the latest reported weather condition.
var CurrentCondition string

// CurrentLocation is the location for the latest forecast.
var CurrentLocation string

// Forecast updates the current location and condition, and returns a summary.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return city + " - current weather condition: " + condition
}
