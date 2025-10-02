// Package twofer provides a function to share something between
// another person and yourself.
package twofer

// ShareWith returns a phrase of the form
// "One for <name>, one for me." if a name is provided,
// or "One for you, one for me." if the name is empty.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
