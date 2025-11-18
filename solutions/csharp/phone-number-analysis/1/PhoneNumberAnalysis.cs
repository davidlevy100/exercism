public static class PhoneNumber
{
    public static (bool IsNewYork, bool IsFake, string LocalNumber) Analyze(string phoneNumber)
    {
        string[] parts = phoneNumber.Split('-');    // area | exchange | local

        bool isNewYork = parts[0] == "212";         // NYC area code
        bool isFake = parts[1] == "555";            // Hollywood fake exchange
        string local = parts[2];                    // local number

        return (isNewYork, isFake, local);
    }

    public static bool IsFake((bool IsNewYork, bool IsFake, string LocalNumber) info)
        => info.IsFake;                             // return fake flag
}
