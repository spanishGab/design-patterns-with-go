package abstractfactory

type IShoe interface {
	setBrand(brand string)
	setSize(size int)
	getBrand() string
	getSize() int
}


type Shoe struct {
	brand string
	size int
}

func (shoe *Shoe) setBrand(brand string) {
	shoe.brand = brand
}

func (shoe *Shoe) setSize(size int) {
	shoe.size = size
}

func (shoe *Shoe) getBrand() string {
	return shoe.brand
}

func (shoe *Shoe) getSize() int {
	return shoe.size
}
