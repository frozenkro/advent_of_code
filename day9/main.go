package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const free = -1

type file struct {
	address int
	size    int
}

func main() {
	d, err := getData("./input")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}

	result := DefragChksum(d)
	fmt.Printf("Result: %d\n", result)
}

func getData(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	sc := bufio.NewScanner(f)

	sc.Scan()
	return []byte(sc.Text()), nil
}

func DefragChksum(data []byte) int {
	dataMap := createDataMap(data)
	result := defragDataMap(dataMap)
	return result
}

func createDataMap(src []byte) []int {
	dataMap := []int{}
	isEven := true
	for i, v := range src {
		var value int
		if isEven {
			value = i / 2
		} else {
			value = free
		}

		intV, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}
		frag := createFrag(value, intV)

		dataMap = append(dataMap, frag...)
		isEven = !isEven
	}
	return dataMap
}

func createFrag(value int, ln int) []int {
	f := make([]int, ln)
	for i := 0; i < ln; i++ {
		f[i] = value
	}
	return f
}

func defragDataMap(d []int) int {
	i := 0
	ii := len(d) - 1

	total := 0

	for i < ii {

		for d[i] != free {
			total += i * int(d[i])
			i++
		}
		for d[ii] == free {
			ii--
		}

		d[i] = d[ii]
		d[ii] = free

	}
	return total
}
