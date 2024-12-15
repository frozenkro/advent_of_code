package main

import (
	"bufio"
	"fmt"
	"os"
)

type frqGrps map[rune][]node
type node struct {
	x int
	y int
}

var (
	bound node
	grid  [][]byte
)

func main() {
	g, err := readData("./input")
	grid = g
	if err != nil {
		fmt.Printf("Error getting data: %v\n", err)
		os.Exit(1)
	}
	bound = node{x: len(grid[0]) - 1, y: len(grid) - 1}

	fg := freqGroups()
	result := antinodes(fg)

	fmt.Printf("Result: %d\n", result)
}

func readData(path string) ([][]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error reading input file: %v", err)
	}

	data := [][]byte{}
	sn := bufio.NewScanner(f)
	for sn.Scan() {
		data = append(data, []byte(sn.Text()))
	}

	return data, nil
}

func freqGroups() frqGrps {
	res := make(frqGrps)
	for y, row := range grid {
		for x, val := range row {
			if val != byte('.') {
				res.addNode(rune(val), x, y)
			}
		}
	}
	return res
}

func (f frqGrps) addNode(freq rune, x int, y int) {
	n := node{x: x, y: y}

	if f[freq] == nil {
		f[freq] = []node{n}
		return
	}

	f[freq] = append(f[freq], n)
}

func antinodes(frqGrps frqGrps) int {
	total := 0
	for _, nodes := range frqGrps {
		for i, n := range nodes {
			if i < len(nodes)-1 {
				for _, n2 := range nodes[i+1:] {
					total += compareNodes(n, n2)
				}
			}
		}
	}
	return total
}

func compareNodes(n1 node, n2 node) int {
	total := 0

	if valid, node := validAnti(n1, n2); valid {
		if grid[node.y][node.x] != byte('#') {
			grid[node.y][node.x] = byte('#')
			total++
		}
	}

	if valid, node := validAnti(n2, n1); valid {
		if grid[node.y][node.x] != byte('#') {
			grid[node.y][node.x] = byte('#')
			total++
		}
	}
	return total
}

func validAnti(n node, nx node) (bool, node) {
	anti := node{}

	if n.x > nx.x {
		anti.x = n.x + n.x - nx.x
	} else {
		anti.x = n.x - (nx.x - n.x)
	}

	if n.y > nx.y {
		anti.y = n.y + n.y - nx.y
	} else {
		anti.y = n.y - (nx.y - n.y)
	}

	valid := anti.y >= 0 &&
		anti.x >= 0 &&
		anti.y <= bound.y &&
		anti.x <= bound.x

	return valid, anti
}
