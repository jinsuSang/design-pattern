package pizzaStore

type PizzaStore interface {
	createPizza(types string) *Pizza
	OrderPizza(types string) *Pizza
}

type NYPizzaStore struct{}

func (ny NYPizzaStore) createPizza(types string) *Pizza {
	switch types {
	case "cheese":
		return &NewNyStyleCheesePizza().Pizza
	default:
		return nil
	}
}

func (ny NYPizzaStore) OrderPizza(types string) *Pizza {

	pizza := ny.createPizza(types)

	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()

	return pizza
}
