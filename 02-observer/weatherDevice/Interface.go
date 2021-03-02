package weatherdevice

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObserver()
}

type Observer interface {
	Update(temp, humidity, pressure float64)
}

type DisplayElement interface {
	Display()
}
