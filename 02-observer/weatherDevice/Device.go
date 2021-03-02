package weatherDevice

import (
	"fmt"
)

type CurrentConditionsDisplay struct {
	temperature float64
	humidity    float64
	weatherData Subject
}

func NewCurrentConditionsDisplay(weatherData Subject) *CurrentConditionsDisplay {
	ccd := CurrentConditionsDisplay{weatherData: weatherData}
	weatherData.RegisterObserver(&ccd)
	return &ccd
}

func (ccd *CurrentConditionsDisplay) Update(temperature, humidity, pressure float64) {
	ccd.temperature = temperature
	ccd.humidity = humidity
	ccd.Display()
}

func (ccd CurrentConditionsDisplay) Display() {
	temperature := fmt.Sprintf("%f", ccd.temperature)
	humidity := fmt.Sprintf("%f", ccd.humidity)
	fmt.Println("Current conditions: " + temperature + "F degrees and " + humidity + "% humidity")
}
