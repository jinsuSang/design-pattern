package duck

import "fmt"

// Duck 은 오리 클래스
type Duck struct {
	FlyBehavior   FlyBehavior
	QuackBehavior QuackBehavior
	Display       func()
}

func (d Duck) PerformFly() {
	d.FlyBehavior.Fly()
}

func (d Duck) PerformQuack() {
	d.QuackBehavior.Quack()
}

func (d Duck) Swim() {
	fmt.Println("수영")
}

func (d *Duck) SetFlyBehavior(fb FlyBehavior) {
	d.FlyBehavior = fb
}

func (d *Duck) SetQuackBehavior(qb QuackBehavior) {
	d.QuackBehavior = qb
}

type MallardDuck struct {
	Duck
}

func (m MallardDuck) Display() {
	fmt.Println("물오리")
}

type ModelDuck struct {
	Duck
}

func (m ModelDuck) Display() {
	fmt.Println("모형 오리")
}
