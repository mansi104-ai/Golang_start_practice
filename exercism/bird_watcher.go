package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	// panic("Please implement the TotalBirdCount() function")
    var x int = 0
    for i := 0 ;i <= len(birdsPerDay)-1;i++{
        x += birdsPerDay[i]
    }
    return x
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// panic("Please implement the BirdsInWeek() function")
    startIndex := (week - 1) * 7
	endIndex := startIndex + 7

	// Initialize total count
	totalBirds := 0

	// Sum up the bird counts for the specified week
	for i := startIndex; i < endIndex; i++ {
		totalBirds += birdsPerDay[i]
	}

	return totalBirds
    
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	// panic("Please implement the FixBirdCountLog() function")
    for i:= 0 ;i < len(birdsPerDay)-1 ; i+=2{
        birdsPerDay[i]+= 1
        
    }
    return birdsPerDay
    
}