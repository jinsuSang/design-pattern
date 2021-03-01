package duck

import "fmt"

// FlyBehavior 는 모든 나는 행동을 클래스에서 구현한다.
type FlyBehavior interface {
	Fly()
}

// FlyWithWings 오리들이 나는 행동을 구현
type FlyWithWings struct{}

func (f FlyWithWings) Fly() {
	fmt.Println("fly")
}

// FlyNoWay 날지 못하는 오리들의 행동 구현
type FlyNoWay struct{}

func (f FlyNoWay) Fly() {
	fmt.Println("can't fly")
}

type FlyRocketPowered struct{}

func (f FlyRocketPowered) Fly() {
	fmt.Println("Rocket Launch!")
}
