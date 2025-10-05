// Package space converts a given age in seconds to years on different planets.
package space

// Planet represents the name of a planet in the solar system.
type Planet string

const secsInOneYear = 31557600.0

var planetPeriods = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age returns the age in planetary years on planet p, or -1 if the planet is invalid.
func Age(secs float64, p Planet) float64 {
	period, ok := planetPeriods[p]
	if !ok {
		return -1
	}
	return secs / (secsInOneYear * period)
}
