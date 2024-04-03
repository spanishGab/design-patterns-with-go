package builder

type Director struct {
	builder IBuilder
}

func newDirector(builder IBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (director *Director) setBuilder(builder IBuilder) {
	director.builder = builder
}

func (director *Director) buildHouse() House {
	director.builder.setDoorType()
	director.builder.setWindowType()
	director.builder.setNumFloor()
	return director.builder.getHouse()
}