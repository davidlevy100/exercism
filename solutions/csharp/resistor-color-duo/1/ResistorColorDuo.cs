public static class ResistorColorDuo
{
    private static readonly string[] _colors = 
    ["black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"];

    public static int Value(string[] colors) =>
        Array.IndexOf(_colors, colors[0]) * 10 + Array.IndexOf(_colors, colors[1]);
}
