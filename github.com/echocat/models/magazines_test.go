package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadMagazines(t *testing.T) {
	testCases := []struct {
		name              string
		path              string
		isSuccessful      bool
		expectedMagazines int
	}{
		{"ok", "../magazines.csv", true, 6},
		{"missing file", "../not_existing_file.csv", false, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			magazines, err := LoadMagazines(tc.path)
			if tc.isSuccessful {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
			assert.Equal(t, tc.expectedMagazines, len(magazines))
			for _, magazine := range magazines {
				assert.NotEmpty(t, magazine.Title)
				assert.NotEmpty(t, magazine.ISBN)
				assert.GreaterOrEqual(t, len(magazine.Authors), 1)
				assert.False(t, magazine.PublishedAt.IsZero())
			}
		})
	}
}

func TestMagazine_Print(t *testing.T) {
	date, _ := time.Parse("02.03.2006", "25.02.2020")
	m := Magazine{
		Publication: Publication{
			Title: "Title",
			ISBN: "111-222-333",
			Authors: []string{"Me", "Me Too"},
		},
		PublishedAt: date,
	}
	assert.Equal(t, "Title\t111-222-333\tMe, Me Too\t2020-02-25", m.Print())
}
