using System;

// Represents the type of forum account.
public enum AccountType
{
    Guest,
    User,
    Moderator
}

// Represents permissions that can be combined using bitwise operations.
[Flags]
public enum Permission
{
    None   = 0,        // No permissions set
    Read   = 1 << 0,   // 0001
    Write  = 1 << 1,   // 0010
    Delete = 1 << 2,   // 0100
    All    = Read | Write | Delete
}

// Provides helper methods for working with permissions.
public static class Permissions
{
    // Returns the default permissions for a given account type.
    public static Permission Default(AccountType accountType) =>
        accountType switch
        {
            AccountType.Guest     => Permission.Read,
            AccountType.User      => Permission.Read | Permission.Write,
            AccountType.Moderator => Permission.All,
            _ => Permission.None
        };

    // Grants one or more permissions by setting their bits.
    public static Permission Grant(Permission current, Permission grant) =>
        current | grant;

    // Revokes one or more permissions by clearing their bits.
    public static Permission Revoke(Permission current, Permission revoke) =>
        current & ~revoke;

    // Checks whether all requested permissions are present.
    public static bool Check(Permission current, Permission check) =>
        (current & check) == check;
}
