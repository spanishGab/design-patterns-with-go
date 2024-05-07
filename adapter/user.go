package adapter

import "fmt"

type User struct {
}

func (user *User) InsertLightningConnectorIntoComputer(computer IComputer) {
	fmt.Println("User inserts Lightning connector into computer.")
	computer.insertIntoLightningPort()
}
