package walletgo

import (
	"testing"
)

func TestWallet(t *testing.T)  {

	assertBalance := func(t *testing.T,wallet Wallet,want Bitcoin){
		got := wallet.Balance()
		if got != want{
			t.Errorf("got %s want %s",got,want)
		}
	}

	assertError := func(t *testing.T,got error,want error){
		if got == nil {
			t.Fatal("wanted an error but didnt get one")
		}
		if got!= want{
			t.Errorf("got '%s',want '%s'",got,want)
		}
	}



    t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		assertBalance(t,wallet,Bitcoin(10))

	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t,wallet,Bitcoin(10))
		assertNoError(t,err)
	})

	t.Run("Withdraw insufficient funds",func(t *testing.T){

		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t,wallet,Bitcoin(20))
		assertError(t,err,InsufficientFundsError)
	})

}

func assertNoError(t *testing.T,got error){
	if got !=nil{
		t.Fatal("got an error but didnt want one")
	}
}