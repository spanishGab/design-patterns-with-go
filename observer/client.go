package observer

func ObserverClient() {
	shirtItem := newItem("Nike Shirt")

	observer1 := &Customer{id: "abc@gmail.com"}
	observer2 := &Customer{id: "xyz@gmail.com"}

	shirtItem.subscribe(observer1)
	shirtItem.subscribe(observer2)

	shirtItem.updateAvailability()
}
