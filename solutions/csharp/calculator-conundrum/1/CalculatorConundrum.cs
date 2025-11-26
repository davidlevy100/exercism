public static class SimpleCalculator
{
    public static string Calculate(int operand1, int operand2, string? operation)
    {
        if (operation is null)
            throw new ArgumentNullException(nameof(operation));

        if (operation == "")
            throw new ArgumentException(nameof(operation));

        switch (operation)
        {
            case "+":
                return $"{operand1} + {operand2} = {operand1 + operand2}";

            case "*":
                return $"{operand1} * {operand2} = {operand1 * operand2}";

            case "/":
                return operand2 == 0
                    ? "Division by zero is not allowed."
                    : $"{operand1} / {operand2} = {operand1 / operand2}";

            default:
                throw new ArgumentOutOfRangeException(nameof(operation));
        }
    }
}
