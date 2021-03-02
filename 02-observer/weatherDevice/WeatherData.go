package weatherDevice

type WeatherData struct {
	observers   []Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func newWeatherData() *WeatherData {
	wd := WeatherData{observers: []Observer{}}
	return &wd
}

func (wd *WeatherData) RegisterObserver(o Observer) {
	wd.observers = append(wd.observers, o)
}

func (wd *WeatherData) RemoveObserver(o Observer) {
	index := -1
	for i, v := range wd.observers {
		if o == v {
			index = i
		}
	}
	if index > 0 {
		wd.observers = append(wd.observers[:index], wd.observers[index+1:]...)
	}
}

func (wd *WeatherData) NotifyObserver() {
	for _, observer := range wd.observers {
		observer.Update(wd.temperature, wd.humidity, wd.pressure)
	}
}

func (wd *WeatherData) measurementsChanged() {
	wd.NotifyObserver()
}

func (wd *WeatherData) SetMeasurements(temperature, humidity, pressure float64) {
	wd.temperature = temperature
	wd.humidity = humidity
	wd.pressure = pressure
	wd.measurementsChanged()
}
