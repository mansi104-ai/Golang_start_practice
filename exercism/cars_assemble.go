package cars

// CalculateWorkingCarsPerHour calculates how many working cars are produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    x := productionRate
    y := successRate / 100 // Convert success rate to a fraction

    NoOfProduceHour := float64(x) * y
    return NoOfProduceHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    NoOfProduceHour := CalculateWorkingCarsPerHour(productionRate, successRate)
    NoOfProduce := NoOfProduceHour / 60 // Cars produced per minute

    return int(NoOfProduce) // Convert to int before returning
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    totalCarsCost := (carsCount % 10) * 10000 + (carsCount / 10) * 95000
    return uint(totalCarsCost) // Convert to uint before returning
}