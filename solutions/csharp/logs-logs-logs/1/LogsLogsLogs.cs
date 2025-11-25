// Represents the numerical encoding for log message severity.
// These values are intentionally non-sequential because the short-log format
// uses fixed numeric codes for each level.
public enum LogLevel
{
    Unknown = 0,
    Trace   = 1,
    Debug   = 2,
    Info    = 4,
    Warning = 5,
    Error   = 6,
    Fatal   = 42
}

// Provides helpers for parsing and compressing log lines.
static class LogLine
{
    // Extracts the log level from a line in the form "[XXX]: message".
    // Assumes the input is well-formed and the code occupies characters 1-3.
    public static LogLevel ParseLogLevel(string logLine)
    {
        string code = logLine.Substring(1, 3);

        switch (code)
        {
            case "TRC": return LogLevel.Trace;
            case "DBG": return LogLevel.Debug;
            case "INF": return LogLevel.Info;
            case "WRN": return LogLevel.Warning;
            case "ERR": return LogLevel.Error;
            case "FTL": return LogLevel.Fatal;
            default:     return LogLevel.Unknown;
        }
    }

    // Produces the compact "<LEVEL>:<MESSAGE>" representation.
    // Only the numeric encoding is used; message is written verbatim.
    public static string OutputForShortLog(LogLevel logLevel, string message) =>
        $"{(int)logLevel}:{message}";
}
