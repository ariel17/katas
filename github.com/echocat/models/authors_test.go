package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadAuthors(t *testing.T) {
	testCases := []struct {
		name            string
		path            string
		isSuccessful    bool
		expectedAuthors int
	}{
		{"ok", "../authors.csv", true, 6},
		{"missing file", "../not_existing_file.csv", false, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			authors, err := LoadAuthors(tc.path)
			if tc.isSuccessful {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
			assert.Equal(t, tc.expectedAuthors, len(authors))
			for _, author := range authors {
				assert.NotEmpty(t, author.Email)
				assert.NotEmpty(t, author.FirstName)
				assert.NotEmpty(t, author.LastName)
			}
		})
	}
}
