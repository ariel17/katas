package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublication_Print(t *testing.T) {
	p := Publication{
		Title: "A title",
		ISBN: "111-222-333",
		Authors: []string{"Me", "Me Too"},
	}
	assert.Equal(t, "A title\t111-222-333\tMe, Me Too", p.Print())
}
