package models

import (
	"strings"
	"github.com/ariel17/katas/github.com/echocat/files"
)

type Book struct {
	Publication
	Description string
}

func LoadBooks(path string) ([]Book, error) {
	books := []Book{}
	lines, err := files.LoadContent(path)
	if err != nil {
		return []Book{}, err
	}
	for _, line := range lines {
		book := parseBook(line)
		books = append(books, book)
	}

	return books, nil
}

func parseBook(line string) Book {
	fields := strings.Split(line, ";")
	return Book{
		Publication: parsePublication(line),
		Description: fields[3],
	}
}