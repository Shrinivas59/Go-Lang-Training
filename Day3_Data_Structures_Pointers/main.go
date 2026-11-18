package main

import (
    "fmt"
    "banking-app/account"
)

func main() {
    a := account.Account{
        Holder:  "John",
        Balance: 5000,
    }

    a.Deposit(1000)

    fmt.Println(a.Balance)
}