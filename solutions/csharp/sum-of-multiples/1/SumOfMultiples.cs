public static class SumOfMultiples
{
    public static int Sum(IEnumerable<int> multiples, int max)
    {
        var vals = new HashSet<int>();
        foreach (int n in multiples)
        {
            if (n == 0) continue;
            for (int i = n; i < max; i += n)
            {
                vals.Add(i);
            }
        }
        return vals.Sum();
    }
}