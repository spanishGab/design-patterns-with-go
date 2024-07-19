package command

import "fmt"

type Tv struct {
	isRunning bool
}

func (tv *Tv) on() {
	tv.isRunning = true
	fmt.Println("Turning tv on")
}

func (tv *Tv) off() {
	tv.isRunning = false
	fmt.Println("Turning tv off")
}
