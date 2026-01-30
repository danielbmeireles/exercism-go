// Package weather provides a way to get weather forecasts for different locations.
package weather

var (
	// CurrentCondition holds the current weather condition.
	CurrentCondition string
	// CurrentLocation holds the current location for the weather forecast.
	CurrentLocation  string
)

// Forecast returns a formatted string with the current weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
