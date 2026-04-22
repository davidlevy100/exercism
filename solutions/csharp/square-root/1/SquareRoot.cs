public static class SquareRoot
{
    public static int Root(int number)
    {
        if (number == 0) return 0;

        int x = number;
        while (true)
        {
            int next = (x + number / x) / 2;
            if (next >= x) return x;
            x = next;
        }
    }
}