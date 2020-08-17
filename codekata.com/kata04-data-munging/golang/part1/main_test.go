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
}
