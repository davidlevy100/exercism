using System;
using System.Runtime.InteropServices;
using System.Globalization;

public enum Location
{
    NewYork,
    London,
    Paris
}

public enum AlertLevel
{
    Early,
    Standard,
    Late
}

public static class Appointment
{
    public static DateTime ShowLocalTime(DateTime dtUtc) => dtUtc.ToLocalTime();

    public static DateTime Schedule(string appointmentDateDescription, Location location)
    {
        var local = DateTime.Parse(appointmentDateDescription);
        var tz    = GetTimeZone(location);

        return TimeZoneInfo.ConvertTimeToUtc(local, tz);
    }

    public static DateTime GetAlertTime(DateTime appointment, AlertLevel alertLevel)
    {
        switch (alertLevel)
        {
            case AlertLevel.Early:
                return appointment.AddDays(-1);
            case AlertLevel.Standard:
                return appointment.AddMinutes(-105);
            case AlertLevel.Late:
                return appointment.AddMinutes(-30);
            default:
                throw new ArgumentOutOfRangeException(nameof(alertLevel));
        }
    }

    public static bool HasDaylightSavingChanged(DateTime dt, Location location)
    {
        var tz = GetTimeZone(location);
        return tz.IsDaylightSavingTime(dt) != tz.IsDaylightSavingTime(dt.AddDays(-7));
    }

    public static DateTime NormalizeDateTime(string dtStr, Location location)
    {
        string format =
            location == Location.NewYork ? "MM/dd/yyyy HH:mm:ss" :
            "dd/MM/yyyy HH:mm:ss"; // London + Paris

        if (!DateTime.TryParseExact(
            dtStr,
            format,
            CultureInfo.InvariantCulture,
            DateTimeStyles.None,
            out DateTime local))
        {
            return DateTime.MinValue;
        }

        var tz = GetTimeZone(location);
        var utc = TimeZoneInfo.ConvertTimeToUtc(local, tz);
        return TimeZoneInfo.ConvertTimeFromUtc(utc, tz);
    }

    private static TimeZoneInfo GetTimeZone(Location location)
    {
        bool win = RuntimeInformation.IsOSPlatform(OSPlatform.Windows);

        string tzId =
            location == Location.NewYork ? (win ? "Eastern Standard Time"   : "America/New_York") :
            location == Location.London  ? (win ? "GMT Standard Time"       : "Europe/London") :
            location == Location.Paris   ? (win ? "W. Europe Standard Time" : "Europe/Paris") :
            throw new ArgumentOutOfRangeException(nameof(location));

        return TimeZoneInfo.FindSystemTimeZoneById(tzId);
    }
}
