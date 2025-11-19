using System;

class BirdCount
{
    // Internal record of birds seen each day (one count per day)
    private int[] birdsPerDay;

    // Store the week's data when creating an instance
    public BirdCount(int[] birdsPerDay)
    {
        this.birdsPerDay = birdsPerDay;
    }

    // Return the fixed data set representing last week's counts
    public static int[] LastWeek() =>
        new[] { 0, 2, 5, 3, 7, 8, 4 };

    // Return today's count (last element in the array)
    public int Today() =>
        birdsPerDay[birdsPerDay.Length - 1];

    // Increase today's count by one
    public void IncrementTodaysCount() =>
        birdsPerDay[birdsPerDay.Length - 1]++;

    // Check if any day had zero birds
    public bool HasDayWithoutBirds() =>
        Array.IndexOf(birdsPerDay, 0) != -1;

    // Sum counts for the first N days
    public int CountForFirstDays(int numberOfDays)
    {
        int sum = 0;

        // Add the values for the requested number of days
        for (int i = 0; i < numberOfDays; i++)
            sum += birdsPerDay[i];

        return sum;
    }

    // Count how many days had 5 or more birds
    public int BusyDays()
    {
        int count = 0;

        foreach (int c in birdsPerDay)
            if (c >= 5)
                count++;

        return count;
    }
}
