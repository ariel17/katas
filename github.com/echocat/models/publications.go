package models

import (
	"fmt"
	"strings"
)

type Publication struct {
	Title   string
	ISBN    string
	Authors []string
}

func (p Publication) Print() string {
	return fmt.Sprintf("%s\t%s\t%s", p.Title, p.ISBN, strings.Join(p.Authors, ", "))
}

func parsePublication(line string) Publication {
	fields := strings.Split(line, ";")
	return Publication{
		Title:   fields[0],
		ISBN:    fields[1],
		Authors: strings.Split(fields[2], ","),
	}
}
