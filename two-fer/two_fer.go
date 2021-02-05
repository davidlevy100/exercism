// Package twofer adds names to strings
package twofer

import "fmt"

// ShareWith inserts a name into a string (One for <name>, one for me.)
func ShareWith(name string) string {

	if len(name) == 0 {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)

}
