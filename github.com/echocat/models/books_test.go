package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadBooks(t *testing.T) {
	testCases := []struct{
		name string
		path string
		amount int
		isSuccessful bool
	}{
		{"ok CSV", "../books.csv", 8, true},
		{"file does not exist", "./not_exists.csv", 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			books, err := LoadBooks(tc.path)
			if tc.isSuccessful {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
			assert.Equal(t, tc.amount, len(books))
			for _, book := range books {
				assert.NotEmpty(t, book.Title)
				assert.NotEmpty(t, book.ISBN)
				assert.GreaterOrEqual(t, len(book.Authors), 1)
				assert.NotEmpty(t, book.Description)
			}
		})
	}
}