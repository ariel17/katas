package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type data struct {
	Day int
	Min float32
}

func main() {
	content, err := loadData("weather.dat")
	if err != nil {
		panic(err)
	}
	if len(content) == 0 {
		fmt.Println("File is empty.")
	}
	for _, d := range content {
		fmt.Printf("* Day: %d, Min: %.2f\n", d.Day, d.Min)
	}
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

		rawMin := strings.Replace(parsed[2], "*", "", -1)
		min, err := strconv.ParseFloat(rawMin, 32)
		if err != nil {
			return nil, err
		}

		parsedContent = append(parsedContent, data{
			Day: day,
			Min: float32(min),
		})
	}

	return parsedContent, nil
}
