package stars

import (
	"math"
)

// Runtime variables set for performance
var exDx, exDy, exD, exD3, exAx, exAy float64

// TimestepExact updates position and velocity of all stars
// Velocity update is based on ex calculation
func TimestepExact() {

	// Update positions of all stars based on current velocity
	for i := range StarList {
		StarList[i].X = StarList[i].X + StarList[i].vx
		StarList[i].Y = StarList[i].Y + StarList[i].vy
	}

	// Update velocities of all stars based on "ex" gravity calculation
	for i := range StarList {
		exAx = 0.0
		exAy = 0.0
		// Compare star to all other stars and add up acceleration
		for j := range StarList {
			if i == j {
				continue
			}
			exDx = StarList[j].X - StarList[i].X
			exDy = StarList[j].Y - StarList[i].Y
			exD = math.Sqrt(exDx*exDx + exDy*exDy)
			// Skip when stars are closer than 1 pixel
			if exD < 1.0 {
				continue
			}
			exD3 = exD * exD * exD
			exAx += G * exDx / exD3
			exAy += G * exDy / exD3
		}
		StarList[i].vx = StarList[i].vx + exAx
		StarList[i].vy = StarList[i].vy + exAy
	}
}
