err := account.Withdraw(10000)

if err != nil {
    fmt.Println("Transaction failed:", err)
} else {
    fmt.Println("Withdrawal successful")
}