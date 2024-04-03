package builder

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (builder *IglooBuilder) setWindowType() {
	builder.windowType = "Snow Window"
}

func (builder *IglooBuilder) setDoorType() {
	builder.doorType = "Snow Door"
}

func (builder *IglooBuilder) setNumFloor() {
	builder.floor = 1
}

func (builder *IglooBuilder) getHouse() House {
	return House{
		windowType: builder.windowType,
		doorType: builder.doorType,
		floor: builder.floor,
	}
}
