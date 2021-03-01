package models

import "strings"

type Publication struct {
	Title       string
	ISBN        string
	Authors     []string
}

func parsePublication(line string) Publication {
	fields := strings.Split(line, ";")
	return Publication{
		Title: fields[0],
		ISBN: fields[1],
		Authors: strings.Split(fields[2], ","),
	}
}
