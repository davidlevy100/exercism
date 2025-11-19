public class Player
{
    // Shared RNG instance; avoids reseeding issues
    private static readonly Random rand = new Random();

    // Returns a random integer in the range [1, 18] (19 is excluded)
    public int RollDie() => rand.Next(1, 19);

    // Returns a random double in the range [0.0, 100.0)
    public double GenerateSpellStrength() => rand.NextDouble() * 100.0;
}
