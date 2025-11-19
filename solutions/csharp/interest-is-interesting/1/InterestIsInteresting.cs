static class SavingsAccount
{
    // Determine interest rate based on balance tiers
    public static float InterestRate(decimal balance) =>
        balance < 0m    ? 3.213f :
        balance < 1000m ? 0.5f   :
        balance < 5000m ? 1.621f :
                          2.475f;

    // Yearly interest amount
    public static decimal Interest(decimal balance) =>
        balance * (decimal)InterestRate(balance) / 100m;

    // Apply one year's interest
    public static decimal AnnualBalanceUpdate(decimal balance) =>
        balance + Interest(balance);

    // Number of years to reach the target
    public static int YearsBeforeDesiredBalance(decimal balance, decimal targetBalance)
    {
        int years = 0;

        while (balance < targetBalance) {
            balance = AnnualBalanceUpdate(balance);
            years++;
        }

        return years;
    }
}
