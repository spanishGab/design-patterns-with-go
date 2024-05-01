package adapter

import "fmt"

type WindowsAdapter struct {
	windosMachine *Windows
}

func (adapter *WindowsAdapter) insertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	adapter.windosMachine.insertIntoUSBPort()
}