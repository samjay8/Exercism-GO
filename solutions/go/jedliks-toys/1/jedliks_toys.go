package jedlik

import (
	"fmt"
)

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if car.battery < car.batteryDrain {
		return
	}
	car.distance += car.speed
	car.battery -= car.batteryDrain
	fmt.Println("car is now", car)
}

// TODO: define the 'DisplayDistance() string' method
func (car *Car) DisplayDistance() string {
	distance := fmt.Sprintf("Driven %d meters", car.distance)
	return distance
}

// TODO: define the 'DisplayBattery() string' method
func (car *Car) DisplayBattery() string {
	battery := fmt.Sprintf("Battery at %d%%", car.battery)
	return battery
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car *Car) CanFinish(trackDistance int) bool {
	if car.battery == 0 {
		return false
	}
	if car.batteryDrain == 0 {
		return true
	}
	maxDistance := (car.battery / car.batteryDrain) * car.speed
	return maxDistance >= trackDistance
}
