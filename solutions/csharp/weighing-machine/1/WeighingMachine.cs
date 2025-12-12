class WeighingMachine
{
    private int precision;
    private double weight;
    private double tareAdjustment = 5.0; // default bias

    public WeighingMachine(int precision)
    {
        Precision = precision;
    }

    public int Precision
    {
        get => precision;
        private set
        {
            if (value < 0)
                throw new ArgumentException("Precision cannot be negative.");

            precision = value;
        }
    }

    public double Weight
    {
        get => weight;
        set
        {
            if (value < 0)
                throw new ArgumentOutOfRangeException(nameof(value), "Weight cannot be negative.");

            weight = value;
        }
    }

    public double TareAdjustment
    {
        get => tareAdjustment;
        set => tareAdjustment = value;
    }

    public string DisplayWeight
    {
        get
        {
            double adjustedWeight = weight - tareAdjustment;
            double rounded = Math.Round(adjustedWeight, precision);
            return $"{rounded.ToString($"F{precision}")} kg";
        }
    }
}
