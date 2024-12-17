package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("AoC Example 1", func(t *testing.T) {
		input := []byte("2333133121414131402")

		result := DefragChksum(input)
		assert.Equal(t, 1928, result)
	})
}
