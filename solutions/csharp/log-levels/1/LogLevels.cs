static class LogLine
{
    // Split on ": " to avoid ambiguity; Trim handles edge-case whitespace
    public static string Message(string logLine) => 
        logLine.Split(": ")[1].Trim();
    
    // Split on both brackets — "[ERROR]: ..." → {"", "ERROR", ": ..."}
    public static string LogLevel(string logLine) => 
        logLine.Split('[', ']')[1].ToLower();
    
    // Compose from the two methods above
    public static string Reformat(string logLine) => 
        $"{Message(logLine)} ({LogLevel(logLine)})";
}