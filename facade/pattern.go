package facade

import "fmt"

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func newWalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	walletFacade := &WalletFacade{
		account: newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet: newWallet(),
		notification: &Notification{},
		ledger: &Ledger{},
	}
	fmt.Println("Account created")
	return walletFacade
}

func (walletFacade *WalletFacade) addMoneyToWallet(accountID string, secutiryCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := walletFacade.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = walletFacade.securityCode.checkCode(secutiryCode)
	if err != nil {
		return err
	}
	walletFacade.wallet.creditBalance(amount)
	walletFacade.notification.sendWalletCreditNotification()
	walletFacade.ledger.makeEntry(accountID, "credit", amount)
	return nil
}

func (walletFacade *WalletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := walletFacade.account.checkAccount(accountID)
	if err != nil {
		return nil
	}
	err = walletFacade.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = walletFacade.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	walletFacade.notification.sendWalletDebitNotification()
	walletFacade.ledger.makeEntry(accountID, "debit", amount)
	return nil
}
