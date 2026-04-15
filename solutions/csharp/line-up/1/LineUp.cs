public static class LineUp
{
    public static string Format(string name, int number) => 
        $"{name}, you are the {number}{GetSuffix(number)} customer we serve today. Thank you!";

    private static string GetSuffix(int number)
    {
        int n = Math.Abs(number);
        int lastTwo = n % 100;
    
        if (lastTwo >= 11 && lastTwo <= 13) return "th";
    
        return (n % 10) switch
        {
            1 => "st",
            2 => "nd",
            3 => "rd",
            _ => "th"
        };
    }
}
