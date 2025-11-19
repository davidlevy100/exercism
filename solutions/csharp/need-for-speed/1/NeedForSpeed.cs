class RemoteControlCar
{
    private int _speed;
    private int _batteryDrain;
    private int _battery = 100;
    private int _distanceDriven = 0;

    public RemoteControlCar(int speed, int batteryDrain)
    {
        _speed = speed;               // distance gained per drive
        _batteryDrain = batteryDrain; // battery used per drive
    }

    // Battery is considered drained when it can't afford another drive
    public bool BatteryDrained() => _battery < _batteryDrain;

    // Total distance traveled so far
    public int DistanceDriven() => _distanceDriven;

    public void Drive()
    {
        if (BatteryDrained())         // do nothing if not enough battery
            return;

        _battery -= _batteryDrain;    // spend battery
        _distanceDriven += _speed;    // gain distance
    }

    // Factory method for a predefined fast car
    public static RemoteControlCar Nitro() => new RemoteControlCar(50, 4);
}

class RaceTrack
{
    private int _distance;

    public RaceTrack(int distance)
    {
        _distance = distance;         // required distance to finish track
    }

    public bool TryFinishTrack(RemoteControlCar car)
    {
        // Run until the car can no longer drive
        while (!car.BatteryDrained())
        {
            car.Drive();
        }

        // Success if total distance meets or exceeds track length
        return car.DistanceDriven() >= _distance;
    }
}
