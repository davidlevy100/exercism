// Package twelve provides the lyrics for "The Twelve Days of Christmas".
package twelve

import "strings"

var ordinals = []string{
	"", // 0th element unused — aligns day numbers with indices
	"first", "second", "third", "fourth", "fifth", "sixth",
	"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth",
}

var gifts = []string{
	"", // 0th element unused
	"a Partridge in a Pear Tree",
	"two Turtle Doves, and",
	"three French Hens,",
	"four Calling Birds,",
	"five Gold Rings,",
	"six Geese-a-Laying,",
	"seven Swans-a-Swimming,",
	"eight Maids-a-Milking,",
	"nine Ladies Dancing,",
	"ten Lords-a-Leaping,",
	"eleven Pipers Piping,",
	"twelve Drummers Drumming,",
}

// Verse returns the verse for the given day (1–12) of "The Twelve Days of Christmas".
func Verse(day int) string {
	parts := make([]string, 0, day)
	for j := day; j >= 1; j-- {
		parts = append(parts, gifts[j])
	}
	return "On the " + ordinals[day] +
		" day of Christmas my true love gave to me: " +
		strings.Join(parts, " ") + "."
}

// Song returns the full lyrics with all 12 verses separated by newlines.
func Song() string {
	lines := make([]string, 0, 12)
	for i := 1; i <= 12; i++ {
		lines = append(lines, Verse(i))
	}
	return strings.Join(lines, "\n")
}
