// Package jedlik implements a simple electric car simulator
// used to practice basic Go syntax and methods.
package jedlik

import "fmt"

// Drive advances the car by one step of speed,
// consuming batteryDrain from the battery if enough charge remains.
func (c *Car) Drive() {
	if c.battery >= c.batteryDrain {
		c.battery -= c.batteryDrain
		c.distance += c.speed
	}
}

// DisplayDistance returns the car's total distance
// as "Driven NN meters".
func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery returns the car's remaining battery
// as "Battery at NN%".
func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish reports whether the car can finish a track
// of the given distance with its remaining battery.
func (c *Car) CanFinish(trackDistance int) bool {
	return c.battery/c.batteryDrain*c.speed >= trackDistance
}
