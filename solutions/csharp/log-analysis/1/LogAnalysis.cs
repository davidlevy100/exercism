public static class LogAnalysis
{
    // Return substring after last occurrence of a delimiter string
    public static string SubstringAfter(this string s, string delimiter)
    {
        int idx = s.LastIndexOf(delimiter);
        if (idx < 0) return "";

        return s.Substring(idx + delimiter.Length);
    }

    // Return substring between two delimiters
    public static string SubstringBetween(this string s, string start, string end)
    {
        int i1 = s.IndexOf(start);
        if (i1 < 0) return "";

        i1 += start.Length;

        int i2 = s.IndexOf(end, i1);
        if (i2 < 0) return "";

        return s.Substring(i1, i2 - i1);
    }

    // Extract message portion
    public static string Message(this string s)
        => s.SubstringAfter(":").Trim();

    // Extract log level portion
    public static string LogLevel(this string s)
        => s.SubstringBetween("[", "]").Trim();
}
