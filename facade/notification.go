package facade

import "fmt"

type Notification struct {
}

func (notification *Notification) sendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (notification *Notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}
