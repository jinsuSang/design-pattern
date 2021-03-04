package main

import (
	"fmt"

	weatherdevice "github.com/design-pattern/02-observer/weatherdevice"
)

func main() {
	weatherData := weatherdevice.NewWeatherData()

	currentDisplay1 := weatherdevice.NewCurrentConditionsDisplay(weatherData)
	currentDisplay2 := weatherdevice.NewCurrentConditionsDisplay(weatherData)

	currentDisplay1.Display()
	// Current conditions: 0.000000F degrees and 0.000000% humidity
	weatherData.SetMeasurements(80, 65, 30.445)
	// Current conditions: 80.000000F degrees and 65.000000% humidity
	// Current conditions: 80.000000F degrees and 65.000000% humidity
	fmt.Println(weatherData)
	// &{[0xc0000483c0 0xc0000483e0] 80 65 30.445}

	weatherData.RemoveObserver(currentDisplay1)
	fmt.Println(weatherData)
	//&{[0xc0000483e0] 80 65 30.445}

	weatherData.SetMeasurements(75, 55, 45.22)
	// Current conditions: 75.000000F degrees and 55.000000% humidity

	currentDisplay1.Display()
	currentDisplay2.Display()
	//Current conditions: 80.000000F degrees and 65.000000% humidity
	// Current conditions: 75.000000F degrees and 55.000000% humidity
}
