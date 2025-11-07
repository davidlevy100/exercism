static class SavingsAccount
{
    // Returns the applicable interest rate (%) based on balance tiers
    public static float InterestRate(decimal balance)
    {
        if (balance < 0m)
            return 3.213f;
        else if (balance < 1000m)
            return 0.5f;
        else if (balance < 5000m)
            return 1.621f;
        else
            return 2.475f;
    }

    // Calculates annual interest amount based on current balance
    public static decimal Interest(decimal balance) =>
        balance * (decimal)InterestRate(balance) / 100m;

    // Returns the updated balance after applying annual interest
    public static decimal AnnualBalanceUpdate(decimal balance) =>
        balance + Interest(balance);

    // Calculates years required to reach or exceed the target balance
    public static int YearsBeforeDesiredBalance(decimal balance, decimal targetBalance)
    {
        int years = 0;
        while (balance < targetBalance)
        {
            balance = AnnualBalanceUpdate(balance);
            years++;
        }
        return years;
    }
}
