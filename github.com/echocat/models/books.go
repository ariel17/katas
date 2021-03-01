package models

import (
	"bufio"
	"os"
	"strings"
)

type Book struct {
	Title       string
	ISBN        string
	Authors     []string
	Description string
}

func LoadBooks(path string) ([]Book, error) {
	f, err := os.Open(path)
	if err != nil {
		return []Book{}, err
	}
	defer f.Close()

	books := []Book{}
	scanner := bufio.NewScanner(f)
	line := 0
	for scanner.Scan() {
		line++
		if line == 1 {
			continue
		}
		book := buildFromLine(scanner.Text())
		books = append(books, book)
	}

	return books, scanner.Err()
}

func buildFromLine(line string) Book {
	fields := strings.Split(line, ";")
	return Book{
		Title: fields[0],
		ISBN: fields[1],
		Authors: strings.Split(fields[2], ","),
		Description: fields[3],
	}
}