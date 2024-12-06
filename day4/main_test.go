package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountInMatrix(t *testing.T) {
  t.Run("all directions", func(t *testing.T) {
    mtx := [][]byte{
      { 'S', '.', '.', 'S', '.', '.', 'S' },
      { '.', 'A', '.', 'A', '.', 'A', '.' },
      { '.', '.', 'M', 'M', 'M', '.', '.' },
      { 'S', 'A', 'M', 'X', 'M', 'A', 'S' },
      { '.', '.', 'M', 'M', 'M', '.', '.' },
      { '.', 'A', '.', 'A', '.', 'A', '.' },
      { 'S', '.', '.', 'S', '.', '.', 'S' },
    }

    res := CountInMatrix("XMAS", mtx)
    assert.Equal(t, 8, res)
  })
}
