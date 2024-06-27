package chainofresponsibility

type Department interface {
	execute(patient *Patient)
	setNext(department Department)
}
