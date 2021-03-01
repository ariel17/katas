package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadMagazines(t *testing.T) {
	testCases := []struct{
		name string
		path string
		isSuccessful bool
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
