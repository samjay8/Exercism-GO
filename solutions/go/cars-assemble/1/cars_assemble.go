package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	percent := successRate / 100
	return float64(productionRate) * percent

}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	productionperhour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(productionperhour) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	cargroups := carsCount / 10
	car_remainders := carsCount % 10
	return uint(cargroups)*95000 + uint(car_remainders)*10000
}
