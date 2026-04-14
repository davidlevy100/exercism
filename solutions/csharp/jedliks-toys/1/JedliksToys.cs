class RemoteControlCar
{
    public int Battery { get; set; } = 100;
    public int Distance { get; set; } = 0;
    
    public static RemoteControlCar Buy() => new RemoteControlCar();

    public string DistanceDisplay() => $"Driven {Distance} meters";

    public string BatteryDisplay() => Battery == 0 ? "Battery empty" : $"Battery at {Battery}%";

    public void Drive()
    {
        if (Battery <= 0) return;
        Battery -= 1;
        Distance += 20;
    }
}
