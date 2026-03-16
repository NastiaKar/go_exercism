// Package weather is used for forecasting weather by using current location and condition.
package weather

var (
    // CurrentCondition shows the current weather.
	CurrentCondition string
    // CurrentLocation shows the current location.
	CurrentLocation  string
)

// Forecast function shows a message about the current weather in the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
