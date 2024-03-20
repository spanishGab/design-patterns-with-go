package main

import (
	factory "example/designpatterns/factorymethod"
	singleton "example/designpatterns/singleton"
	"fmt"
	"os"
)

func main() {
	var chosenDesignPattern string = os.Args[1]
	switch (chosenDesignPattern) {
		case "singleton":
			singleton.SingletonClient()
			break
		case "factory":
			factory.FactoryClient()
			break
		default:
			fmt.Println("Pattern not found")
			break
	}
}