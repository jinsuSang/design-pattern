package beverage

type ICondimentDecorator interface {
	GetDescription()
}

type CondimentDecorator struct {
	Beverage
}

type Mocha struct {
	CondimentDecorator
}

func (m *Mocha) GetDescription() string {
	return m.CondimentDecorator.GetDescription() + ", 모카"
}

func (m *Mocha) Cost() float64 {
	return .20 + m.CondimentDecorator.Cost()
}

type Soy struct {
	CondimentDecorator
}

func (s *Soy) GetDescription() string {
	return s.CondimentDecorator.GetDescription() + ", 두유"
}

func (s *Soy) Cost() float64 {
	return .20 + s.CondimentDecorator.Cost()
}

type Whip struct {
	CondimentDecorator
}

func (w *Whip) GetDescription() string {
	return w.Beverage.GetDescription() + ", 휘핑 크림"
}

func (w *Whip) Cost() float64 {
	return .20 + w.CondimentDecorator.Cost()
}

type SteamMilk struct {
	CondimentDecorator
}

func (sm SteamMilk) GetDescription() string {
	return sm.Beverage.GetDescription() + ", 스팀 밀크"
}

func (sm SteamMilk) Cost() float64 {
	return .20 + sm.Beverage.cost
}
