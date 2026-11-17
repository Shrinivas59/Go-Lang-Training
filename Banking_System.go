
package main

import (
    "errors"
    "fmt"
)

type Account struct {
    AccountHolder string
    AccountNumber int
    Balance       float64
}

// Deposit money
func (a *Account) Deposit(amount float64) error {
    if amount <= 0 {
        return errors.New("deposit amount must be positive")
    }

    a.Balance += amount
    return nil
}

// Withdraw money
func (a *Account) Withdraw(amount float64) error {
    if amount <= 0 {
        return errors.New("withdrawal amount must be positive")
    }

    if amount > a.Balance {
        return errors.New("insufficient balance")
    }

    a.Balance -= amount
    return nil
}

// Check balance
func (a Account) CheckBalance() float64 {
    return a.Balance
}

// Display account details
func (a Account) String() string {
    return fmt.Sprintf(
        "Account Holder: %s | Account Number: %d | Balance: %.2f",
        a.AccountHolder,
        a.AccountNumber,
        a.Balance,
    )
}

// Banking interface
type Banking interface {
    Deposit(float64) error
    Withdraw(float64) error
    CheckBalance() float64
}

func main() {

    account := Account{
        AccountHolder: "John",
        AccountNumber: 1001,
        Balance:       5000,
    }

    // Interface implementation
    var bank Banking = &account

    fmt.Println("Initial Account Details:")
    fmt.Println(account)

    // Deposit
    err := bank.Deposit(2000)
    if err != nil {
        fmt.Println("Deposit Error:", err)
    } else {
        fmt.Println("Deposit successful")
    }

    // Withdraw
    err = bank.Withdraw(1500)
    if err != nil {
        fmt.Println("Withdrawal Error:", err)
    } else {
        fmt.Println("Withdrawal successful")
    }

    // Check balance
    fmt.Println("Current Balance:", bank.CheckBalance())

    // Display final details
    fmt.Println("Final Account Details:")
    fmt.Println(account)
}
