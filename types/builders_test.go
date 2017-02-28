// Package containing the Game Of Life primitive types and factories
package types

import (
  "fmt"
  "testing"
)

func TestBuildFrameValue(t *testing.T) {
  fmt.Println("Test Started")
  f := BuildFrame(INIT_VALUE, 10, 20, false)
  if f.Rows() != 10 {
    t.Fatal(fmt.Sprintf("There should be 10 rows, but there are %d", f.Rows()))
  }
  if f.Cols() != 20 {
    t.Fatal(fmt.Sprintf("There should be 20 cols, but there are %d", f.Cols()))
  }
  for r:=0;r<f.Rows();r++ {
    for c:=0;c<f.Cols();c++ {
      if f.Cell(r,c).IsAlive() {
        t.Fatal("All cells should be dead, but one was alive")
      }
    }
  }
  fmt.Println("Test Completed")
}
