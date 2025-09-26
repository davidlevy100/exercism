// Package speed simulates a remote-controlled car and race track.
package speed

// Car represents a remote-controlled car.
type Car struct {
	battery, batteryDrain, speed, distance int
}

// NewCar creates a car with full battery.
func NewCar(speed, batteryDrain int) Car {
	return Car{battery: 100, batteryDrain: batteryDrain, speed: speed}
}

// Track represents a race track.
type Track struct {
	distance int
}

// NewTrack creates a track with the given distance.
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive moves the car once if it has enough battery.
func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	}
	car.battery -= car.batteryDrain
	car.distance += car.speed
	return car
}

// CanFinish reports if the car can finish the track.
func CanFinish(car Car, track Track) bool {
	return car.battery/car.batteryDrain >= track.distance/car.speed
}
