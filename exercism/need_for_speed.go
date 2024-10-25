package speed

// TODO: define the 'Car' type struct
type Car struct{
    battery      int // battery percentage
	batteryDrain int // percentage drained per drive
	speed        int // speed in meters
	distance     int // distance driven
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	// Return a new Car instance with the given speed and battery drain
	return Car{
		battery:      100, // Initial battery is 100%
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,    // Initial distance is 0
	}
}
// TODO: define the 'Track' type struct
type Track struct{
    distance int
    
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	// panic("Please implement the NewTrack function")
    return Track{
        distance : distance,
    }
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	// Check if the battery is enough for one drive
	if car.battery >= car.batteryDrain {
		// Reduce the battery and increase the distance
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
	return car
}
// CanFinish checks if a car is able to finish a certain track.
// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
    // Calculate the total number of drives needed to cover the track distance
    distanceToCover := track.distance - car.distance
    
    // Calculate how many times the car can drive with the remaining battery
    totalDrivesPossible := car.battery / car.batteryDrain
    
    // Calculate how many drives are needed to finish the race
    drivesNeeded := (distanceToCover + car.speed - 1) / car.speed // rounding up
    
    // If the car can make at least as many drives as needed, return true
    return totalDrivesPossible >= drivesNeeded
}