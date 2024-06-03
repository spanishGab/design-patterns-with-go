package main

import (
	abstractfactory "example/designpatterns/abstractfactory"
	"example/designpatterns/adapter"
	"example/designpatterns/bridge"
	"example/designpatterns/builder"
	"example/designpatterns/composite"
	"example/designpatterns/decorator"
	"example/designpatterns/facade"
	factory "example/designpatterns/factorymethod"
	prototype "example/designpatterns/prototype"
	singleton "example/designpatterns/singleton"
	"fmt"
	"os"
)

func main() {
	var chosenDesignPattern string = os.Args[1]
	switch (chosenDesignPattern) {
		case "abstractfactory":
			abstractfactory.AbstractFactoryClient()
			break
		case "adapter":
			adapter.AdapterClient()
			break
		case "bridge":
			bridge.BridgeClient()
			break
		case "builder":
			builder.BuilderClient()
			break
		case "composite":
			composite.CompositeClient()
			break
		case "decorator":
			decorator.DecoratorClient()
			break
		case "facade":
			facade.FacadeClient()
			break
		case "factory":
			factory.FactoryClient()
			break
		case "prototype":
			prototype.PrototypeClient()
			break
		case "singleton":
			singleton.SingletonClient()
			break
		default:
			fmt.Println("Pattern not found")
			break
	}
}