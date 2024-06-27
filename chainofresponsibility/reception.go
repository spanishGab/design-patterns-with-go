package chainofresponsibility

import "fmt"

type Reception struct {
	next Department
}

func (reception *Reception) execute(patient *Patient) {
	if patient.registrationDone {
		fmt.Println("Patient registration already done")
		reception.next.execute(patient)
		return
	}
	fmt.Println("Reception registering patient")
	patient.registrationDone = true
	reception.next.execute(patient)
}

func (reception *Reception) setNext(next Department) {
	reception.next = next
}
