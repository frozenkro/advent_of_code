package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	total    int
	operands []int
}

func main() {

}

func getData(path string) ([]equation, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	results := []equation{}
	sn := bufio.NewScanner(f)
	for sn.Scan() {
		nums := strings.Split(sn.Text(), " ")
		if len(nums) < 2 {
			return nil, fmt.Errorf("Invalid row found: '%v'", sn.Text())
		}

		totalStr := strings.Trim(nums[0], ":")
		total, err := strconv.Atoi(totalStr)
		if err != nil {
			return nil, fmt.Errorf("Error converting total to int for row '%v', value: '%v'", sn.Text(), totalStr)
		}

		o := []int{}
		for _, v := range nums[1:] {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("Error converting operand to int for row '%v', value: '%v'", sn.Text(), n)
			}
			o = append(o, n)
		}

		equation := equation{total: total, operands: o}
		results = append(results, equation)
	}

	return results, nil
}

func validCount(equations []equation) int {
	total := 0
	for _, e := range equations {
		if e.isValid() {
			total++
		}
	}
	return total
}

func (e equation) isValid() bool {
	 return e.check(e.operators)
}

func (e equation)check(ops []int) bool {
  if ops[0] > e.total {
    return false
  }
  if len(ops) == 1 {
    return ops[0] == e.total
  }

  rMul := ops[0] * ops[1]
  rAdd := ops[0] + ops[1]
  
  return e.check(append([]int{rMul}, ops[1:]...)) || e.check(append([]int{rAdd}, ops[1:]...))

}