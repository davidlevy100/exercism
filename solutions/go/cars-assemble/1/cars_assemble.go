// Package cars provides simple calculations for car production rates and costs.
package cars

const (
	// singleCarCost is the cost of producing one car.
	singleCarCost = 10000

	// tenCarCost is the discounted cost of producing 10 cars as a batch.
	tenCarCost = 95000
)

// CalculateWorkingCarsPerHour returns the number of successful cars produced per hour.
// successRate is expressed as a percentage (e.g. 90.0 means 90%).
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100.0
}

// CalculateWorkingCarsPerMinute returns the whole number of successful cars
// produced per minute. Fractional cars are truncated.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60.0)
}

// CalculateCost returns the total production cost for the given number of cars.
// Groups of 10 cars are charged at a discounted rate.
func CalculateCost(carsCount int) uint {
	return (uint(carsCount)/10)*tenCarCost + (uint(carsCount)%10)*singleCarCost
}
