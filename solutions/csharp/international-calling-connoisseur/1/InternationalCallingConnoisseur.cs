public static class DialingCodes
{
    // Factory: empty dictionary
    public static Dictionary<int, string> GetEmptyDictionary() => new();

    // Factory: pre-populated with sample data
    public static Dictionary<int, string> GetExistingDictionary() => new()
    {
        [1]  = "United States of America",
        [55] = "Brazil",
        [91] = "India"
    };

    // Upsert primitive: adds or overwrites regardless of key existence
    public static Dictionary<int, string> AddCountryToExistingDictionary(
        Dictionary<int, string> existingDictionary, int countryCode, string countryName)
    {
        existingDictionary[countryCode] = countryName;
        return existingDictionary;
    }

    // Composes GetEmptyDictionary + AddCountryToExistingDictionary
    public static Dictionary<int, string> AddCountryToEmptyDictionary(int countryCode, string countryName) =>
        AddCountryToExistingDictionary(GetEmptyDictionary(), countryCode, countryName);

    // Update only: no-op if key doesn't exist
    public static Dictionary<int, string> UpdateDictionary(
        Dictionary<int, string> existingDictionary, int countryCode, string countryName)
    {
        if (existingDictionary.ContainsKey(countryCode))
            existingDictionary[countryCode] = countryName;
        return existingDictionary;
    }

    // Removes entry if present; no-op if missing
    public static Dictionary<int, string> RemoveCountryFromDictionary(
        Dictionary<int, string> existingDictionary, int countryCode)
    {
        existingDictionary.Remove(countryCode);
        return existingDictionary;
    }

    // Returns empty string if key not found
    public static string GetCountryNameFromDictionary(
        Dictionary<int, string> existingDictionary, int countryCode) =>
        existingDictionary.TryGetValue(countryCode, out var name) ? name : string.Empty;

    // Key existence check
    public static bool CheckCodeExists(Dictionary<int, string> existingDictionary, int countryCode) =>
        existingDictionary.ContainsKey(countryCode);

    // Returns longest country name by character count; empty string if dictionary is empty
    public static string FindLongestCountryName(Dictionary<int, string> existingDictionary) =>
        existingDictionary.Values.MaxBy(name => name.Length) ?? string.Empty;
}