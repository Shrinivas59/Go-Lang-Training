func TestOverdraft(t *testing.T) {
    a := Account{
        Holder:  "John",
        Balance: 5000,
    }

    err := a.Withdraw(10000)

    if err == nil {
        t.Fatal("Expected overdraft error")
    }

    var overdraft *OverdraftError

    if !errors.As(err, &overdraft) {
        t.Fatal("Expected OverdraftError")
    }

    if a.Balance != 5000 {
        t.Error("Balance should not change")
    }
}