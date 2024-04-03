package builder

type WoodenBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newWoodenBuilder() *WoodenBuilder {
	return &WoodenBuilder{}
}

func (builder *WoodenBuilder) setWindowType() {
	builder.windowType = "Wooden Window"
}

func (builder *WoodenBuilder) setDoorType() {
	builder.doorType = "Wooden Door"
}

func (builder *WoodenBuilder) setNumFloor() {
	builder.floor = 2
}


func (builder *WoodenBuilder) getHouse() House {
	return House{
		windowType: builder.windowType,
		doorType:   builder.doorType,
		floor:      builder.floor,
	}
}
