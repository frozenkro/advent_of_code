package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	t.Run("simple correct string", func(t *testing.T) {
		input := []byte("mul(5,6)mul(6,6)")
		res, err := parse(input)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, int64(66), res)
	})
	t.Run("simple incorrect string", func(t *testing.T) {
		input := []byte("mul(5,x)mul(6,6)")
		res, err := parse(input)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, int64(36), res)
	})
	t.Run("3 digit nums", func(t *testing.T) {
		input := []byte("mul(1,655)mul(100,1)")
		res, err := parse(input)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, int64(755), res)
	})
	t.Run("4 digit nums", func(t *testing.T) {
		input := []byte("mul(1,6555)mul(1000,1)")
		res, err := parse(input)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, int64(0), res)
	})
	t.Run("keyword interrupt parse", func(t *testing.T) {
		input := []byte("mul(1,mul(1,1)")
		res, err := parse(input)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, int64(1), res)
	})
}
