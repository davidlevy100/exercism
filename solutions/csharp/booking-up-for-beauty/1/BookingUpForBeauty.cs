static class Appointment
{
    // Parse a date string into a DateTime
    public static DateTime Schedule(string appointmentDateDescription) 
        => DateTime.Parse(appointmentDateDescription);

    // Check if the appointment already happened
    public static bool HasPassed(DateTime appointmentDate) 
        => DateTime.Now > appointmentDate;

    // Check if time is between noon and 6pm
    public static bool IsAfternoonAppointment(DateTime appointmentDate) 
        => appointmentDate.Hour >= 12 && appointmentDate.Hour < 18;

    // Describe the appointment in plain language
    public static string Description(DateTime appointmentDate) 
        => $"You have an appointment on {appointmentDate}.";

    // Return a fixed anniversary date
    public static DateTime AnniversaryDate() 
        => new DateTime(DateTime.Now.Year, 9, 15);
}
