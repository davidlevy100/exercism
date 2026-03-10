public class Player
{
    // Shared RNG instance; avoids reseeding issues
    private static readonly Random rand = new();

    // [1, 18] inclusive
    public int RollDie() => rand.Next(1, 19);

    // [0.0, 100.0)
    public double GenerateSpellStrength() => rand.NextDouble() * 100.0;
}