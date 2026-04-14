class RemoteControlCar
{
    public int Battery { get; set; }
    public int BatteryDrain { get; set; }
    public int Distance { get; set; }
    public int Speed { get; set; }

    public RemoteControlCar(int speed, int batteryDrain)
    {
        Battery = 100;
        BatteryDrain = batteryDrain;
        Speed = speed;
    }

    public bool BatteryDrained() => Battery < BatteryDrain;
    public int DistanceDriven() => Distance;

    public void Drive()
    {
        if (BatteryDrained()) return;
        Battery -= BatteryDrain;
        Distance += Speed;
    }

    public static RemoteControlCar Nitro() => new RemoteControlCar(50, 4);
}

class RaceTrack
{
    public int Distance { get; set; }

    public RaceTrack(int distance) => Distance = distance;

    public bool TryFinishTrack(RemoteControlCar car) =>
        car.Battery / car.BatteryDrain * car.Speed >= Distance;
}