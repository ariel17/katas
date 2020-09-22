package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitLine(t *testing.T) {
	line := "     col1  col2 col3   col4"
	expected := []string{"col1", "col2", "col3", "col4"}
	assert.Equal(t, expected, splitLine(line))
}

func TestLoadData(t *testing.T) {
	testCases := []struct {
		name          string
		path          string
		expectedLines int
		linesToAvoid  []int
		isSuccessful  bool
	}{
		{"ok for version 1", "./weather.dat", 31, []int{0, 10}, true},
		{"ok for version 2", "./football.dat", 20, []int{0, 10}, true},
		{"failed by not found", "./not-found.dat", 0, []int{}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			content, err := loadData(tc.path, tc.linesToAvoid)
			assert.Equal(t, tc.isSuccessful, err == nil)
			assert.Equal(t, tc.isSuccessful, content != nil)
			if tc.isSuccessful {
				assert.Equal(t, tc.expectedLines, len(content))
			}
		})
	}
}

func TestDataIntoTeams(t *testing.T) {
	testCases := []struct {
		name         string
		data         [][]string
		isSuccessful bool
	}{
		{
			"ok",
			[][]string{
				[]string{"1.", "Arsenal", "38", "26", "9", "3", "79", "-", "36", "87"},
			},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			teams, err := dataIntoTeams(tc.data)
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
			assert.Equal(t, tc.isSuccessful, err == nil)
			assert.Equal(t, tc.isSuccessful, teams != nil)
		})
	}
}

func TestDataIntoWeather(t *testing.T) {
	testCases := []struct {
		name         string
		data         [][]string
		isSuccessful bool
	}{
		{
			"ok",
			[][]string{
				[]string{"1", "88", "59", "74", "53.8", "0.00", "F", "280", "9.6", "270", "17", "1.6", "93", "23", "1004.5"},
			},
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			weathers, err := dataIntoWeather(tc.data)
			assert.Equal(t, tc.isSuccessful, err == nil)
			assert.Equal(t, tc.isSuccessful, weathers != nil)
		})
	}
}


func TestSmallestSpread(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		data := []weather{
			weather{1, 10, 25},
			weather{2, 5, 5.5},
			weather{3, 6, 9},
			weather{4, -6, 20},
			weather{5, 5, 6},
			weather{5, -20, -6},
		}

		r := smallestSpread(data)
		assert.Equal(t, 2, r.Day)
	})

	t.Run("empty array", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("should have failed!")
			}
		}()
		smallestSpread([]weather{})
	})
}
