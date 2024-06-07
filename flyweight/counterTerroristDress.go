package flyweight

type CounterTerroristDress struct {
	color string
}

func (counterTerroristDress *CounterTerroristDress) getColor() string {
	return counterTerroristDress.color
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "green"}
}
