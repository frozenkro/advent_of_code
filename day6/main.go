package main

import (
	"bufio"
	"fmt"
	"os"
)
type direction byte
const (
  up direction = '^'
  right direction = '>'
  down direction = 'v'
  left direction = '<'
)
var directions = []direction{ up, right, down, left }

type location struct {
  x int
  y int
  d direction
}

func main() {
  data, err := getData("./input")
  if err != nil {
    fmt.Printf("Error occurred during data retrieval: %v\n", err)
    os.Exit(1)
  }

  result, err := totalGuardSpaces(data)
  if err != nil {
    fmt.Printf("Error occurred during calculation\n")
    os.Exit(1)
  }

  fmt.Printf("Result: %d\n", result)
}

func getData(path string) ([][]byte, error) {
  f, err := os.Open(path)
  if err != nil {
    return nil, fmt.Errorf("Error occurred when reading file at '%v': %v\n", path, err)
  }

  data := [][]byte{}
  sn := bufio.NewScanner(f)
  for sn.Scan() {
    data = append(data, []byte(sn.Text()))
  }
  return data, nil
}

func totalGuardSpaces(data [][]byte) (int, error) {
  loc, err := initCoordinates(data)
  if err != nil {
    return 0, err
  }
  var nextLoc location

  total := 0
  for {
    if data[loc.y][loc.x] != 'x' {
      total++
    }

    nextLoc = getNextLoc(*loc)
    if (isOob(data, nextLoc)) {
      return total, nil
    }

    for (data[nextLoc.y][nextLoc.x] == byte('#')) {
      nextLoc.turn()
      if (isOob(data, nextLoc)) {
        return total, nil
      }
    }

    loc = &nextLoc
  }
}

func isOob(data [][]byte, loc location) bool {
  return loc.y < 0 || loc.x < 0 || loc.y >= len(data) || loc.x >= len(data[loc.y])
}

func getNextLoc(loc location) location {
  res := loc
  if loc.d == up {
    res.y -= 1
  }
  if loc.d == right {
    res.x += 1
  }
  if loc.d == down {
    res.y += 1 
  }
  if loc.d == left {
    res.x -= 1
  }
  return res
}

func (l *location) turn() {
  for i, v := range directions {
    if l.d == v {
      l.d = directions[(i+1)%len(directions)]
    }
  }
}

func initCoordinates(data [][]byte) (*location, error) {
  for y, row := range data {
    for x, ch := range row {
      for _, d := range directions {
        if ch == byte(d) {
          return &location{ x: x, y: y, d: d }, nil
        }
      }
    }
  }
  return nil, fmt.Errorf("Initial coordinates not found")
}
