package adapter

import "fmt"

type Windows struct {
}

func (windows *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
