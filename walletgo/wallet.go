package walletgo

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
  balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance +=  amount
}

func (b Bitcoin) String() string{
	return fmt.Sprintf("%d BTC",b)
}


var InsufficientFundsError = errors.New("cannot withdraw,insuffcient funds")
func (w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance{
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}