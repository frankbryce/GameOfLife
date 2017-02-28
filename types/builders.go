// Package containing the Game Of Life primitive types and factories
package types

import (
  "fmt"
  "reflect"
)

// enum constants for intialization strategies for Game of Life frames
const (
  INIT_DEAD = iota
  INIT_VALUE = iota
  INIT_RANDOM = iota
  INIT_FILE = iota
)

const (
  // Example file for testing and getting started.  See INIT_FILE option in
  // the BuildFrame() method
  ExampleFile = "http://github.com/frankbryce/GameOfLife/types/ExampleFile.golh"
)

type golCell struct {
  alive bool
}

type golFrame struct {
  cells [][]GolCell
}

func (c golCell) IsAlive() bool {
  return c.alive
}

func (c golCell) Set(a bool) {
  c.alive = a
}

func (f golFrame) checkCells() {
  if f.cells == nil {
    panic("Frame has not been initialized")
  }
}

func (f golFrame) Cell(row,col int) GolCell {
  f.checkCells()
  if row >= f.Rows() || row < 0 {
    panic(fmt.Sprintf("Row %d is invalid for a frame with only %d rows",row,f.Rows()))
  }
  if col >= f.Cols() || col < 0 {
    panic(fmt.Sprintf("Column %d is invalid for a frame with only %d colums",col,f.Cols()))
  }
  return f.cells[row][col]
}

func (f golFrame) Rows() int {
  f.checkCells()
  return len(f.cells)
}

func (f golFrame) Cols() int {
  f.checkCells()
  if (len(f.cells)==0) {
    return 0
  }
  if f.cells[0] == nil {
    panic("Frame has not been initialized properly, there is an internal error")
  }
  return len(f.cells[0])
}

// Build a frame, with a given initialization strategy
//
// INIT_VALUE: Initialize all cells with the alive value given
// Example: Initialize a 100x100 frame with all cells dead
//     types.BuildFrame(types.INIT_VALUE, 100, 100, false)
//
// INIT_RANDOM: Initialize all cells with a random value, with an alive
// probability given
// Example: 100x100 frame where ~30% of cells are alive
//     types.BuildFrame(types.INIT_RANDOM, 100, 100, 0.3)
//
// INIT_FILE: Initialize the board from the given file
// Example: Build the frame from the example file given in this package
//     types.BuildFrame(types.INIT_FILE, types.ExampleFile)
func BuildFrame(init_strategy int, a ...interface{}) GolFrame {
  switch init_strategy {
  case INIT_VALUE:
    checkForNArgs("INIT_VALUE", 3, a...)
    rows, cols := checkForRowsAndCols(a...)
    isAlive, ok := a[2].(bool)
    if !ok {
      panic(fmt.Sprintf("Type of value to initialize frame should be boolean, actual type was %s", reflect.TypeOf(a[2])))
    }
    return buildFrameValue(rows, cols, isAlive)
  case INIT_RANDOM:
    checkForNArgs("INIT_RANDOM", 3, a...)
    rows, cols := checkForRowsAndCols(a...)
    pAlive, ok := a[2].(float64)
    if !ok {
      panic(fmt.Sprintf("Type of parameter for alive percentage should be float64, actual type is %s", reflect.TypeOf(a[2])))
    }
    return buildFrameRandom(rows, cols, pAlive)
  case INIT_FILE:
    checkForNArgs("INIT_FILE", 1, a...)
    file, ok := a[0].(string)
    if !ok {
      panic(fmt.Sprintf("Type of parameter for file should be string, actual type is %s", reflect.TypeOf(a[0])))
    }
    return buildFrameFile(file)
  default:
    panic(fmt.Sprintf("Invalid initialization strategy: %d",init_strategy))
  }
}

func checkForNArgs(strategy string, n int, a ...interface{}) {
  if len(a) != n {
    panic(fmt.Sprintf("BuildFrame() called with init strategy %s should have %d additioanl arguments.  Actual number of arguments was %d", strategy, n, len(a)))
  }
}

func checkForRowsAndCols(a ...interface{}) (rows int, cols int) {
  rows, ok1 := a[0].(int)
  cols, ok2 := a[1].(int)
  if !ok1 || !ok2 {
    panic(fmt.Sprintf("BuildFrame() should have integers for rows and cols arguments.  Actual types were %s and %s", reflect.TypeOf(a[0]), reflect.TypeOf(a[1])))
  }
  return rows, cols
}

func buildFrameValue(rows int, cols int, isAlive bool) golFrame {
  f := golFrame{}
  f.cells = make([][]GolCell, rows)
  for r:=0;r<rows;r++ {
    f.cells[r] = make([]GolCell, cols)
    for c:=0;c<cols;c++ {
      f.cells[r][c] = golCell{isAlive}
    }
  }
  return f
}

func buildFrameRandom(rows int, cols int, pAlive float64) golFrame {
  return golFrame{}
}

func buildFrameFile(file string) golFrame {
  return golFrame{}
}

