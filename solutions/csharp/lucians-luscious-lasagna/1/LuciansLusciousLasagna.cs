/// <summary>
/// Calculates preparation and cooking times for a lasagna.
/// </summary>
class Lasagna
{
    /// <summary>Expected baking time in minutes.</summary>
    public int ExpectedMinutesInOven() => 40;

    /// <summary>Minutes remaining in the oven.</summary>
    public int RemainingMinutesInOven(int time) => ExpectedMinutesInOven() - time;

    /// <summary>Prep time is 2 minutes per layer.</summary>
    public int PreparationTimeInMinutes(int layers) => layers * 2;

    /// <summary>Total elapsed time (prep + oven).</summary>
    public int ElapsedTimeInMinutes(int layers, int time) =>
        PreparationTimeInMinutes(layers) + time;
}
