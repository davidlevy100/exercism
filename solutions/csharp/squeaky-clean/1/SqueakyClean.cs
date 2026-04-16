using System.Text;

public static class Identifier
{
    public static string Clean(string identifier)
    {
        var parts = identifier.Split('-').Select(CleanPart);
        return string.Concat(parts.Select((p, i) => i == 0 ? p : Capitalize(p)));
    }

    private static string CleanPart(string s)
    {
        var sb = new StringBuilder();
        foreach (char c in s)
        {
            sb.Append(c switch
            {
                _ when char.IsWhiteSpace(c) => "_",
                _ when char.IsControl(c)    => "CTRL",
                _ when c >= 'α' && c <= 'ω' => "",
                _ when char.IsLetter(c)     => c.ToString(),
                _                           => ""
            });
        }
        return sb.ToString();
    }

    private static string Capitalize(string s) =>
        s.Length == 0 ? s : char.ToUpper(s[0]) + s[1..];
}