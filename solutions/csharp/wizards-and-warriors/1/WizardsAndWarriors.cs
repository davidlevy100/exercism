// Base type for all characters.
// Handles shared behavior like type name and vulnerability defaults.
abstract class Character
{
    private readonly string _type;

    // Each character supplies its own type name.
    protected Character(string characterType) =>
        _type = characterType;

    // Subclasses must define how much damage they do to a target.
    public abstract int DamagePoints(Character target);

    // Default behavior: characters are NOT vulnerable unless they say otherwise.
    public virtual bool Vulnerable() => false;

    // Simple string representation.
    public override string ToString() => $"Character is a {_type}";
}

// Warrior: basic melee class with flat damage and no prep mechanics.
class Warrior : Character
{
    public Warrior() : base("Warrior") { }

    // Warriors do more damage if the target happens to be vulnerable.
    public override int DamagePoints(Character target) =>
        target.Vulnerable() ? 10 : 6;
}

// Wizard: can prepare a spell to become stronger and less vulnerable.
class Wizard : Character
{
    private bool _spellPrepared;

    public Wizard() : base("Wizard") { }

    // A wizard is only vulnerable when they haven't prepared a spell.
    public override bool Vulnerable() =>
        !_spellPrepared;

    // Spell prepared = big damage, otherwise weak.
    public override int DamagePoints(Character target) =>
        _spellPrepared ? 12 : 3;

    // Prepares the spell, changing both damage and vulnerability state.
    public void PrepareSpell() =>
        _spellPrepared = true;
}
