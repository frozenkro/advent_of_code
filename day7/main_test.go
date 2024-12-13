package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestValidCount(t *testing.T) {
  t.Run("AoC example", func(t *testing.T) {
    data := []equation{
      { total: 190, operands: []int{ 10, 19 } },
      { total: 3267, operands: []int{ 81, 40, 27 } },
      { total: 83, operands: []int{ 17, 5 } },
      { total: 156, operands: []int{ 15, 6 } },
      { total: 7290, operands: []int{ 6, 8, 6, 15 } },
      { total: 161011, operands: []int{ 16, 10, 13 } },
      { total: 192, operands: []int{ 17, 8, 14 } },
      { total: 21037, operands: []int{ 9, 7, 18, 13 } },
      { total: 292, operands: []int{ 11, 6, 16, 20 } },
    }

    total := validSum(data, false)
    assert.Equal(t, 3749, total)
  })
  t.Run("AoC example II", func(t *testing.T) {
    data := []equation{
      { total: 190, operands: []int{ 10, 19 } },
      { total: 3267, operands: []int{ 81, 40, 27 } },
      { total: 83, operands: []int{ 17, 5 } },
      { total: 156, operands: []int{ 15, 6 } },
      { total: 7290, operands: []int{ 6, 8, 6, 15 } },
      { total: 161011, operands: []int{ 16, 10, 13 } },
      { total: 192, operands: []int{ 17, 8, 14 } },
      { total: 21037, operands: []int{ 9, 7, 18, 13 } },
      { total: 292, operands: []int{ 11, 6, 16, 20 } },
    }

    total := validSum(data, true)
    assert.Equal(t, 11387, total)
  })
}
