package ariel17;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.util.List;
import java.util.function.ToIntFunction;


/**
 * Based on https://richardstartin.github.io/posts/building-a-bloom-filter-from-scratch
 */
public class BloomFilter {

    private static final Long WORD_BASE = 1L;

    private static final int WORD_INDEX = 6;

    private final int size;

    private final List<ToIntFunction<String>> functions;

    private final long[] array;

    public BloomFilter(int size, List<ToIntFunction<String>> functions) {
        this.size = size;
        this.functions = functions;
        array = new long[size];
    }

    public void add(String word) {
        if (word == null) {
            return;
        }
        for (ToIntFunction<String> f : functions) {
            int hash = mapHash(f.applyAsInt(word));
            array[hash >>> WORD_INDEX] |= WORD_BASE << hash;
        }
    }

    public boolean contains(String word) {
        if (word == null) {
            return false;
        }
        for (ToIntFunction<String> f : functions) {
            int hash = mapHash(f.applyAsInt(word));
            if ((array[hash >>> WORD_INDEX] & (WORD_BASE << hash)) == 0) {
                return false;
            }
        }
        return true;
    }

    public void load(String filePath) throws IOException {
        InputStream is = this.getClass().getResourceAsStream(filePath);
        if (is == null) {
            throw new IOException("file not found");
        }
        BufferedReader br = new BufferedReader(new InputStreamReader(is));
        String line;
        do {
            line = br.readLine();
            add(line);
        } while (line != null);
    }

    private int mapHash(int hash) {
        return hash & (size - 1);
    }
}

