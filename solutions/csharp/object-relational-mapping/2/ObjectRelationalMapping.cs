using System;

public sealed class Orm : IDisposable
{
    private readonly Database _database;

    public Orm(Database database) => _database = database;

    // Start a transaction.
    // Database enforces correct state and throws if invalid.
    public void Begin() => _database.BeginTransaction();

    // Write within a transaction.
    // Any failure resets the database to a safe (Closed) state.
    public void Write(string data)
    {
        try
        {
            _database.Write(data);
        }
        catch
        {
            _database.Dispose();
        }
    }

    // Commit the transaction.
    // Failures are handled defensively by resetting the database.
    public void Commit()
    {
        try
        {
            _database.EndTransaction();
        }
        catch
        {
            _database.Dispose();
        }
    }

    // Ensure no transaction is left open if ORM is disposed early.
    public void Dispose()
    {
        if (_database.DbState != Database.State.Closed)
            _database.Dispose();
    }
}
