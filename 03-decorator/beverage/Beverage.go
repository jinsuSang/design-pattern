package beverage

type Beverage struct {
	description string
	cost        float64
}

type IBeverage interface {
	Cost() float64
	Description() string
}

type Espresso struct {
	beverage Beverage
}

func (e Espresso) Cost() float64 {
	return e.beverage.cost
}

func (e Espresso) Description() string {
	return e.beverage.description
}

func NewEspresso() *Espresso {
	return &Espresso{
		beverage: Beverage{cost: 1.99, description: "에스프레소"},
	}
}

type HouseBlend struct {
	beverage Beverage
}

func (hb HouseBlend) Cost() float64 {
	return hb.beverage.cost
}

func (hb HouseBlend) Description() string {
	return hb.beverage.description
}

func NewHouseBlend() *HouseBlend {
	return &HouseBlend{
		beverage: Beverage{cost: 8.9, description: "하우스 블렌드"},
	}
}
