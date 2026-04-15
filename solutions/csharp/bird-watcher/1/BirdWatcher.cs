class BirdCount
{
    private int[] birdsPerDay;

    public BirdCount(int[] birdsPerDay) =>
        this.birdsPerDay = birdsPerDay;

    public static int[] LastWeek() => 
        [0, 2, 5, 3, 7, 8, 4];

    public int Today() => 
        this.birdsPerDay[^1];

    public void IncrementTodaysCount() => 
        this.birdsPerDay[^1]++;

    public bool HasDayWithoutBirds() => 
        this.birdsPerDay.Any(b => b == 0);

    public int CountForFirstDays(int numberOfDays) =>
        this.birdsPerDay[..numberOfDays].Sum();

    public int BusyDays() => 
        this.birdsPerDay.Count(b => b >= 5);
}
