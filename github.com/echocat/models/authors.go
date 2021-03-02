package models

import (
	"strings"

	"github.com/ariel17/katas/github.com/echocat/files"
)

type Author struct {
	Email     string
	FirstName string
	LastName  string
}

func (a Author) Print() string {
	return ""
}

func LoadAuthors(path string) ([]Author, error) {
	lines, err := files.LoadContent(path)
	if err != nil {
		return []Author{}, err
	}
	authors := []Author{}
	for _, line := range lines {
		authors = append(authors, parseAuthor(line))
	}
	return authors, nil
}

func parseAuthor(line string) Author {
	fields := strings.Split(line, ";")
	return Author{
		Email:     fields[0],
		FirstName: fields[1],
		LastName:  fields[2],
	}
}
