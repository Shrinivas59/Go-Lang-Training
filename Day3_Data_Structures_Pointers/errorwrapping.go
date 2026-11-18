var ErrInsufficientBalance = errors.New("insufficient balance")

func Withdraw(amount float64) error {
    if amount > 5000 {
        return fmt.Errorf(
            "withdrawal failed: %w",
            ErrInsufficientBalance,
        )
    }

    return nil
}