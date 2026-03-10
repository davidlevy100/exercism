public class Player {

    // [1, 18] inclusive
    public int RollDie() => Random.Shared.Next(1, 19);

    // [0.0, 100.0)
    public double GenerateSpellStrength() => Random.Shared.NextDouble() * 100.0;
}