package bridge

import "fmt"

type Hp struct {
}

func (printer *Hp) PrintFile() {
	fmt.Println("Printing by a HP Printer")
}
