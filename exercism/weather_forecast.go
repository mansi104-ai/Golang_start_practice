//Package weather provides tools to display.
package weather

// CurrentCondition tells the current condition of the weather.
var CurrentCondition string

// CurrentLocation tells the current location of consideration.
var CurrentLocation string


// Forecast helps to display the weather condition at the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}