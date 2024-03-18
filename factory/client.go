package factory

import "fmt"

func FactoryClient() {
	ak47, _ := createGun("ak47")
	musket, _ := createGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
