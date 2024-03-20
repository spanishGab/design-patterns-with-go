package abstractfactory

type ISportsFactory interface {
    makeShirt(brand string, size int) IShirt
    makeShoe(brand string, size int) IShoe
}
