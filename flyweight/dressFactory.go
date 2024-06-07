package flyweight

import "fmt"

const (
	TerroristDressType = "tDress"
	CounterTerroristDressType = "ctDress"
)

var (
	dressFactorySingleInstante = &DressFactory{
		dressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	dressMap map[string]Dress
}

func (dressFactory *DressFactory) getDressByType(dressType string) (Dress, error) {
	if dressFactory.dressMap[dressType] != nil {
		return dressFactory.dressMap[dressType], nil
	}
	if dressType == TerroristDressType {
		fmt.Printf("Creating terrorist dress type %s\n", dressType)
		dressFactory.dressMap[dressType] = newTerroristDress()
		return dressFactory.dressMap[dressType], nil
	}
	if dressType == CounterTerroristDressType {
		fmt.Printf("Creating counter-terrorist dress type %s\n", dressType)
		dressFactory.dressMap[dressType] = newCounterTerroristDress()
		return dressFactory.dressMap[dressType], nil
	}
	return nil, fmt.Errorf("Wrong dress type passed")
}

func getDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstante
}