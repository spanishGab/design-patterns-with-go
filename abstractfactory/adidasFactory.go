package abstractfactory

type AdidasShirt struct {
	Shirt
}

type AdidasShoe struct {
	Shoe
}

type AdidasFactory struct {
}

func (adidasFactory *AdidasFactory) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			brand: "adidas",
			size: 14,
		},
	}
}

func (adidasFactory *AdidasFactory) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			brand: "adidas",
			size: 41,
		},
	}
}
