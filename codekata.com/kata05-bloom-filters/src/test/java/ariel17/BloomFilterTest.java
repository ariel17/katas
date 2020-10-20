package ariel17;


import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.function.ToIntFunction;

import static org.junit.jupiter.api.Assertions.*;

public class BloomFilterTest {

    public static final String DICT_PATH = "/words";

    private BloomFilter bf;

    private String word;

    private List<ToIntFunction<String>> functions;

    @BeforeEach
    public void setUp() {
        functions = new ArrayList<>(){{
            add(new MD5HashToInt());
        }};
        bf = new BloomFilter(1024, functions);
        word = "Hello";
    }

    @Test
    public void loadOK() throws IOException {
        bf.load(DICT_PATH);
    }

    @Test
    public void loadFailed() {
        assertThrows(IOException.class, () -> {
            bf.load("/file/path/unknown");
        });
    }

    @Test
    public void doContainsNotNull() {
        bf.add(word);
        assertTrue(bf.contains(word));
    }

    @Test
    public void notContainsNull() {
        assertFalse(bf.contains(null));
    }

    @Test
    public void notContains() {
        String otherWord = "xdsfkso";
        assertFalse(bf.contains(otherWord));
    }
}
