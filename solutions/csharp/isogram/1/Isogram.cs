public static class Isogram
{
    public static bool IsIsogram(string word)
    {
        var seen = new HashSet<char>();
        foreach (char c in word)
        {
            if (Char.IsLetter(c) && !seen.Add(Char.ToLower(c)))
                return false;
        }
        return true;
    }
}