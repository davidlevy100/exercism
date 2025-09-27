// Package interest provides functions for interest rate calculations.
package interest

// InterestRate returns the annual interest rate (in percent)
// for the given balance.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return 3.213
	case balance < 1000:
		return 0.5
	case balance < 5000:
		return 1.621
	default:
		return 2.475
	}
}

// Interest returns the annual interest earned on the balance.
func Interest(balance float64) float64 {
	return float64(InterestRate(balance)) / 100 * balance
}

// AnnualBalanceUpdate returns the balance after one year of interest.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance returns the number of years required
// for balance to reach or exceed targetBalance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	years := 0
	for balance < targetBalance {
		balance = AnnualBalanceUpdate(balance)
		years++
	}
	return years
}
