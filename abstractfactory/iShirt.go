package abstractfactory

type IShirt interface {
	setBrand(brand string)
	setSize(size int)
	getBrand() string
	getSize() int
}

type Shirt struct {
	brand string
	size int
}

func (shirt *Shirt) setBrand(brand string) {
	shirt.brand = brand
}

func (shirt *Shirt) setSize(size int) {
	shirt.size = size
}

func (shirt *Shirt) getBrand() string {
	return shirt.brand
}

func (shirt *Shirt) getSize() int {
	return shirt.size
}
