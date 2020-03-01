package stars

import (
	"fmt"
	"image/color"
	"math"
)

// Simulation variables
var (
	dt = 100.0
	G  = 10.0
	// Width
	W int
	// Height
	H      int
	fW, fH float64
	// The color white
	White = color.RGBA{
		byte(255),
		byte(255),
		byte(255),
		byte(0xff),
	}
)

// Star is a simple position and velocity coordinate
type Star struct {
	X, Y, vx, vy float64
}

// StarList is a global slice of stars
var StarList []*Star

// StartValues set starting position and velocity
// Forms an ellipse shape with velocities approx tangential
func StartValues(nstars int) error {

	fW = float64(W)
	fH = float64(H)

	rx := 10 * math.Round(math.Sqrt(float64(nstars)/math.Pi))
	ry := 0.5 * rx
	tx := fW / 2
	ty := fH / 2
	for i := -rx; i <= rx; i += 10 {
		for j := -ry; j <= ry; j += 5 {
			if i*i/(rx*rx)+j*j/(ry*ry) <= 1 {

				// Logical starting position
				x := float64(i)
				y := float64(j)

				var vx, vy float64

				// Velocity perpendicular to circle
				if x == 0 && y == 0 {
					continue
				} else if x == 0 {
					vx = y
					vy = 0.0
				} else if x > 0 {
					vx = y / x
					vy = -x
				} else {
					vx = -y / x
					vy = -x
				}

				// Velocity vector with fixed length
				d := math.Sqrt(vx*vx + vy*vy)
				vxs := 6.0 * vx / d
				vys := 6.0 * vy / d

				// Translate position to middle of screen
				x += tx
				y += ty

				// Construct object
				s := Star{X: x, Y: y, vx: vxs, vy: vys}
				StarList = append(StarList, &s)
			}
		}
	}
	fmt.Printf("Requested number of stars %d resulted in %d stars\n", nstars, len(StarList))
	return nil
}
