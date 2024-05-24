package decorator

type VeggieMania struct {
}

func (pizza *VeggieMania) getPrice() int {
	return 15
}
