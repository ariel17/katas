package files

import (
	"bufio"
	"os"
)

func LoadContent(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()

	lines := []string{}
	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		count++
		if count == 1 {
			continue
		}
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}