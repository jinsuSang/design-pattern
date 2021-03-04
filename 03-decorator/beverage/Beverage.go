package beverage

type Beverage struct {
	description string
	cost        float64
}

func (b Beverage) GetDescription() string {
	return b.description
}

func (b Beverage) Cost() float64 {
	return b.cost
}

func NewBeverage() *Beverage {
	nb := &Beverage{description: "제목 없음"}
	return nb
}

type Espresso struct {
	Beverage
}

func (e Espresso) Cost() float64 {
	return e.cost
}

func (e Espresso) GetDescription() string {
	return e.description
}

func NewEspresso() *Espresso {
	espresso := &Espresso{Beverage: Beverage{description: "에스프레소", cost: 1.99}}
	return espresso
}

type HouseBlend struct {
	Beverage
}

func (hb HouseBlend) Cost() float64 {
	return hb.cost
}

func (hb HouseBlend) GetDescription() string {
	return hb.description
}

func NewHouseBlend() *HouseBlend {
	houseBlend := &HouseBlend{Beverage: Beverage{description: "하우스 블렌드 커피", cost: 8.9}}
	return houseBlend
}
