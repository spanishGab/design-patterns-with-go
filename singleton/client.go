package designpatterns

import "fmt"

func SingletonClient() {
	for i := 0;  i < 30; i++ {
		go getInstance()
	}

	instance := getInstance()
	fmt.Println("sum(1, 2)", instance.sum(1, 2))
	fmt.Println("subtract(1, 2)", instance.subtract(1, 2))
	fmt.Println("multiply(1, 2)", instance.multiply(1, 2))
	fmt.Println("divide(1, 2)", instance.divide(1, 2))
	fmt.Scanln()
}