package flyweight

type TerroristDress struct {
	color string
}

func (terroristDress *TerroristDress) getColor() string {
	return terroristDress.color
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}
