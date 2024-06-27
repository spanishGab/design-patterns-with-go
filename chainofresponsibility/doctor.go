package chainofresponsibility

import "fmt"

type Doctor struct {
	next Department
}

func (doctor *Doctor) execute(patient *Patient) {
	if patient.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		doctor.next.execute(patient)
		return
	}
	fmt.Println("Doctor checking patient")
	patient.doctorCheckUpDone = true
	doctor.next.execute(patient)
}

func (doctor *Doctor) setNext(next Department) {
	doctor.next = next
}
