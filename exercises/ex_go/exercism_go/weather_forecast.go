// Package main provides tools to forecast the weather conditions 
// of various cities in Globlinocus.
package main

// CurrentCondition represents the actual weather conditions of a place.
var CurrentCondition string

// CurrentLocation represents the location on which the weather conditions
// are checked.
var CurrentLocation string

// Forecast returns a string with the current weather condition in a said location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
