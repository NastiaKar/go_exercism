package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (0.01 * successRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var perHour = float64(productionRate) * (0.01 * successRate)
    return int(perHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    var cost int = 0
    var tens int = carsCount / 10
    var singles int = carsCount % 10
    cost = (tens * 95000) + (singles * 10000)
    return uint(cost)
}
