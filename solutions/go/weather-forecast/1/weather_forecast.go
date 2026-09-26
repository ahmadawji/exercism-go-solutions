// Package weather provides variables and a function for information about Goblinocus's weather.
package weather

var (
// CurrentCondition presents the weather condition.
	CurrentCondition string
// CurrentLocation presents the city where the weather is taken.
	CurrentLocation  string
)

// Forecast takes current location and condtion and returs a formatted output.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
