package pizzaStore

import "fmt"

type Pizza struct {
	name     string
	dough    string
	sauce    string
	toppings []string
}

func (p Pizza) Prepare() {
	fmt.Println(p.name + " 준비")
	fmt.Println(p.dough + " 도우 준비")
	fmt.Println(p.sauce + " 소스 준비")
	for _, t := range p.toppings {
		fmt.Println(t + " 토핑 준비")
	}
}

func (p Pizza) Bake() {
	fmt.Println("피자 굽기")
}

func (p Pizza) Cut() {
	fmt.Println("피자 자르기")
}

func (p Pizza) Box() {
	fmt.Println("피자 박스에 넣기")
}

func (p *Pizza) GetName() string {
	return p.name
}

type NYStyleCheesePizza struct {
	Pizza
}

func NewNyStyleCheesePizza() *NYStyleCheesePizza {
	ny := new(NYStyleCheesePizza)
	ny.dough = "얇은 도우"
	ny.name = "뉴욕 스타일"
	ny.sauce = "마리나라 소스"
	ny.toppings = append(ny.toppings, "레지아노 치즈")
	return ny
}

type ChicagoStyleCheesePizza struct {
	Pizza
}

func NewChicagoStyleCheesePizza() *ChicagoStyleCheesePizza {
	ch := new(ChicagoStyleCheesePizza)
	ch.dough = "두꺼운 도우"
	ch.name = "시카고 스타일"
	ch.sauce = "토마토 소스"
	ch.toppings = append(ch.toppings, "모짜렐라 치즈")
	return ch
}

func (ch ChicagoStyleCheesePizza) Cut() {
	fmt.Println("사각형으로 피자 자르기")
}
