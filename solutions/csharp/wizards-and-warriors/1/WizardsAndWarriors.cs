using System;

/// <summary>Base character class defining shared behavior.</summary>
abstract class Character
{
    private readonly string characterType;
    private bool isVulnerable = false;

    protected Character(string characterType) => this.characterType = characterType;

    /// <summary>Calculates damage dealt to a target.</summary>
    public abstract int DamagePoints(Character target);

    /// <summary>Returns whether the character is currently vulnerable.</summary>
    public virtual bool Vulnerable() => isVulnerable;

    public override string ToString() => $"Character is a {characterType}";
}

/// <summary>A warrior with fixed damage values depending on target vulnerability.</summary>
class Warrior : Character
{
    public Warrior() : base("Warrior") { }

    public override int DamagePoints(Character target) => target.Vulnerable() ? 10 : 6;
}

/// <summary>A wizard who can prepare a spell for stronger attacks and less vulnerability.</summary>
class Wizard : Character
{
    private bool hasSpellPrepared = false;

    public Wizard() : base("Wizard") { }

    public override int DamagePoints(Character target) => hasSpellPrepared ? 12 : 3;

    public void PrepareSpell() => hasSpellPrepared = true;

    public override bool Vulnerable() => !hasSpellPrepared;
}
