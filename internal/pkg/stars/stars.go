package stars

import (
	"fmt"
	"image/color"
	"math"
)

// Simulation variables
// Note: dt is set to 1
const (
	G = 0.05
	V0 = 0.5
)

var (
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

	r := 5 * math.Round(math.Sqrt(float64(nstars)/math.Pi))
	tx := fW / 2
	ty := fH / 2
	for i := -r; i <= r; i += 5 {
		for j := -r; j <= r; j += 5 {
			if i*i+j*j <= r*r {

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
				} else if y == 0 {
					vx = 0.0
					vy = -x
				} else if x > 0 && y > 0 {
					vx = y / x
					vy = -y / x
				} else if x > 0 && y < 0 {
					vx = y / x
					vy = y / x
				} else if x < 0 && y > 0 {
					vx = -y / x
					vy = -y / x
				} else if x < 0 && y < 0 {
					vx = -y / x
					vy = y / x
				} else {
					fmt.Println("ERROR: Start position bug", x, y)
				}

				// Velocity vector with fixed length
				d := math.Sqrt(vx*vx + vy*vy)
				vxs := V0 * vx / d
				vys := V0 * vy / d

				// Translate position to middle of screen
				x += tx
				y += ty

				// Construct object
				s := Star{X: x, Y: y, vx: vxs, vy: vys}
				StarList = append(StarList, &s)
			}
		}
	}

	//for i := range StarList {
	//	fmt.Println("Star ", i, StarList[i])
	//}
	fmt.Printf("Requested number of stars %d resulted in %d stars\n", nstars, len(StarList))
	return nil
}
