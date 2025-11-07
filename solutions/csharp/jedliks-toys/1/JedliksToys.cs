class RemoteControlCar
{
    private int distanceDriven = 0;
    private int batteryPercentage = 100;

    /// <summary>Creates a new remote control car.</summary>
    public static RemoteControlCar Buy() => new RemoteControlCar();

    /// <summary>Returns the total distance driven.</summary>
    public string DistanceDisplay() => $"Driven {distanceDriven} meters";

    /// <summary>Returns the current battery status.</summary>
    public string BatteryDisplay() =>
        batteryPercentage > 0 ? $"Battery at {batteryPercentage}%" : "Battery empty";

    /// <summary>Drives the car forward 20 meters and drains 1% battery if available.</summary>
    public void Drive()
    {
        if (batteryPercentage <= 0) return;
        distanceDriven += 20;
        batteryPercentage = Math.Max(0, batteryPercentage - 1);
    }
}
