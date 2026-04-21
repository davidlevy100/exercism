using System;
using System.Text;

public static class RomanNumeralExtension
{
    private static readonly int[] Values  = { 1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1 };
    private static readonly string[] Symbols = { "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I" };

    public static string ToRoman(this int value)
    {
        if (value < 1 || value > 3999)
            throw new ArgumentOutOfRangeException(nameof(value), "Roman numerals support 1–3999.");

        var sb = new StringBuilder();
        for (int i = 0; i < Values.Length; i++)
        {
            while (value >= Values[i])
            {
                sb.Append(Symbols[i]);
                value -= Values[i];
            }
        }
        return sb.ToString();
    }
}