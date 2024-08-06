package observer

import "fmt"

type Item struct {
	observers []Observer
	name      string
	inStock   bool
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (item *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", item.name)
	item.inStock = true
	item.notify()
}

func (item *Item) subscribe(obsever Observer) {
	item.observers = append(item.observers, obsever)
}

func (item *Item) unsubscribe(obsever Observer) {
	item.observers = removeFromSlice(item.observers, obsever)
}

func (item *Item) notify() {
	for _, observer := range item.observers {
		observer.update(item.name)
	}
}

func removeFromSlice(observers []Observer, observerToRemove Observer) []Observer {
	observersLength := len(observers)
	for i, observer := range observers {
		if observerToRemove.getID() == observer.getID() {
			observers[observersLength-1], observers[i] = observers[i], observers[observersLength-1]
			return observers[:observersLength-1]
		}
	}
	return observers
}
