static class QuestLogic
{
    // Fast attack only works if the knight is out cold
    public static bool CanFastAttack(bool knightIsAwake) => !knightIsAwake;

    // You can spy if anyone is awake — simple OR chain
    public static bool CanSpy(bool knightIsAwake, bool archerIsAwake, bool prisonerIsAwake)
        => knightIsAwake || archerIsAwake || prisonerIsAwake;

    // Only signal if the archer won't notice and the prisoner is awake
    public static bool CanSignalPrisoner(bool archerIsAwake, bool prisonerIsAwake)
        => !archerIsAwake && prisonerIsAwake;

    // Two paths:
    // 1. Dog present AND archer asleep
    // 2. No knight, no archer, prisoner awake
    public static bool CanFreePrisoner(
        bool knightIsAwake,
        bool archerIsAwake,
        bool prisonerIsAwake,
        bool petDogIsPresent
    )
        => (petDogIsPresent && !archerIsAwake)
           || (!knightIsAwake && !archerIsAwake && prisonerIsAwake);
}
