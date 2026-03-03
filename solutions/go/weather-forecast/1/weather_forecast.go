// Package weather provides weather forecast functions.
package weather

var (
    // CurrentCondition is the current weather condition.
	CurrentCondition string
    // CurrentLocation is the current location.
	CurrentLocation  string
)

// Forecast returns the local weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
