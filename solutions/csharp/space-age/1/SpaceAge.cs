public class SpaceAge
{
    private readonly double _earthYears;

    public SpaceAge(int seconds) =>
        _earthYears = seconds / 31557600.0;

    public double OnMercury() => _earthYears / 0.2408467;
    public double OnVenus()   => _earthYears / 0.61519726;
    public double OnEarth()   => _earthYears;
    public double OnMars()    => _earthYears / 1.8808158;
    public double OnJupiter() => _earthYears / 11.862615;
    public double OnSaturn()  => _earthYears / 29.447498;
    public double OnUranus()  => _earthYears / 84.016846;
    public double OnNeptune() => _earthYears / 164.79132;
}