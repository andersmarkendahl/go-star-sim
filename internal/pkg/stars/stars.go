package stars

import (
	"math"
	"fmt"
	"image/color"
)

// Simulation variables
// dt is set to 1
const (
	// Gravitational constant
	G = 0.05
)

var (
	// Data is storage for simulation data
	Data SimData
	// Float representiation of grid
	fW, fH float64
	// White is ... the color white
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
// Forms a circle with start velocities approx tangential counterclockwise
func StartValues(nstars int) (int error) {

	fW = float64(Data.Width)
	fH = float64(Data.Height)

	r := 3 * math.Round(math.Sqrt(float64(nstars)/math.Pi))
	tx := fW / 2
	ty := fH / 2
	for i := -r; i <= r; i += 3 {
		for j := -r; j <= r; j += 3 {
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
				} else if y > 0 {
					vx = 1.0
					vy = -x / y
				} else if y < 0 {
					vx = -1.0
					vy = x / y
				} else {
					fmt.Println("ERROR: Start position bug", x, y)
				}

				// Velocity vector with fixed length
				d := math.Sqrt(vx*vx + vy*vy)
				vxs := Data.Velocity * vx / d
				vys := Data.Velocity * vy / d

				// Translate position to middle of grid
				x += tx
				y += ty

				// Construct object
				s := Star{X: x, Y: y, vx: vxs, vy: vys}
				StarList = append(StarList, &s)
			}
		}
	}
	return nil
}

// Move updates position for all stars
func Move() {

	// Update positions of all stars based on current velocity
	for i := range StarList {
		StarList[i].X = StarList[i].X + StarList[i].vx
		StarList[i].Y = StarList[i].Y + StarList[i].vy
	}

}
