static class AssemblyLine
{
    public static double SuccessRate(int speed)
    {
        if (speed <= 0) return 0.0;   // invalid or zero speed → no production
        if (speed <= 4) return 1.0;   // speeds 1–4 → perfect
        if (speed <= 8) return 0.9;   // speeds 5–8 → minor errors
        if (speed == 9) return 0.8;   // speed 9 → noticeable errors
        return 0.77;                  // 10 → messy
    }

    // cars per hour with error rate factored in
    public static double ProductionRatePerHour(int speed) =>
        speed * 221.0 * SuccessRate(speed);

    // whole cars per minute
    public static int WorkingItemsPerMinute(int speed) =>
        (int)(ProductionRatePerHour(speed) / 60.0);
}