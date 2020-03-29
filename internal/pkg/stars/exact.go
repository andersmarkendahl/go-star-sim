package stars

import (
	"math"
	"sync"
)

// Runtime variables set for performance
var exDx, exDy, exD, exD3, exAx, exAy float64

// VelocityExact updates position and velocity of all stars
// Velocity update is based on exact calculation
func VelocityExact() {
	// Update velocities of all stars based on "exact" gravity calculation
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
			if exD < 0.5 {
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

// Helper function to be able to use go routines
func exactAcc(star *Star, i int, wg *sync.WaitGroup) {
	defer wg.Done()

	var ax, ay float64
	for j := range StarList {
		if i == j {
			continue
		}
		dx := StarList[j].X - star.X
		dy := StarList[j].Y - star.Y
		d := math.Sqrt(dx*dx + dy*dy)
		// Skip when stars are too close
		if d < 0.5 {
			continue
		}
		d3 := d * d * d
		ax += G * dx / d3
		ay += G * dy / d3
	}
	star.vx = star.vx + ax
	star.vy = star.vy + ay
}

// VelocityExactGR updates position and velocity of all stars
// Velocity update is based on exact calculation with go routines
func VelocityExactGR() {
	// Update velocities of all stars based on "exact" gravity calculation
	var wg sync.WaitGroup
	for i := range StarList {
		// Compare star to all other stars and add up acceleration
		wg.Add(1)
		go exactAcc(StarList[i], i, &wg)
	}
	wg.Wait()
}
