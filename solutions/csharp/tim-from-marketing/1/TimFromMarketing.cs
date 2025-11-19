static class Badge
{
    public static string Print(int? id, string name, string? department)
    {
        // Normalize department: uppercase if present, fallback to "OWNER" if null
        var dept = department?.ToUpper() ?? "OWNER";

        // If id is null, omit the ID section; otherwise include [id]
        return id == null
            ? $"{name} - {dept}"
            : $"[{id}] - {name} - {dept}";
    }
}
