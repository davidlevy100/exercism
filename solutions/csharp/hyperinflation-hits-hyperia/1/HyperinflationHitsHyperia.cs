public static class CentralBank
{
    public static string DisplayDenomination(long a, long b)
    {
        try
        {
            checked { return (a * b).ToString(); }
        }
        catch (OverflowException)
        {
            return "*** Too Big ***";
        }
    }

    public static string DisplayGDP(float a, float b)
    {
        float result = a * b;

        if (float.IsNaN(result) || float.IsInfinity(result))
            return "*** Too Big ***";

        return result.ToString();
    }

    public static string DisplayChiefEconomistSalary(decimal a, decimal b)
    {
        try
        {
            checked { return (a * b).ToString(); }
        }
        catch (OverflowException)
        {
            return "*** Much Too Big ***";
        }
    }
}
