package designpatterns

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type simpleCalculator struct {}

var instance *simpleCalculator

func getInstance() *simpleCalculator {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating single instance now.")
			instance = &simpleCalculator{}
		} else {
			fmt.Println("1 - Single instance already created.")
		}
	} else {
		fmt.Println("2 - Single instance already created.")
	}
	return instance
}

func (calc *simpleCalculator) sum(a, b int) int {
	return a + b
}

func (calc *simpleCalculator) subtract(a, b int) int {
	return a - b
}

func (calc *simpleCalculator) multiply(a, b int) int {
	return a * b
}

func (calc *simpleCalculator) divide(a, b int) float32 {
	if b == 0 {
		panic("Cannot divide by zero")
	}
	return float32(a) / float32(b)
}