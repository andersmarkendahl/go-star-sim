package stars

import (
	"fmt"
	"github.com/Aoana/ball-sim-go/pkg/mathutil"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"image/color"
)

// Simulation variables
var (
	dt = 1.0
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
)

// StarList is a global slice of objects
var StarList []*objects.Object

func init() {
}

// StartValues set starting position and velocity
// Fixed starting position and velocity is random
func StartValues(r int) error {

	for i := -r; i <= r; i++ {
		for j := -r; j <= r; j++ {
			if i*i+j*j <= r*r {
				// Random starting velocity
				vx0, _ := mathutil.RandInRange(-0.5, 0.5)
				vy0, _ := mathutil.RandInRange(-0.5, 0.5)
				// Construct objects
				s, err := objects.New(float64(W/2+i), float64(H/2+j), vx0, vy0)
				if err != nil {
					return err
				}
				StarList = append(StarList, s)
			}
		}
	}
	fmt.Println("Number of stars: ", len(StarList))
	return nil
}

// TimestepStars updates position and velocity of all stars
func TimestepStars() error {

	// Update positions of all stars based on current velocity
	for i := range StarList {
		err := StarList[i].Position(dt)
		if err != nil {
			return err
		}
	}

	// Update velocities of all stars based on gravity calculation
	for i := range StarList {
		err := StarList[i].Velocity(0, 0, dt)
		if err != nil {
			return err
		}
	}
	return nil

}
