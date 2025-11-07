public class Player
{
    private static readonly Random _random = new Random();

    // Generates a number between 1 and 18 inclusive
    public int RollDie() => _random.Next(1, 19);

    // Generates a number in [0.0, 100.0)
    public double GenerateSpellStrength() => _random.NextDouble() * 100.0;
}