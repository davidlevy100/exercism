class RemoteControlCar
{
    private int _distanceDriven = 0;   // meters driven
    private int _battery = 100;        // percentage remaining
    
    public static RemoteControlCar Buy() => new RemoteControlCar();

    // Display driven distance
    public string DistanceDisplay() => $"Driven {_distanceDriven} meters";

    // Display battery status
    public string BatteryDisplay() =>
        _battery > 0
            ? $"Battery at {_battery}%"
            : "Battery empty";

    // Attempt to drive 20 meters, draining 1% battery
    public void Drive()
    {
        if (_battery <= 0)
            return;

        _distanceDriven += 20;
        _battery -= 1;
    }
}
