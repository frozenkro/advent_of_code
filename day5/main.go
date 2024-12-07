package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
  first int
  last int
}

func main() {
  rules, rows, err := getData("./input")
  if err != nil {
    fmt.Printf("Error occurred when getting input data: %v\n", err)
  }

  res := GetSumValidRows(rows, rules)
  fmt.Printf("Result: %v", res)
}

// Returns sum of middle values of all valid rows
func GetSumValidRows(rows [][]int, rules []rule) int {
  sum := 0
  for _, v := range rows {
    sum += GetValueForRow(v, rules)
  }
  return sum
}

// Returns 0 for invalid update row
func GetValueForRow(row []int, rules []rule) int {
  for i, _ := range row {
    if IsAnyRuleBroken(row, rules, i) {
      return 0
    }
  }
  return row[(len(row)-1)/2]
}

func IsAnyRuleBroken(row []int, rules []rule, idx int) bool {
  for _, rule := range rules {
    if rule.first == row[idx] && IsRuleBroken(row, rule, idx) {
      return true
    }
  }
  return false
}

func IsRuleBroken(row []int, rule rule, firstNumIdx int) bool {
  for i := 0; i < firstNumIdx; i++ {
    if row[i] == rule.last {
      return true
    }
  }
  return false
}

func getData(filePath string) ([]rule, [][]int, error) {
  f, err := os.Open(filePath)
  if err != nil {
    return nil, nil, fmt.Errorf("Error reading file: %w\n", err)
  }
  sc := bufio.NewScanner(f)

  line := 0
  rules := []rule{}
  for sc.Scan() {
    line++
    ruleArr := strings.Split(sc.Text(), "|")
    if len(ruleArr) == 1 {
      break
    } else if len(ruleArr) != 2 {
      return nil, nil, fmt.Errorf("Unexpected rule entry length on line %v: %v\n", line, len(ruleArr))
    }
    first, e1 := strconv.Atoi(ruleArr[0])
    last, e2 := strconv.Atoi(ruleArr[1])
    if e1 != nil || e2 != nil {
      return nil, nil, fmt.Errorf(
        "Error(s) occurred while parsing rule to ints on line %v: \n%w\n\n%w\n", 
        line, e1, e2)
    }

    r := rule{ first: first, last: last }
    rules = append(rules, r)
  }

  updates := [][]int{}
  for sc.Scan() {
    line++
    stArr := strings.Split(sc.Text(), ",")
    if len(stArr) % 2 != 1 {
      return nil, nil, fmt.Errorf("Expected all updates to be odd length, line %v was even", line)
    }

    pages := []int{}
    for i, v := range stArr {
      p, err := strconv.Atoi(v)
      if err != nil {
        return nil, nil, fmt.Errorf(
          "Error occurred parsing page num; line %v, idx %v, value %v\nError: %w\n",
          line, i, v, err)
      }

      pages = append(pages, p)
    }

    if len(pages) > 0 {
      updates = append(updates, pages)
    }
  }
  return rules, updates, nil
}


