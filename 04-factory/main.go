package main

import (
	"github.com/design-pattern/04-factory/pizzaStore"
)

func main() {
	nyStore := pizzaStore.NYPizzaStore{}
	nyStore.OrderPizza("cheese")
}
