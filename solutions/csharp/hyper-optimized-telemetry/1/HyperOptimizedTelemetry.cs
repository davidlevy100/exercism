public static class TelemetryBuffer
{
    public static byte[] ToBuffer(long reading)
    {
        var buffer = new byte[9];
        
        if (reading >= 4_294_967_296L)
        {
            buffer[0] = 248;
            BitConverter.GetBytes((long)reading).CopyTo(buffer, 1);
        }
        else if (reading >= 2_147_483_648L)
        {
            buffer[0] = 4;
            BitConverter.GetBytes((uint)reading).CopyTo(buffer, 1);
        }
        else if (reading >= 65_536)
        {
            buffer[0] = 252;
            BitConverter.GetBytes((int)reading).CopyTo(buffer, 1);
        }
        else if (reading >= 0)
        {
            buffer[0] = 2;
            BitConverter.GetBytes((ushort)reading).CopyTo(buffer, 1);
        }
        else if (reading >= -32_768)
        {
            buffer[0] = 254;
            BitConverter.GetBytes((short)reading).CopyTo(buffer, 1);
        }
        else if (reading >= -2_147_483_648L)
        {
            buffer[0] = 252;
            BitConverter.GetBytes((int)reading).CopyTo(buffer, 1);
        }
        else
        {
            buffer[0] = 248;
            BitConverter.GetBytes(reading).CopyTo(buffer, 1);
        }
        
        return buffer;
    }

    public static long FromBuffer(byte[] buffer) =>
        buffer[0] switch
        {
            2   => BitConverter.ToUInt16(buffer, 1),
            4   => BitConverter.ToUInt32(buffer, 1),
            254 => BitConverter.ToInt16(buffer, 1),
            252 => BitConverter.ToInt32(buffer, 1),
            248 => BitConverter.ToInt64(buffer, 1),
            _   => 0
        };
}