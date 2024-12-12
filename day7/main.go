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
	data, err := getData("./input")
	if err != nil {
		fmt.Printf("Error occurred while processing input data: %v", err)
		os.Exit(1)
	}

	result := validSum(data, false)
	fmt.Printf("Result 1: '%d'", result)

	result = validSum(data, true)
	fmt.Printf("Result 2: '%d'", result)
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

func validSum(equations []equation, concatOperator bool) int {
	total := 0
	for _, e := range equations {
		if e.isValid(concatOperator) {
			total += e.total
		}
	}
	return total
}

func (e equation) isValid(concatOperator bool) bool {
	return e.check(e.operands, concatOperator)
}

func (e equation) check(ops []int, concatOperator bool) bool {
	if ops[0] > e.total {
		return false
	}
	if len(ops) == 1 {
		return ops[0] == e.total
	}

	mul := ops[0] * ops[1]
	mOps := make([]int, len(ops)-1)
	copy(mOps, ops[1:])
	mOps[0] = mul

	sum := ops[0] + ops[1]
	aOps := make([]int, len(ops)-1)
	copy(aOps, ops[1:])
	aOps[0] = sum

	if concatOperator {
		catS := strconv.Itoa(ops[0]) + strconv.Itoa(ops[1])
		cOps := make([]int, len(ops)-1)
		copy(cOps, ops[1:])
		cat, err := strconv.Atoi(catS)
		cOps[0] = cat
		if err != nil {
			panic(err)
		}
		return e.check(mOps, concatOperator) ||
			e.check(aOps, concatOperator) ||
			e.check(cOps, concatOperator)
	}

	return e.check(mOps, concatOperator) ||
		e.check(aOps, concatOperator)

}

