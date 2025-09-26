// Package annalyn implements the logic for Annalyn's infiltration scenario.
// It provides functions that determine whether certain actions are possible
// based on the states of the knight, archer, prisoner, and Annalyn's pet dog.
package annalyn

// CanFastAttack reports whether a fast attack is possible.
// It is only possible if the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// CanSpy reports whether spying is possible.
// It is possible if at least one character is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

// CanSignalPrisoner reports whether the prisoner can be signaled.
// This requires the prisoner to be awake while the archer is asleep.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return !archerIsAwake && prisonerIsAwake
}

// CanFreePrisoner reports whether the prisoner can be freed.
// This is possible in either of two cases:
//  1. The prisoner is awake, and both the knight and archer are asleep.
//  2. Annalyn's pet dog is present, and the archer is asleep.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	return (prisonerIsAwake && !knightIsAwake && !archerIsAwake) ||
		(petDogIsPresent && !archerIsAwake)
}
