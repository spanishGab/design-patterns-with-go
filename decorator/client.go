package decorator

import "fmt"

func DecoratorClient() {
	pizza := &VeggieMania{}

	pizzaWithTopping := &Topping{
		pizza: pizza,
	}

	pizzaWithCheese := &CheeseTopping{
		pizza: pizzaWithTopping,
	}

	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}