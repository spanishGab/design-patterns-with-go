package bridge

type IComputer interface {
	Print()
	SetPrinter(IPrinter)
}