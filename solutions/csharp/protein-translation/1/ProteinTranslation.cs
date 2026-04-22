public static class ProteinTranslation
{
    public static string[] Proteins(string strand) =>
        strand.Chunk(3)
            .Select(c => Translate(new string(c)))
            .TakeWhile(p => p != "STOP")
            .ToArray();

    private static string Translate(string codon) => codon switch
    {
        "AUG" => "Methionine",
        "UUU" or "UUC" => "Phenylalanine",
        "UUA" or "UUG" => "Leucine",
        "UCU" or "UCC" or "UCA" or "UCG" => "Serine",
        "UAU" or "UAC" => "Tyrosine",
        "UGU" or "UGC" => "Cysteine",
        "UGG" => "Tryptophan",
        "UAA" or "UAG" or "UGA" => "STOP",
        _ => throw new ArgumentException($"Invalid codon: {codon}")
    };
}