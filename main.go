package main

import (
	abstractfactory "example/designpatterns/abstractfactory"
	"example/designpatterns/builder"
	factory "example/designpatterns/factorymethod"
	prototype "example/designpatterns/prototype"
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
		case "abstractfactory":
			abstractfactory.AbstractFactoryClient()
			break
		case "builder":
			builder.BuilderClient()
			break
		case "prototype":
			prototype.PrototypeClient()
			break
		default:
			fmt.Println("Pattern not found")
			break
	}
}