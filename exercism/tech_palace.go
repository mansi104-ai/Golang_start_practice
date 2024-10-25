package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	// panic("Please implement the WelcomeMessage() function")
    var final = "Welcome to the Tech Palace,"+ " " + strings.ToUpper(customer)
    return final 
    
}

// AddBorder adds a border to a welcome message.
// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    // Create the top and bottom border
    border := strings.Repeat("*", numStarsPerLine) // Create border without newline
    
    // Construct the final message with borders
    final := border + "\n" + welcomeMsg + "\n" + border // Add newlines around welcomeMsg
    return final
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// panic("Please implement the CleanupMessage() function")
    // Remove all stars from the message
    cleanedMessage := strings.ReplaceAll(oldMsg, "*", "")
    // Trim leading and trailing whitespace
    cleanedMessage = strings.TrimSpace(cleanedMessage)
    return cleanedMessage
}
    
    
