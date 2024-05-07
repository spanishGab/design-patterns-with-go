package bridge

import "fmt"

type Epson struct {
}

func (printer *Epson) PrintFile() {
	fmt.Println("Printing by a EPSON Printer")
}
