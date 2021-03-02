package main

import (
	"fmt"

	weatherdevice "github.com/design-pattern/02-observer/weatherdevice"
)

func main() {
	weatherData := weatherdevice.NewWeatherData()
	currentDisplay := weatherdevice.NewCurrentConditionsDisplay(weatherData)
	weatherData.SetMeasurements(80, 65, 30.445)
	fmt.Println(weatherData)
	currentDisplay.Display()
	weatherData.RemoveObserver(currentDisplay)
	fmt.Println(weatherData)
}
