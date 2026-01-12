using System;

// Defines the available forum account types.
public enum AccountType
{
    Guest,
    User,
    Moderator
}

// Defines combinable permissions using individual bits.
[Flags]
public enum Permission
{
    None   = 0,        // 0000 — no permissions
    Read   = 1 << 0,   // 0001 — read access
    Write  = 1 << 1,   // 0010 — write access
    Delete = 1 << 2,   // 0100 — delete access
    All    = Read | Write | Delete // 0111 — all permissions
}

// Helper methods for assigning and checking permissions.
public static class Permissions
{
    // Maps each account type to its default permission set.
    public static Permission Default(AccountType accountType) =>
        accountType switch
        {
            AccountType.Guest     => Permission.Read,
            AccountType.User      => Permission.Read | Permission.Write,
            AccountType.Moderator => Permission.All,
            _ => Permission.None
        };

    // Grants permissions by turning on the specified bits (bitwise OR).
    public static Permission Grant(Permission current, Permission grant) =>
        current | grant;

    // Revokes permissions by turning off the specified bits (AND with inverted mask).
    public static Permission Revoke(Permission current, Permission revoke) =>
        current & ~revoke;

    // Checks that all requested permission bits are set in the current value.
    public static bool Check(Permission current, Permission check) =>
        (current & check) == check;
}
