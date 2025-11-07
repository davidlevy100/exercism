using System;
using System.Linq;

class BirdCount
{
    private int[] birdsPerDay;

    public BirdCount(int[] birdsPerDay) => this.birdsPerDay = birdsPerDay;

    /// <summary>Returns last week's sample data.</summary>
    public static int[] LastWeek() => new[] { 0, 2, 5, 3, 7, 8, 4 };

    /// <summary>Returns today's bird count.</summary>
    public int Today() => birdsPerDay[birdsPerDay.Length - 1];

    /// <summary>Increments today's count by one.</summary>
    public void IncrementTodaysCount() => birdsPerDay[birdsPerDay.Length - 1]++;

    /// <summary>Checks if any day had zero birds.</summary>
    public bool HasDayWithoutBirds() => birdsPerDay.Any(c => c == 0);

    /// <summary>Returns total birds for the first N days.</summary>
    public int CountForFirstDays(int numberOfDays) => birdsPerDay.Take(numberOfDays).Sum();

    /// <summary>Counts the number of busy days (5+ birds).</summary>
    public int BusyDays() => birdsPerDay.Count(c => c >= 5);
}
