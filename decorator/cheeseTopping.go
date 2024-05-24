package decorator

type CheeseTopping struct {
	pizza IPizza
}

func (topping *CheeseTopping) getPrice() int {
	return topping.pizza.getPrice() + 8
}
