package builder

import "fmt"

func BuilderClient() {
    woodenBuilder := newWoodenBuilder()
    iglooBuilder := newIglooBuilder()

    director := newDirector(woodenBuilder)
    woodenHouse := director.buildHouse()
    fmt.Printf("Wooden House Door Type: %s\n", woodenHouse.doorType)
    fmt.Printf("Wooden House Window Type: %s\n", woodenHouse.windowType)
    fmt.Printf("Wooden House Num Floor: %d\n", woodenHouse.floor)

    director.setBuilder(iglooBuilder)
    iglooHouse := director.buildHouse()
    fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
    fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
    fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floor)
}