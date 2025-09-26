// Package partyrobot generates friendly party messages.
package partyrobot

import "fmt"

// Welcome returns a simple welcome message.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday returns a birthday message with the person's age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable returns a table assignment message with details.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return fmt.Sprintf(
		"Welcome to my party, %s!\n"+
			"You have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\n"+
			"You will be sitting next to %s.",
		name, table, direction, distance, neighbor,
	)
}
