package partyrobot
import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	// panic("Please implement the Welcome function")
    string1 := fmt.Sprintf("Welcome to my party, %s!",name)
    return string1 
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
    // Correctly format the birthday message
    message := fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
    return message
}


// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
// 	panic("Please implement the AssignTable function")
     // Format the table number to be 3 digits with leading zeros
    tableNumberFormatted := fmt.Sprintf("%03d", table)
    
    // Format the distance to one decimal place
    distanceFormatted := fmt.Sprintf("%.1f", distance)
    
    // Construct the welcome message
    message := fmt.Sprintf("Welcome to my party, %s!\n", name) +
        fmt.Sprintf("You have been assigned to table %s. Your table is %s, exactly %s meters from here.\n", 
            tableNumberFormatted, direction, distanceFormatted) +
        fmt.Sprintf("You will be sitting next to %s.", neighbor)

    return message
    
}