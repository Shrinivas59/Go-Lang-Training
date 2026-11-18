package account

import "testing"

func TestWithdraw(t *testing.T) {
    a := Account{
        Holder:  "John",
        Balance: 5000,
    }

    err := a.Withdraw(1000)

    if err != nil {
        t.Fatal("Unexpected error:", err)
    }

    if a.Balance != 4000 {
        t.Errorf(
            "Expected balance 4000, got %.2f",
            a.Balance,
        )
    }
}