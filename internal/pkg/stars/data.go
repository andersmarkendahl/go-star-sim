package stars

import (
	"time"
)

// Pixel is internal data of pixel
// Maximum grid size is 2^16 x 2^16
type Pixel struct {
	Px []uint16
	Py []uint16
}

// SimData contains simulation data
// Written by star-calc
// Read by star-sim
type SimData struct {
	Width  int
	Height int
	Steps  int
	Model  string
	Time   time.Duration
	Stars  []Pixel
}
