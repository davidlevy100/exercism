// Package birdwatcher provides utilities for analyzing daily bird counts.
package birdwatcher

// TotalBirdCount returns the total number of birds counted
// by summing the counts for each day.
func TotalBirdCount(birdsPerDay []int) int {
	var sum int
	for _, value := range birdsPerDay {
		sum += value
	}
	return sum
}

// BirdsInWeek returns the total number of birds counted
// in the given week. Weeks are 1-indexed and consist of
// seven consecutive days.
func BirdsInWeek(birdsPerDay []int, week int) int {
	end := 7 * week
	start := end - 7
	return TotalBirdCount(birdsPerDay[start:end])
}

// FixBirdCountLog returns the bird counts after correcting
// the counts for alternate days. It increments the count
// for every second day, starting with the first day.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}
