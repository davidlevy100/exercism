// Package resistorcolor provides resistor color codes.
package resistorcolor

var colorMap = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Colors returns the list of all resistor colors in order.
func Colors() []string {
	return []string{
		"black", "brown", "red", "orange", "yellow",
		"green", "blue", "violet", "grey", "white",
	}
}

// ColorCode returns the numeric code of the given color.
func ColorCode(color string) int {
	return colorMap[color]
}
