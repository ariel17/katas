package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type weather struct {
	Day int
	Min float32
	Max float32
}

func (w weather) Diff() float32 {
	return w.Max-w.Min
}

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

func loadData(path string, linesToAvoid []int) ([][]string, error) {
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
	var data [][]string

	for i, line := range lines {
		mustIgnore := false
		for _, lta := range linesToAvoid {
			if i == lta {
				mustIgnore = true
			}
		}
		if mustIgnore {
			continue
		}
		data = append(data, splitLine(line))
	}

	return data, nil
}

func dataIntoWeather(data [][]string) ([]weather, error) {
	var weathers []weather
	for _, parsed := range data {
		day, err := strconv.Atoi(parsed[0])
		if err != nil {
			return nil, err
		}

		rawMax := strings.Replace(parsed[1], "*", "", -1)
		max, err := strconv.ParseFloat(rawMax, 32)
		if err != nil {
			return nil, err
		}

		rawMin := strings.Replace(parsed[2], "*", "", -1)
		min, err := strconv.ParseFloat(rawMin, 32)
		if err != nil {
			return nil, err
		}

		weathers = append(weathers, weather{
			Day: day,
			Min: float32(min),
			Max: float32(max),
		})
	}
	return weathers, nil
}

func dataIntoTeams(data [][]string) ([]team, error) {
	var teams []team
	for _, parts := range data {
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
	weathersData, err := loadData("weather.dat", []int{0, 1, 32})
	if err != nil {
		panic(err)
	}
	weathers, err := dataIntoWeather(weathersData)
	if err != nil {
		panic(err)
	}
	min := smallestSpread(weathers)
	fmt.Printf("* Day: %d, Min: %.2f, Max: %.2f, Spread: %.2f\n", min.Day, min.Min, min.Max, min.Diff())

	teamsData, err := loadData("football.dat", []int{0, 18})
	if err != nil {
		panic(err)
	}
	teams, err := dataIntoTeams(teamsData)
	if err != nil {
		panic(err)
	}
	tt := getMinimumDiff(teams)
	fmt.Printf("* Team: %s, Against: %d, For: %d, diff: %d\n", tt.Name, tt.Against, tt.For, tt.Diff())
}

func smallestSpread(d []weather) weather {
	if len(d) == 0 {
		panic(errors.New("weather data cannot be empty"))
	}
	var (
		min weather
		v = float32(math.MaxFloat32)
	)
	for _, day := range d {
		diff := day.Diff()
		if diff < v {
			v = diff
			min = day
		}
	}
	return min
}
