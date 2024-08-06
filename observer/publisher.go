package observer

type Publisher interface {
	subscribe(observer Observer)
	unsubscribe(observer Observer)
	notify()
}
