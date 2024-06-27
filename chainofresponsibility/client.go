package chainofresponsibility

func CoRClient() {
	reception := &Reception{}
	doctor := &Doctor{}
	medical := &Medical{}
	cashier := &Cashier{}
	
	reception.setNext(doctor)
	doctor.setNext(medical)
	medical.setNext(cashier)

	patient := &Patient{name: "Gabriel"}
	reception.execute(patient)
}