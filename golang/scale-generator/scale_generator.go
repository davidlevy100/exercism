// Package scale generates musical scales
package scale

import "strings"

var sharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

var jumps = map[string]int{"A": 3, "M": 2, "m": 1}

var scales = map[string][]string{
	"C":  sharps,
	"a":  sharps,
	"G":  sharps,
	"D":  sharps,
	"A":  sharps,
	"E":  sharps,
	"B":  sharps,
	"F#": sharps,
	"e":  sharps,
	"b":  sharps,
	"f#": sharps,
	"c#": sharps,
	"g#": sharps,
	"d#": sharps,
	"F":  flats,
	"Bb": flats,
	"Eb": flats,
	"Ab": flats,
	"Db": flats,
	"Gb": flats,
	"d":  flats,
	"g":  flats,
	"c":  flats,
	"f":  flats,
	"bb": flats,
	"eb": flats,
}

// Scale takes tonic and interval as strings and returns notes represented as a slice of strings
func Scale(tonic, intervals string) []string {

	result := []string{}

	thisScale := scales[tonic]

	index := indexOf(strings.Title(tonic), thisScale)

	// No intervals just return 12 notes starting from the first note given
	if len(intervals) == 0 {

		for i := index; i < index+12; i++ {
			result = append(result, getNote(i, thisScale))
		}

		return result

	}

	for _, interval := range intervals {
		result = append(result, getNote(index, thisScale))
		index += jumps[string(interval)]
	}

	return result
}

func indexOf(myString string, stringSlice []string) int {

	for index, value := range stringSlice {
		if value == myString {
			return index
		}
	}
	return -1
}

func getNote(index int, scale []string) string {
	return scale[index%len(scale)]
}
