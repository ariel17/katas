package ariel17;

import java.nio.ByteBuffer;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.function.ToIntFunction;

public class MD5HashToInt implements ToIntFunction<String> {

    private static final String MD5 = "MD5";

    @Override
    public int applyAsInt(String s) {
        MessageDigest md = null;
        try {
            md = MessageDigest.getInstance(MD5);
        } catch (NoSuchAlgorithmException e) {
            return 0;
        }
        byte[] result = md.digest(s.getBytes());
        return ByteBuffer.wrap(result).getInt();
    }
}
