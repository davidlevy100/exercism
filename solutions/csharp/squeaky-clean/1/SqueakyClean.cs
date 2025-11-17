using System.Text;

public static class Identifier
{
    public static string Clean(string identifier)
    {
        if (identifier == null)
            return null;

        var sb = new StringBuilder();
        bool kebabFlag = false;

        foreach (char c in identifier)
        {
            // Convert spaces to underscores
            if (c == ' ')
            {
                sb.Append('_');
                continue;
            }

            // Replace control characters with literal "CTRL"
            if (char.IsControl(c))
            {
                sb.Append("CTRL");
                continue;
            }

            // A hyphen signals we should capitalize the next character
            if (c == '-')
            {
                kebabFlag = true;
                continue;
            }

            // Uppercase the character following a hyphen
            if (kebabFlag)
            {
                sb.Append(char.ToUpperInvariant(c));
                kebabFlag = false;
                continue;
            }

            // Skip Greek lowercase alpha through omega (per assignment spec)
            if (c >= 'α' && c <= 'ω')
                continue;

            // Only letters are allowed past this point
            if (!char.IsLetter(c))
                continue;

            // Append normal letter
            sb.Append(c);
        }

        return sb.ToString();
    }
}
