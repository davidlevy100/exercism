public static class ResistorColor
{
    private static readonly string[] ColorCodes =
    {
        "black", "brown", "red", "orange", "yellow",
        "green", "blue", "violet", "grey", "white"
    };

    public static int ColorCode(string color)
    {
        // case-sensitive by spec, but we can normalize once
        color = color.ToLowerInvariant();

        int index = Array.IndexOf(ColorCodes, color);
        if (index < 0)
            throw new ArgumentException("Invalid color", nameof(color));

        return index;
    }

    public static string[] Colors()
    {
        return ColorCodes;
    }
}
