package decorator

type TomatoTopping struct {
	pizza IPizza
}

func (topping *TomatoTopping) getPrice() int {
	return topping.pizza.getPrice() + 5
}
