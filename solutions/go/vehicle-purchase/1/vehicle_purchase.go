package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	vehicle := "is clearly the better choice."

	if option1 < option2 {
		vehicle = option1 + " " + vehicle
	}
	if option1 > option2 {
		vehicle = option2 + " " + vehicle
	}
	return vehicle
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {

	var calculate float64
	if age < 3 {
		calculate = originalPrice * 0.8
	} else if age >= 10 {
		calculate = originalPrice * 0.5
	} else if age >= 3 && age < 10 {
		calculate = originalPrice * 0.7
	}
	return calculate
}
