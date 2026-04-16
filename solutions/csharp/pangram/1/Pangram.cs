public static class Pangram
{
    public static bool IsPangram(string input)
    {
        var letters = new HashSet<char>();
        foreach (char c in input)
        {
            if (c >= 'a' && c <= 'z') letters.Add(c);
            else if (c >= 'A' && c <= 'Z') letters.Add((char)(c + 32));
        }
        return letters.Count == 26;
    }
}