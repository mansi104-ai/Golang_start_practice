package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    if balance < 0 {
        return 3.213
    }
    if balance >= 0 && balance < 1000 {
        return 0.5
    }
    if balance >= 1000 && balance < 5000 {
        return 1.621
    }
    return 2.475
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    return (float64(InterestRate(balance)) * balance) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
    return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    years := 0

    // Loop until the current balance is greater than or equal to targetBalance
    for balance < targetBalance {
        balance = AnnualBalanceUpdate(balance)
        years++
    }

    return years
}