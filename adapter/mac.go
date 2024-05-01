package adapter

import "fmt"

type Mac struct {
}

func (mac *Mac) insertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}
