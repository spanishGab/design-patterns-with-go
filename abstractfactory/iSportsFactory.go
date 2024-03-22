package abstractfactory

import "fmt"

type ISportsFactory interface {
    makeShirt() IShirt
    makeShoe() IShoe
}

func CreateSportsFactory(brand string) (ISportsFactory, error) {
    if brand == "adidas" {
        return &AdidasFactory{}, nil
    }

    if brand == "nike" {
        return &NikeFactory{}, nil
    }

    return nil, fmt.Errorf("Wrong brand type passed")
}