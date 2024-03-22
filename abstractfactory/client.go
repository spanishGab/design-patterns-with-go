package abstractfactory

import "fmt"

func AbstractFactoryClient() {
	adidasFactory, _ := CreateSportsFactory("adidas")
	nikeFactory, _ := CreateSportsFactory("nike")

	adidasShirt := adidasFactory.makeShirt()
	adidasShoe := adidasFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()
	nikeShoe := nikeFactory.makeShoe()

	printShirtDetails(nikeShirt)
	printShoeDetails(nikeShoe)
	printShirtDetails(adidasShirt)
	printShoeDetails(adidasShoe)
}

func printShoeDetails(shoe IShoe) {
	fmt.Printf("Shoe Brand: %s", shoe.getBrand())
	fmt.Println()
	fmt.Printf("Shoe Size: %d", shoe.getSize())
	fmt.Println()
}

func printShirtDetails(shirt IShoe) {
	fmt.Printf("Shirt Brand: %s", shirt.getBrand())
	fmt.Println()
	fmt.Printf("Shirt Size: %d", shirt.getSize())
	fmt.Println()
}