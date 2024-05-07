package bridge

import "fmt"

type Mac struct {
	printer IPrinter
}

func (computer *Mac) Print() {
	fmt.Println("Print request for mac")
	computer.printer.PrintFile()
}

func (computer *Mac) SetPrinter(printer IPrinter) {
	computer.printer = printer
}
