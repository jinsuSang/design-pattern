package main

import (
	"fmt"

	"github.com/design-pattern/03-decorator/beverage"
)

func main() {
	espresso := beverage.NewEspresso()

	fmt.Println(espresso.GetDescription() + " $" + fmt.Sprint(espresso.Cost()))

	hb := beverage.NewHouseBlend()
	mocha := beverage.Mocha{CondimentDecorator: beverage.CondimentDecorator(*hb)}
	fmt.Println(mocha)
	mocha = beverage.Mocha{CondimentDecorator: mocha.CondimentDecorator}
	fmt.Println(mocha)
	fmt.Println(mocha.GetDescription() + " $" + fmt.Sprint(mocha.Cost()))

	// 	에스프레소 $1.99
	// {{{하우스 블렌드 커피 8.9}}}
	// {{{하우스 블렌드 커피 8.9}}}
	// 하우스 블렌드 커피, 모카 $9.1
	// 이유가 뭘까??
}
