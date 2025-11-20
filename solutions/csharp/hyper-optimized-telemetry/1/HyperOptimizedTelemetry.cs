public static class TelemetryBuffer
{
    public static byte[] ToBuffer(long reading)
    {
        var buffer = new byte[9];

        // --- POSITIVE VALUES -------------------------------------------------

        if (reading >= 0 && reading <= ushort.MaxValue)
        {
            buffer[0] = 2;
            BitConverter.GetBytes((ushort)reading).CopyTo(buffer, 1);
            return buffer;
        }

        if (reading >= 0 && reading <= int.MaxValue)
        {
            buffer[0] = (byte)(256 - 4);     // 252
            BitConverter.GetBytes((int)reading).CopyTo(buffer, 1);
            return buffer;
        }

        if (reading >= 0 && reading <= uint.MaxValue)
        {
            buffer[0] = 4;
            BitConverter.GetBytes((uint)reading).CopyTo(buffer, 1);
            return buffer;
        }

        // --- NEGATIVE VALUES -------------------------------------------------

        if (reading >= short.MinValue && reading <= -1)
        {
            buffer[0] = (byte)(256 - 2);     // 254
            BitConverter.GetBytes((short)reading).CopyTo(buffer, 1);
            return buffer;
        }

        if (reading >= int.MinValue && reading < short.MinValue)
        {
            buffer[0] = (byte)(256 - 4);     // 252
            BitConverter.GetBytes((int)reading).CopyTo(buffer, 1);
            return buffer;
        }

        // --- EVERYTHING ELSE → LONG -----------------------------------------

        buffer[0] = (byte)(256 - 8);         // 248
        BitConverter.GetBytes(reading).CopyTo(buffer, 1);
        return buffer;
    }


    public static long FromBuffer(byte[] buffer)
    {
        int prefix = buffer[0];
    
        bool isSigned = prefix >= 128;
        int payloadBytes = isSigned ? 256 - prefix : prefix;
    
        // Prefix claims more bytes than can exist → value = 0 (per tests)
        if (payloadBytes > 8)
            return 0;
    
        long result = 0;
    
        // Little-endian decode
        for (int i = 0; i < payloadBytes; i++)
        {
            result |= ((long)buffer[1 + i]) << (8 * i);
        }
    
        // Unsigned? Done.
        if (!isSigned)
            return result;
    
        // Signed: only sign-extend if payloadBytes < 8
        if (payloadBytes < 8)
        {
            // Check sign bit of highest byte
            bool signBitSet = (buffer[1 + (payloadBytes - 1)] & 0x80) != 0;
    
            if (signBitSet)
            {
                long mask = -1L << (payloadBytes * 8);
                result |= mask;
            }
        }
    
        return result;
    }


}
