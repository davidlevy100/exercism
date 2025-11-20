public static class TelemetryBuffer
{
    public static byte[] ToBuffer(long reading)
    {
        byte[] bytes;
        bool signed;

        // Determine smallest fitting type
        if (reading >= 0 && reading <= ushort.MaxValue)
        {
            signed = false;
            bytes = BitConverter.GetBytes((ushort)reading);
        }
        else if (reading >= 0 && reading <= int.MaxValue)
        {
            signed = true;
            bytes = BitConverter.GetBytes((int)reading);
        }
        else if (reading >= 0 && reading <= uint.MaxValue)
        {
            signed = false;
            bytes = BitConverter.GetBytes((uint)reading);
        }
        else if (reading >= short.MinValue && reading <= -1)
        {
            signed = true;
            bytes = BitConverter.GetBytes((short)reading);
        }
        else if (reading >= int.MinValue && reading < short.MinValue)
        {
            signed = true;
            bytes = BitConverter.GetBytes((int)reading);
        }
        else
        {
            // long catch-all
            signed = true;
            bytes = BitConverter.GetBytes(reading);
        }

        // prefix = payload length for unsigned
        // prefix = 256 - payload length for signed
        byte prefix = signed
            ? (byte)(256 - bytes.Length)
            : (byte)bytes.Length;

        // Build fixed 9-byte buffer
        var buffer = new byte[9];
        buffer[0] = prefix;
        Array.Copy(bytes, 0, buffer, 1, bytes.Length);

        return buffer;
    }


    public static long FromBuffer(byte[] buffer)
    {
        byte prefix = buffer[0];

        // Signed long (8 bytes)
        if (prefix == 256 - sizeof(long))
            return BitConverter.ToInt64(buffer, 1);

        // Signed int (4 bytes)
        if (prefix == 256 - sizeof(int))
            return BitConverter.ToInt32(buffer, 1);

        // Signed short (2 bytes)
        if (prefix == 256 - sizeof(short))
            return BitConverter.ToInt16(buffer, 1);

        // Unsigned ushort (2 bytes)
        if (prefix == sizeof(ushort))
            return BitConverter.ToUInt16(buffer, 1);

        // Unsigned uint (4 bytes)
        if (prefix == sizeof(uint))
            return BitConverter.ToUInt32(buffer, 1);

        // Anything else is not a valid prefix → value = 0
        return 0;
    }
}
