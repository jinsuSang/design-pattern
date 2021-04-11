package duck

import "fmt"

type QuackBehavior interface {
	Quack()
}

type Quack struct{}

func (q Quack) Quack() {
	fmt.Println("꽥")
}

type MuteQuack struct{}

func (m MuteQuack) Quack() {
	fmt.Println("---")
}

type Squeak struct{}

func (s Squeak) Quack() {
	fmt.Println("삑")
}
