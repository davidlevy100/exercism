public static class LogAnalysis
{
    /// <summary>Returns substring after last occurrence of delimiter.</summary>
    public static string SubstringAfter(this string str, string delimiter)
    {
        int index = str.LastIndexOf(delimiter);
        return index == -1 ? string.Empty : str.Substring(index + delimiter.Length);
    }

    /// <summary>Returns substring between last occurrences of start and end.</summary>
    public static string SubstringBetween(this string str, string start, string end)
    {
        int startIndex = str.LastIndexOf(start);
        int endIndex = str.LastIndexOf(end);

        if (startIndex == -1 || endIndex == -1 || startIndex >= endIndex)
            return string.Empty;

        return str.Substring(startIndex + start.Length, endIndex - startIndex - start.Length);
    }

    /// <summary>Extracts message after ": ".</summary>
    public static string Message(this string str) => str.SubstringAfter(": ");

    /// <summary>Extracts text between "[" and "]".</summary>
    public static string LogLevel(this string str) => str.SubstringBetween("[", "]");
}