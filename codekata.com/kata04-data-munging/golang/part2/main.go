package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type team struct {
	Name    string
	For     int
	Against int
}

func (t team) Diff() int {
	return int(math.Abs(float64(t.For - t.Against)))
}

func splitLine(line string) []string {
	parts := strings.Split(line, " ")
	var s []string
	for _, v := range parts {
		if v != "" {
			s = append(s, v)
		}
	}
	return s
}

func loadData(path string) ([]team, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(b)), "\n")
	teams := []team{}

	for i, line := range lines {
		if i == 0 || i == 18 {
			continue
		}
		parts := splitLine(line)
		f, err := strconv.Atoi(parts[6])
		if err != nil {
			return nil, err
		}
		a, err := strconv.Atoi(parts[8])
		if err != nil {
			return nil, err
		}
		t := team{
			Name:    parts[1],
			For:     f,
			Against: a,
		}
		teams = append(teams, t)
	}

	return teams, nil
}

func getMinimumDiff(teams []team) team {
	var tt team
	minorDiff := math.MaxInt32
	for _, t := range teams {
		diff := t.Diff()
		if diff < minorDiff {
			minorDiff = diff
			tt = t
		}
	}
	return tt
}

func main() {
	teams, err := loadData("football.dat")
	if err != nil {
		panic(err)
	}
	tt := getMinimumDiff(teams)
	fmt.Printf("team: %s, diff: %d\n", tt.Name, tt.Diff())
}
