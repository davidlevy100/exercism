// Package chance provides simple randomization utilities such as dice rolls,
// energy generation, and shuffling of animals.
package chance

import "math/rand"

// RollADie returns a random integer between 1 and 20 inclusive.
func RollADie() int {
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 between 0.0 (inclusive)
// and 12.0 (exclusive).
func GenerateWandEnergy() float64 {
	return rand.Float64() * 12
}

// ShuffleAnimals returns a slice containing all eight animal strings
// in random order.
func ShuffleAnimals() []string {
	x := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

	rand.Shuffle(len(x), func(i, j int) {
		x[i], x[j] = x[j], x[i]
	})

	return x
}
