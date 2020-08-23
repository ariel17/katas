package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTeam(t *testing.T) {
	tt := team{
		Name:    "River",
		For:     10,
		Against: 5,
	}
	assert.Equal(t, 5, tt.Diff())
}

func TestSplitLine(t *testing.T) {
	line := "     col1  col2 col3   col4"
	expected := []string{"col1", "col2", "col3", "col4"}
	assert.Equal(t, expected, splitLine(line))
}

func TestLoadData(t *testing.T) {
	testCases := []struct {
		name         string
		path         string
		isSuccessful bool
	}{
		{"ok", "./football.dat", true},
		{"failed by not found", "./not-found.dat", false},
		{"failed by non-numeric field", "./non-numeric.dat", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			teams, err := loadData(tc.path)
			if tc.isSuccessful {
				assert.Nil(t, err)
				assert.NotNil(t, teams)
				assert.Equal(t, 8, teams[10].Diff())
			} else {
				assert.NotNil(t, err)
				assert.Nil(t, teams)
			}
		})
	}
}

func TestGetMinimumDiff(t *testing.T) {
	teams := []team{
		team{Name: "A", For: 10, Against: 13},
		team{Name: "B", For: 8, Against: 9},
		team{Name: "C", For: 20, Against: 13},
	}

	tt := getMinimumDiff(teams)
	assert.Equal(t, "B", tt.Name)
	assert.Equal(t, 1, tt.Diff())
}
