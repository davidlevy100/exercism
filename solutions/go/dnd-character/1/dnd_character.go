// Package dndcharacter generates random D&D characters.
package dndcharacter

import "math/rand"

// Character holds D&D ability scores and hit points.
type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier returns the D&D ability modifier for a score,
// rounding down toward negative infinity.
func Modifier(score int) int {
	mod := (score - 10) / 2
	if (score-10)%2 != 0 && score < 10 {
		mod--
	}
	return mod
}

// Ability rolls 4d6 and returns the sum of the highest three.
func Ability() int {
	lowest, sum := 7, 0
	for range 4 {
		roll := rand.Intn(6) + 1
		sum += roll
		if roll < lowest {
			lowest = roll
		}
	}
	return sum - lowest
}

// GenerateCharacter creates a random character with ability scores
// and hit points based on Constitution.
func GenerateCharacter() Character {
	c := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}
	c.Hitpoints = 10 + Modifier(c.Constitution)
	return c
}
