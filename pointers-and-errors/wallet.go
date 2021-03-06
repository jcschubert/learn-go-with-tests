package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin represents an amount of bitcoin.
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// ErrInsufficientFunds that is returned if balance is insufficient for a withdrawal from a wallet
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Wallet represents a Bitcoin wallet.
type Wallet struct {
	balance Bitcoin
}

// Deposit adds the provided amount to the Wallet's balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance shows us the wallet's balance
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw withdraws the shown amount from the Wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Stringer allows us to convert a type to a string
type Stringer interface {
	String() string
}
