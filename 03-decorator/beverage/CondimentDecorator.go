package beverage

type Mocha struct {
	b IBeverage
}

func (m Mocha) Description() string {
	return m.b.Description() + ", 모카"
}

func (m Mocha) Cost() float64 {
	return .20 + m.b.Cost()
}

func NewMocha(b interface{}) Mocha {
	return Mocha{b: b.(IBeverage)}
}

type Soy struct {
	b IBeverage
}

func (s Soy) Description() string {
	return s.b.Description() + ", 두유"
}

func (s Soy) Cost() float64 {
	return .30 + s.b.Cost()
}

func NewSoy(b interface{}) *Soy {
	return &Soy{b: b.(IBeverage)}
}
