using System;

public sealed class Orm : IDisposable
{
    private readonly Database _database;

    // ORM owns the database instance and is responsible for returning it
    // to a safe (Closed) state on failure or early shutdown.
    public Orm(Database database) => _database = database;

    // Starts a new transaction.
    // This is a strict state transition: callers must begin only from Closed.
    public void Begin()
    {
        if (_database.DbState != Database.State.Closed)
            throw new InvalidOperationException();

        _database.BeginTransaction();
    }

    // Writes data within the current transaction.
    // Any failure (bad data, invalid state, etc.) triggers cleanup to ensure
    // the database is not left in a partial or invalid transaction state.
    public void Write(string data)
    {
        try
        {
            _database.Write(data);
        }
        catch
        {
            // Database.Dispose() acts as a rollback/reset per the database contract.
            _database.Dispose();
        }
    }

    // Commits the active transaction.
    // Failures are handled defensively by resetting the database to Closed.
    public void Commit()
    {
        try
        {
            _database.EndTransaction();
        }
        catch
        {
            // Ensure no dangling or corrupted transaction remains.
            _database.Dispose();
        }
    }

    // Ensures the database is not left mid-transaction if the ORM
    // is disposed early by the caller.
    public void Dispose()
    {
        if (_database.DbState != Database.State.Closed)
            _database.Dispose();
    }
}
