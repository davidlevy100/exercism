public static class DialingCodes
{
    // Returns a new, empty dictionary.
    public static Dictionary<int, string> GetEmptyDictionary() =>
        new Dictionary<int, string>();

    // Returns a dictionary preset with a few country codes.
    public static Dictionary<int, string> GetExistingDictionary() =>
        new Dictionary<int, string>
        {
            { 1,  "United States of America" },
            { 55, "Brazil" },
            { 91, "India" }
        };

    // Adds a single entry to a new empty dictionary.
    public static Dictionary<int, string> AddCountryToEmptyDictionary(
        int countryCode, string countryName)
    {
        var result = GetEmptyDictionary();
        result.Add(countryCode, countryName);
        return result;
    }

    // Adds a new entry to an existing dictionary.
    public static Dictionary<int, string> AddCountryToExistingDictionary(
        Dictionary<int, string> existingDictionary, 
        int countryCode, 
        string countryName)
    {
        existingDictionary.Add(countryCode, countryName);
        return existingDictionary;
    }

    // Returns the country name if present; otherwise an empty string.
    public static string GetCountryNameFromDictionary(
        Dictionary<int, string> existingDictionary, int countryCode)
        => CheckCodeExists(existingDictionary, countryCode)
            ? existingDictionary[countryCode]
            : string.Empty;

    // Checks whether the dictionary contains a given key.
    public static bool CheckCodeExists(
        Dictionary<int, string> existingDictionary, int countryCode)
        => existingDictionary.ContainsKey(countryCode);

    // Replaces the value only if the key already exists.
    public static Dictionary<int, string> UpdateDictionary(
        Dictionary<int, string> existingDictionary, 
        int countryCode, 
        string countryName)
    {
        if (CheckCodeExists(existingDictionary, countryCode))
        {
            existingDictionary[countryCode] = countryName;
        }

        return existingDictionary;
    }

    // Removes a key only if it already exists.
    public static Dictionary<int, string> RemoveCountryFromDictionary(
        Dictionary<int, string> existingDictionary, int countryCode)
    {
        if (CheckCodeExists(existingDictionary, countryCode))
        {
            existingDictionary.Remove(countryCode);
        }

        return existingDictionary;
    }

    // Finds and returns the longest country name.
    public static string FindLongestCountryName(
        Dictionary<int, string> existingDictionary)
    {
        string longestName = "";
        int longestLength = 0;

        foreach (string name in existingDictionary.Values)
        {
            if (name.Length > longestLength)
            {
                longestLength = name.Length;
                longestName = name;
            }
        }

        return longestName;
    }
}
