package stars

import (
	"fmt"
	"image/color"
	"math"
)

// Simulation variables
var (
	dt = 100.0
	G  = 0.1
	// Width
	W int
	// Height
	H int
	// The color white
	White = color.RGBA{
		byte(255),
		byte(255),
		byte(255),
		byte(0xff),
	}
	x, y, vx, vy, ax, ay, dx, dy, d, d3 float64
)

// Star is a simple position and velocity coordinate
type Star struct {
	X, Y, vx, vy float64
}

// StarList is a global slice of stars
var StarList []*Star

func init() {
}

// StartValues set starting position and velocity
// Fixed starting position and velocity is random
func StartValues(nstars int) error {

	r := math.Round(math.Sqrt(float64(nstars) / math.Pi))
	tx := float64(W / 2)
	ty := float64(H / 2)

	for i := -r; i <= r; i++ {
		for j := -r; j <= r; j++ {
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
				} else if x > 0 {
					vx = y / x
					vy = -x
				} else {
					vx = -y / x
					vy = -x
				}

				// Velocity vector with fixed length
				d := math.Sqrt(vx*vx + vy*vy)
				vxs := 20 * vx / d
				vys := 20 * vy / d

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

// TimestepStars updates position and velocity of all stars
func TimestepStars() error {

	// Update positions of all stars based on current velocity
	for i := range StarList {
		StarList[i].X = StarList[i].X + StarList[i].vx/dt
		StarList[i].Y = StarList[i].Y + StarList[i].vy/dt
	}

	// Update velocities of all stars based on gravity calculation
	for i := range StarList {

		ax = 0.0
		ay = 0.0
		for j := range StarList {
			if i == j {
				continue
			}
			dx = StarList[j].X - StarList[i].X
			dy = StarList[j].Y - StarList[i].Y
			d = math.Sqrt(dx*dx + dy*dy)
			if d < 2.0 {
				continue
			}
			d3 = d * d * d
			ax += G * dx / d3
			ay += G * dy / d3
		}

		StarList[i].vx = StarList[i].vx + ax/dt
		StarList[i].vy = StarList[i].vy + ay/dt
	}
	return nil

}
