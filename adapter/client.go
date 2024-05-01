package adapter

func AdapterClient() {
	user := &User{}
	mac := &Mac{}

	user.InsertLightningConnectorIntoComputer(mac)

	windosMachine := &Windows{}
	windowsMachineAdapter := &WindowsAdapter{
		windosMachine: windosMachine,
	}
	user.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}