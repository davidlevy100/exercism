public static class Languages
{
    // Creates a new, empty language list.
    public static List<string> NewList() =>
        new List<string>();

    // Returns a preset list of existing languages.
    public static List<string> GetExistingLanguages() =>
        new List<string> { "C#", "Clojure", "Elm" };

    // Adds a language to the list and returns the same list (mutating in place).
    public static List<string> AddLanguage(List<string> languages, string language)
    {
        languages.Add(language);
        return languages;
    }

    // Returns the number of languages in the list.
    public static int CountLanguages(List<string> languages) =>
        languages.Count;

    // Checks whether a specific language exists in the list.
    public static bool HasLanguage(List<string> languages, string language) =>
        languages.Contains(language);

    // Reverses the list in place and returns it.
    public static List<string> ReverseList(List<string> languages)
    {
        languages.Reverse();
        return languages;
    }

    // The "exciting" rule:
    // - If the first item is "C#", it's exciting.
    // - If the second item is "C#" AND the list length is 2 or 3, it's exciting.
    public static bool IsExciting(List<string> languages)
    {
        // Empty list can't be exciting.
        if (languages.Count == 0)
            return false;

        // Rule 1: C# is first.
        if (languages[0] == "C#")
            return true;

        // Rule 2: C# is second, and list size is 2 or 3.
        if (languages.Count >= 2 &&
            languages[1] == "C#" &&
            (languages.Count == 2 || languages.Count == 3))
        {
            return true;
        }

        return false;
    }

    // Removes a language from the list if it exists, returning the mutated list.
    public static List<string> RemoveLanguage(List<string> languages, string language)
    {
        languages.Remove(language);
        return languages;
    }

    // Returns true if all items are unique.
    // HashSet automatically removes duplicates, so comparing size is enough.
    public static bool IsUnique(List<string> languages)
    {
        HashSet<string> unique = new HashSet<string>(languages);
        return languages.Count == unique.Count;
    }
}
