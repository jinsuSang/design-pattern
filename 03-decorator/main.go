package main

import (
	"fmt"

	"github.com/design-pattern/03-decorator/beverage"
)

func main() {
	espresso := beverage.NewEspresso()

	fmt.Println(espresso.Description(), espresso.Cost())

	mocha := beverage.NewMocha(espresso)
	fmt.Println(mocha.Description(), mocha.Cost())

	mocha = beverage.NewMocha(mocha)
	fmt.Println(mocha.Description(), mocha.Cost())

	soy := beverage.NewSoy(mocha)
	fmt.Println(soy.Description(), soy.Cost())
	// Output:
	// 에스프레소 1.99
	// 에스프레소, 모카 2.19
	// 에스프레소, 모카, 모카 2.39
	// 에스프레소, 모카, 모카, 두유 2.69
}
