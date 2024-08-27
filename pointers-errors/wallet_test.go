package wallet

import (
	"testing"
)

// ========= v1. =========
// func TestWallet(t *testing.T) {
// 	t.Run("Deposit", func(t *testing.T) {

// 		wallet := Wallet{}

// 		wallet.Deposit(Bitcoin(10))

// 		got := wallet.Balance()

// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s, want %s", got, want)
// 		}
// 	})

// 	t.Run("Withdraw", func(t *testing.T) {
// 		wallet := Wallet{balance: Bitcoin(100)}
// 		wallet.Withdraw(Bitcoin(50))

// 		got := wallet.Balance()
// 		want := Bitcoin(50)

// 		if got != want {
// 			t.Errorf("Got %s, want %s", got, want)
// 		}
// 	})
// }

// ========= v2. =========

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		// t.Fatal will stop the test when called.
		t.Fatal("didn't get an error but wanted one")
	}
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 100}
		err := wallet.Withdraw(Bitcoin(40))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(60))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		initialBalance := Bitcoin(50)
		wallet := Wallet{balance: initialBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, initialBalance)
	})
}
