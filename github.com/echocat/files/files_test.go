package files

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadContent(t *testing.T) {
	testCases := []struct {
		name          string
		path          string
		isSuccessful  bool
		expectedLines int
	}{
		{"ok", "../magazines.csv", true, 6},
		{"missing file", "./not_existing.file", false, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lines, err := LoadContent(tc.path)
			if tc.isSuccessful {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
			assert.Equal(t, tc.expectedLines, len(lines))
		})
	}
}
