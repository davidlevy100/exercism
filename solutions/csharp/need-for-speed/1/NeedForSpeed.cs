/// <summary>Simulates a remote control car.</summary>
class RemoteControlCar
{
    private int _speed;
    private int _batteryDrain;
    private int _distanceDriven = 0;
    private int _batteryLevel = 100;

    /// <summary>Creates a car with speed and battery drain.</summary>
    public RemoteControlCar(int speed, int batteryDrain)
    {
        _speed = speed;
        _batteryDrain = batteryDrain;
    }

    /// <summary>Checks if the battery is too low to drive.</summary>
    public bool BatteryDrained() => _batteryLevel < _batteryDrain;

    /// <summary>Gets the distance driven so far.</summary>
    public int DistanceDriven() => _distanceDriven;

    /// <summary>Drives once if enough battery remains.</summary>
    public void Drive()
    {
        if (!BatteryDrained())
        {
            _distanceDriven += _speed;
            _batteryLevel -= _batteryDrain;
        }
    }

    /// <summary>Returns a high-performance Nitro car.</summary>
    public static RemoteControlCar Nitro() => new RemoteControlCar(50, 4);
}

/// <summary>Represents a racetrack.</summary>
class RaceTrack
{
    private int _distance;

    /// <summary>Creates a track with a distance.</summary>
    public RaceTrack(int distance) => _distance = distance;

    /// <summary>Tests if a car can finish before draining.</summary>
    public bool TryFinishTrack(RemoteControlCar car)
    {
        while (!car.BatteryDrained())
        {
            car.Drive();
            if (car.DistanceDriven() >= _distance) return true;
        }
        return false;
    }
}
