package stars

import (
	"github.com/Aoana/ball-sim-go/pkg/mathutil"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"image/color"
)

// Simulation variables
var (
	dt float64
	// Width
	W int
	// Height
	H int
	// The color white
	White color.RGBA
)

// StarList is a global slice of objects
var StarList []*objects.Object

func init() {
	// Timestep difference
	dt = 1
	// The color white
	White = color.RGBA{
		byte(255),
		byte(255),
		byte(255),
		byte(0xff),
	}
}

// StartValues set starting position and velocity
// Fixed starting position and velocity is random
func StartValues() error {

	for i := 0; i < 50; i++ {
		// Random starting velocity
		vx0, _ := mathutil.RandInRange(0, 0.5)
		vy0, _ := mathutil.RandInRange(0, 0.5)
		// Ball constructor
		s, err := objects.New(float64(W/2+2*i), float64(H/2+2*i), vx0, vy0)
		if err != nil {
			return err
		}
		StarList = append(StarList, s)
	}
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
