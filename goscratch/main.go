package main

import (
	"fmt"
	"strconv"
)

func main() {
	badNumChars := []rune{
		'o',
		'1',
		'5',
	}

	num, err := strconv.Atoi(string(badNumChars))
	if err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	fmt.Printf("result: %d", num)
}
