package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type parsingState int
const (
  searching parsingState = iota
  parsingKeyword
  keywordFound
  parsingFirstNum
  parsingSecondNum
  reset
)

var (
  state parsingState
  num1Buf = []rune{}
  num1 int
  num2Buf = []rune{}
  runningTotal int64 = 0
  keywordTemplate = []rune{
    'm',
    'u',
    'l',
  }
  maxNumSize = 3
  currBufferIndex = 0
)


func main() {
  fmt.Println("starting..")

  data, err := os.ReadFile("./input")
  if err != nil {
    fmt.Printf("error on file read: %v\n", err)
    os.Exit(1)
  }


  res, err := parse(data)
  if err != nil {
    fmt.Printf("parsing error: %v", err)
    os.Exit(1)
  }
  fmt.Printf("Result: %v", res)
}

func parse(data []byte) (int64, error) {
  res := int64(0)
  i := 0
  curd := data
  for i < len(data) {
    l, err := nextChar(curd, &res)
    if err != nil {
      return 0, fmt.Errorf("error reading char: %w", err)
    }
    if (state == reset) {
      state = searching
    } else {
      i += l
      curd = curd[l:]
    }
  }
  return res, nil
}

func nextChar(data []byte, total *int64) (int, error) {
  c, size := utf8.DecodeRune(data)
  if state == searching {
    handleSearching(c)
  } else if state == parsingKeyword {
    handleParsingKeyword(c)
  } else if state == keywordFound {
    handleKeywordFound(c)
  } else if state == parsingFirstNum {
    handleParsingFirstNum(c)
  } else if state == parsingSecondNum {
    *total += handleParsingSecondNum(c)
  }
  return size, nil
} 

func handleSearching(ch rune) {
  if ch == keywordTemplate[0] {
    state = parsingKeyword
    currBufferIndex = 1
    return
  } 
  currBufferIndex = 0
}

func handleParsingKeyword(ch rune) {
  if ch != keywordTemplate[currBufferIndex] {
    resetState()
    return
  }

  if currBufferIndex == 2 {
    state = keywordFound
    currBufferIndex = 0
    return
  }

  currBufferIndex += 1
}

func handleKeywordFound(ch rune) {
  if ch == '(' {
    state = parsingFirstNum
    currBufferIndex = 0
  } else {
    resetState()
  }
}

func handleParsingFirstNum(ch rune) {
  if unicode.IsDigit(ch) {
    num1Buf = append(num1Buf, ch)

    if currBufferIndex >= maxNumSize {
      resetState()
      return
    }
    currBufferIndex += 1
  } else if ch == ',' {
    if currBufferIndex == 0 || currBufferIndex > maxNumSize {
      resetState()
      return
    }

    num, err := strconv.Atoi(string(num1Buf))
    if err != nil {
      resetState()
    }
    num1 = num

    state = parsingSecondNum
    currBufferIndex = 0
  } else {
    resetState()
  }
}
func handleParsingSecondNum(ch rune) int64{
  if unicode.IsDigit(ch) {
    num2Buf = append(num2Buf, ch)

    if currBufferIndex >= maxNumSize {
      resetState()
      return 0
    }
    currBufferIndex += 1
  } else if ch == ')' {
    if currBufferIndex == 0 || currBufferIndex > maxNumSize {
      resetState()
      return 0
    }

    num2, err := strconv.Atoi(string(num2Buf))
    if err != nil {
      resetState()
      return 0
    }

    result := int64(num1) * int64(num2)
    resetState()
    return result
  } else {
    resetState()
  }

  return 0
}

func resetState() {
  num1Buf = []rune{}
  num2Buf = []rune{}
  currBufferIndex = 0
  state = reset
}
