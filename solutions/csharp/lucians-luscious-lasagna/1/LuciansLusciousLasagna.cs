class Lasagna
{
    // Total expected oven time (in minutes)
    private const int OvenTime = 40;

    // Time per layer (in minutes)
    private const int LayerTime = 2;

    // Returns the total expected oven time
    public int ExpectedMinutesInOven() => OvenTime;

    // Returns remaining minutes based on time in oven so far
    public int RemainingMinutesInOven(int minutesIn) =>
        OvenTime - minutesIn;

    // Returns total prep time based on number of layers
    public int PreparationTimeInMinutes(int layers) =>
        layers * LayerTime;

    // Returns total elapsed time: prep + bake time so far
    public int ElapsedTimeInMinutes(int layers, int minutesIn) =>
        PreparationTimeInMinutes(layers) + minutesIn;
}
