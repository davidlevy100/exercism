static class LogLine
{
    // Extract the message after the last colon
    public static string Message(string logLine)
    {
        int idx = logLine.LastIndexOf(':');
        return logLine.Substring(idx + 1).Trim();
    }

    // Extract the log level between brackets, lowercased
    public static string LogLevel(string logLine)
    {
        int idx = logLine.LastIndexOf(':');
        string level = logLine.Substring(0, idx);
        return level.Trim('[', ']', ' ').ToLowerInvariant();
    }

    // Reformat using the two helper functions
    public static string Reformat(string logLine)
        => $"{Message(logLine)} ({LogLevel(logLine)})";
}
