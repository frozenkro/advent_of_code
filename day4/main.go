package main

import (
	"bufio"
	"fmt"
	"os"
)

type State struct {
  matrix [][]byte
  x int
  y int
}
func (s State) Value() byte {
  return s.matrix[s.y][s.x]
}

const word = "XMAS"
var dirs = [][]int{
  []int{ 0, 1 },
  []int{ 1, 1 },
  []int{ 1, 0 },
  []int{ 1, -1 },
  []int{ 0, -1 },
  []int{ -1, -1 },
  []int{ -1, 0 },
  []int{ -1, 1 },
}

func main() {
  fmt.Print("Starting..\n")

  mtx, err := Matrix()
  if err != nil {
    fmt.Printf("Error reading matrix from file: %v", err)
    os.Exit(1)
  }

  res := CountInMatrix(word, mtx)
  if err != nil {
    fmt.Printf("Error during matrix traversal: %v", err)
    os.Exit(1)
  }

  fmt.Printf("Result: %v", res)
}

func Matrix() ([][]byte, error) {
  file, err := os.Open("./input")
  if err != nil {
    return nil, err
  }
  defer file.Close()

  mtx := [][]byte{}
  s := bufio.NewScanner(file)
  for s.Scan() {
    mtx = append(mtx, []byte(s.Text()))
  }
  return mtx, nil
}

func CountInMatrix(word string, matrix [][]byte) int {
  total := 0
  state := State{ matrix: matrix }
  for y, row := range matrix {
    state.y = y
    for x, ch := range row {
      state.x = x
      if ch == []byte(word)[0] {
        total += CountForLocation(&state, word)
      }
    }
  }
  return total
}

func CountForLocation(state *State, word string) int {
  ct := 0

  for _, v := range dirs {
    if CheckWord(state, word, v[0], v[1]) {
      ct++
    }
  }

  return ct
}

func CheckWord(state *State, word string, xMlt int, yMlt int) bool {
  ascii_boys := []byte(word)

  // Starting at the end catches index overflows in first iteration
  // also no need to check 0 index; already confirmed by this point
  for i := len(ascii_boys)-1; i > 0; i-- {
    x := (state.x + i) * xMlt
    y := (state.y + i) * yMlt
    if x < 0 || 
    y < 0 || 
    y >= len(state.matrix) || 
    x >= len(state.matrix[y]) || 
    state.matrix[y][x] != ascii_boys[i] {
      return false
    }
  }
  return true
}
