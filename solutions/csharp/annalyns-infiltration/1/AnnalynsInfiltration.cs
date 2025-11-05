/// <summary>
/// Infiltration logic for Annalyn's quest.
/// </summary>
static class QuestLogic
{
    /// <summary>Attack only if the knight is asleep.</summary>
    public static bool CanFastAttack(bool knightIsAwake) => !knightIsAwake;

    /// <summary>Spy if anyone is awake.</summary>
    public static bool CanSpy(bool knightIsAwake, bool archerIsAwake, bool prisonerIsAwake) =>
        knightIsAwake || archerIsAwake || prisonerIsAwake;

    /// <summary>Signal if the archer sleeps and the prisoner is awake.</summary>
    public static bool CanSignalPrisoner(bool archerIsAwake, bool prisonerIsAwake) =>
        !archerIsAwake && prisonerIsAwake;

    /// <summary>Free prisoner if archer sleeps and (dog helps or knight sleeps and prisoner is awake).</summary>
    public static bool CanFreePrisoner(bool knightIsAwake, bool archerIsAwake, bool prisonerIsAwake, bool petDogIsPresent) =>
        !archerIsAwake && (petDogIsPresent || (prisonerIsAwake && !knightIsAwake));
}
