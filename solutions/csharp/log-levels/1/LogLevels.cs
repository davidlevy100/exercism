/// <summary>
/// Parses and reformats simple log lines of the form "[LEVEL]: message".
/// </summary>
static class LogLine
{
    /// <summary>Extracts the message text from a log line.</summary>
    public static string Message(string logLine) =>
        logLine.Split(':')[1].Trim();

    /// <summary>Extracts the log level (lowercase) from a log line.</summary>
    public static string LogLevel(string logLine) =>
        logLine.Split(':')[0].Trim('[', ']', ' ').ToLowerInvariant();

    /// <summary>Returns "message (level)" format.</summary>
    public static string Reformat(string logLine) =>
        $"{Message(logLine)} ({LogLevel(logLine)})";
}
