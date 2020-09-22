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

type data struct {
	Day int
	Min float32
	Max float32
}

func (d data) Diff() float32 {
	return d.Max-d.Min
}

func main() {
	content, err := loadData("weather.dat")
	if err != nil {
		panic(err)
	}
	if len(content) == 0 {
		fmt.Println("File is empty.")
	}
	min := smallestSpread(content)
	fmt.Printf("* Day: %d, Min: %.2f, Max: %.2f, Spread: %.2f", min.Day, min.Min, min.Max, min.Diff())
}

func loadData(filePath string) ([]data, error) {
	f, err := os.Open(filePath)

	defer f.Close()

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	rawContent := strings.Split(strings.TrimSpace(string(b)), "\n")
	parsedContent := []data{}

	for i, line := range rawContent {
		if i < 2 || i == len(rawContent)-1 {
			continue
		}
		parsed := []string{}
		parts := strings.Split(line, " ")
		for _, p := range parts {
			if len(p) > 0 {
				parsed = append(parsed, p)
			}
		}

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

		parsedContent = append(parsedContent, data{
			Day: day,
			Min: float32(min),
			Max: float32(max),
		})
	}

	return parsedContent, nil
}

func smallestSpread(d []data) data {
	if len(d) == 0 {
		panic(errors.New("weather data cannot be empty"))
	}
	var (
		min data
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