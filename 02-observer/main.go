package main

import (
	"fmt"

	"github.com/design-pattern/02-observer/weatherDevice"
)

func main() {
	weatherData := weatherDevice.WeatherData{}
	currentDisplay := weatherDevice.NewCurrentConditionsDisplay(&weatherData)
	weatherData.SetMeasurements(80, 65, 30.445)
	fmt.Println(weatherData)
	currentDisplay.Display()
}
