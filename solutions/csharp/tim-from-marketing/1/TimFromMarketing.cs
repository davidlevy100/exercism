public static class Badge
{
    /// <summary>Formats a badge with ID, name, and department.</summary>
    public static string Print(int? id, string name, string? department)
    {
        string idPart = id is null ? "" : $"[{id}] - ";
        string deptPart = department?.ToUpper() ?? "OWNER";
        return $"{idPart}{name} - {deptPart}";
    }
}
