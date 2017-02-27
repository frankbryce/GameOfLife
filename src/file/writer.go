package file

import (
  "io"
  "github.com/frankbryce/GameOfLife/types/"
)

type GolWriter struct {
  rows,cols int
}

type GolPersister interface {
  Open(string)
  Create(string)
  Append(GolFrame)
  Modify(frameNo int, frame GolFrame)
  Save(Simulation)
}

func (gw GolWriter) Write(p []byte) (n int, err error) {
  
}

func (gf GolFrame) Write() (n int, err error) {
  
}
