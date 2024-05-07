package bridge

import "fmt"

type Windows struct {
	printer IPrinter
}

func (computer *Windows) Print() {
	fmt.Println("Print request for windows")
	computer.printer.PrintFile()
}

func (computer *Windows) SetPrinter(printer IPrinter) {
	computer.printer = printer
}
