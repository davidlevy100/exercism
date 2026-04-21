public static class Leap
{
    public static bool IsLeapYear(int year) => year switch
    {
        _ when year % 400 == 0 => true,
        _ when year % 100 == 0 => false,
        _ when year % 4   == 0 => true,
        _                      => false,
    };
}