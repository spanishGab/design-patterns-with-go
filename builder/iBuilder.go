package builder

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}
