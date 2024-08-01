package main

import (
	abstractfactory "example/designpatterns/abstractfactory"
	"example/designpatterns/adapter"
	"example/designpatterns/bridge"
	"example/designpatterns/builder"
	"example/designpatterns/chainofresponsibility"
	"example/designpatterns/command"
	"example/designpatterns/composite"
	"example/designpatterns/decorator"
	"example/designpatterns/facade"
	factory "example/designpatterns/factorymethod"
	"example/designpatterns/flyweight"
	"example/designpatterns/iterator"
	prototype "example/designpatterns/prototype"
	"example/designpatterns/proxy"
	singleton "example/designpatterns/singleton"
	"fmt"
	"os"
)

func main() {
	var chosenDesignPattern string = os.Args[1]
	switch chosenDesignPattern {
	case "abstractfactory":
		abstractfactory.AbstractFactoryClient()
	case "adapter":
		adapter.AdapterClient()
	case "bridge":
		bridge.BridgeClient()
	case "builder":
		builder.BuilderClient()
	case "cor":
		chainofresponsibility.CoRClient()
	case "command":
		command.CommandClient()
	case "composite":
		composite.CompositeClient()
	case "decorator":
		decorator.DecoratorClient()
	case "facade":
		facade.FacadeClient()
	case "factory":
		factory.FactoryClient()
	case "flyweight":
		flyweight.FlyweightClient()
	case "iterator":
		iterator.IteratorClient()
	case "prototype":
		prototype.PrototypeClient()
	case "proxy":
		proxy.ProxyClient()
	case "singleton":
		singleton.SingletonClient()
	default:
		fmt.Println("Pattern not found")
	}
}
