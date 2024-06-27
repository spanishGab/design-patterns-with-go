package chainofresponsibility

import "fmt"

type Cashier struct {
	next Department
}

func (cashier *Cashier) execute(patient *Patient) {
	if patient.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient")
	patient.paymentDone = true
}

func (cashier *Cashier) setNext(next Department) {
	cashier.next = next
}
