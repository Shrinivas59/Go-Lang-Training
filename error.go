import (
    "errors"
    "fmt"
)

func (a *Account) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("amount must be positive")
    }

    if amount > a.Balance {
        return errors.New("insufficient balance")
    }

    a.Balance -= amount

    return nil
}