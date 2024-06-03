package facade

import "fmt"

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{
		balance: 0,
	}
}

func (wallet *Wallet) creditBalance(amount int) {
	wallet.balance += amount
	fmt.Println("Wallet balance added successfuly")
	return
}

func (wallet *Wallet) debitBalance(amount int) error {
	if wallet.balance < amount {
		return fmt.Errorf("Balance is not sufficient")
	}
	fmt.Println("Wallet balance is Sufficient")
	wallet.balance -= amount
	return nil
}