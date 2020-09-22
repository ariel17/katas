package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDataOK(t *testing.T) {
	data, err := loadData("weather.dat")
	assert.Nil(t, err)
	assert.Equal(t, 30, len(data))
	assert.Equal(t, float32(59), data[20].Min)
	assert.Equal(t, float32(86), data[20].Max)
}

func TestSmallestSpread(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		d := []data{
			data{1, 10, 25},
			data{2, 5, 5.5},
			data{3, 6, 9},
			data{4, -6, 20},
			data{5, 5, 6},
			data{5, -20, -6},
		}

		r := smallestSpread(d)
		assert.Equal(t, 2, r.Day)
	})

	t.Run("empty array", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("should have failed!")
			}
		}()
		smallestSpread([]data{})
	})
}