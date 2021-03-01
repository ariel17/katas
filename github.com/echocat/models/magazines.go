package models

import (
	"strings"
	"time"
	"github.com/ariel17/katas/github.com/echocat/files"
)

type Magazine struct {
	Publication
	PublishedAt time.Time
}

func LoadMagazines(path string) ([]Magazine, error) {
	lines, err := files.LoadContent(path)
	if err != nil {
		return []Magazine{}, err
	}
	magazines := []Magazine{}
	for _, line := range lines {
		magazine, err := parseMagazine(line)
		if err != nil {
			return magazines, err
		}
		magazines = append(magazines, magazine)
	}
	return magazines, nil
}

func parseMagazine(line string) (Magazine, error) {
	fields := strings.Split(line, ";")
	date, err := time.Parse("02.03.2006", fields[3])
	return Magazine{
		Publication: parsePublication(line),
		PublishedAt: date,
	}, err
}