package chainofresponsibility

import "fmt"

type Medical struct {
	next Department
}

func (medical *Medical) execute(patient *Patient) {
	if patient.medicineDone {
		fmt.Println("Medicine already given to patient")
		medical.next.execute(patient)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	patient.medicineDone = true
	medical.next.execute(patient)
}

func (medical *Medical) setNext(next Department) {
	medical.next = next
}
