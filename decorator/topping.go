package decorator

type Topping struct {
	pizza IPizza
}

func (topping *Topping) getPrice() int {
	return topping.pizza.getPrice() + 4
}
