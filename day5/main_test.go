package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
  t.Run("basic test", func(t *testing.T) {
    rules, rows, err := getData("./test_input")
    assert.Nil(t, err)

    res := GetSumValidRows(rows, rules)
    assert.Equal(t, 143, res)
  })
}
