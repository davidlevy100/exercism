public static class LogAnalysis 
{
    // Returns everything after the first occurrence of delimiter
    public static string SubstringAfter(this string str, string delimiter) => 
        str.Split(delimiter)[1];

    // Get everything after start, then cut at end
    public static string SubstringBetween(this string str, string start, string end) => 
        str.SubstringAfter(start).Split(end)[0];

    // Message follows ": "; Trim handles edge-case whitespace
    public static string Message(this string str) =>
        str.SubstringAfter(": ").Trim();

    // Level sits between "[" and "]"
    public static string LogLevel(this string str) =>
        str.SubstringBetween("[", "]");
}