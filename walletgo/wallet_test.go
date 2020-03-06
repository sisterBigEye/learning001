package walletgo

import (
	"testing"
)

func TestWallet(t *testing.T)  {

	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))
	got := wallet.Balance()
	want := Bitcoin(20)

	if got != want {
		t.Errorf("got %d want %s",got,want)
	}

}
