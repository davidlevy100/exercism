public static class PhoneNumber
{
    /// <summary>Returns if the number is from NY, fake, and its local part.</summary>
    public static (bool IsNewYork, bool IsFake, string LocalNumber) Analyze(string phoneNumber)
    {
        var parts = phoneNumber.Split('-');
        return (parts[0] == "212", parts[1] == "555", parts[2]);
    }

    /// <summary>Checks if a phone number is fake.</summary>
    public static bool IsFake((bool IsNewYork, bool IsFake, string LocalNumber) info)
        => info.IsFake;
}
