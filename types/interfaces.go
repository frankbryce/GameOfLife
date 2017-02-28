// Package containing the Game Of Life primitive types and factories
package types

type GolCell interface {
  Set(alive bool)
  IsAlive() bool
}

type GolFrame interface {
  Cell(row,col int) GolCell
  Rows() int
  Cols() int
}


