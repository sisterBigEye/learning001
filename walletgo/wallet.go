package walletgo

import "fmt"

type Bitcoin int

type Wallet struct {
  balance Bitcoin
}

type String interface {
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