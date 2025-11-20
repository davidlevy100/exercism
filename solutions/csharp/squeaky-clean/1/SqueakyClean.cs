using System;
using System.Text;

public static class Identifier
{
    public static string Clean(string identifier)
    {
        var sb = new StringBuilder();
        bool capitalizeNext = false;

        foreach (char c in identifier)
        {
            // 1. Space → underscore
            if (c == ' ')
            {
                sb.Append('_');
                continue;
            }

            // 2. Control character → "CTRL"
            if (char.IsControl(c))
            {
                sb.Append("CTRL");
                continue;
            }

            // 3. Kebab case: '-' triggers camelCase
            if (c == '-')
            {
                capitalizeNext = true;
                continue;
            }

            // 4. Omit non-letters
            if (!char.IsLetter(c))
            {
                // underscore is only allowed if produced already from spaces
                continue;
            }

            // 5. Omit Greek lowercase letters (α to ω)
            if (c >= 'α' && c <= 'ω')
                continue;

            // Apply camelCase capitalization
            char final = capitalizeNext ? char.ToUpper(c) : c;
            capitalizeNext = false;

            sb.Append(final);
        }

        return sb.ToString();
    }
}
