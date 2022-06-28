// Package weather
// about this package.
package weather

// CurrentCondition else.
var CurrentCondition string

// CurrentLocation other.
var CurrentLocation string

// Forecast function.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
