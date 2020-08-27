package wallet

import (
	"errors"
	"fmt"
)

//Bitcoin type
type Bitcoin int

// Wallet Struc
type Wallet struct {
	balance Bitcoin
}

//Deposit money in the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//ErrInsufficientFunds error
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

//Withdraw money from the wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

//Balance pf the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance

}

//Stringer interface
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
