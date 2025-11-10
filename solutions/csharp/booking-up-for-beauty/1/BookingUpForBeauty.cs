using System;

static class Appointment
{
    /// <summary>Parses a date/time string into a DateTime object.</summary>
    public static DateTime Schedule(string appointmentDateDescription)
        => DateTime.Parse(appointmentDateDescription);

    /// <summary>Returns true if the appointment time has already passed.</summary>
    public static bool HasPassed(DateTime appointmentDate)
        => DateTime.Now > appointmentDate;

    /// <summary>Returns true if the appointment is between noon and 6 PM.</summary>
    public static bool IsAfternoonAppointment(DateTime appointmentDate)
        => appointmentDate.Hour >= 12 && appointmentDate.Hour < 18;

    /// <summary>Returns a human-readable description of the appointment.</summary>
    public static string Description(DateTime appointmentDate)
        => $"You have an appointment on {appointmentDate:G}.";

    /// <summary>Returns the fixed anniversary date for the current year (September 15).</summary>
    public static DateTime AnniversaryDate()
        => new DateTime(DateTime.Now.Year, 9, 15);
}
