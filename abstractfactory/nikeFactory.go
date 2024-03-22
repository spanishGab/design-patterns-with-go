package abstractfactory

type NikeShirt struct {
	Shirt
}

type NikeShoe struct {
	Shoe
}

type NikeFactory struct {
}

func (nikeFactory *NikeFactory) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			brand: "nike",
			size: 14,
		},
	}
}

func (nikeFactory *NikeFactory) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			brand: "nike",
			size: 41,
		},
	}
}
