package main

import "github.com/design-pattern/01-strategy/duck"

func main() {
	mallard := duck.MallardDuck{
		Duck: duck.Duck{
			FlyBehavior:   duck.FlyWithWings{},
			QuackBehavior: duck.Quack{}}}
	mallard.Display()
	mallard.PerformFly()
	mallard.PerformQuack()

	// Output:
	// 물오리
	// fly
	// 꽥

	model := duck.ModelDuck{
		Duck: duck.Duck{
			FlyBehavior:   duck.FlyNoWay{},
			QuackBehavior: duck.Quack{},
		},
	}

	model.PerformFly()
	model.SetFlyBehavior(duck.FlyRocketPowered{})
	model.PerformFly()

	// Output:
	// can't fly
	// Rocket Launch!
}
