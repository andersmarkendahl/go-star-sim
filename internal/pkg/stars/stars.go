package stars

import (
	"fmt"
	"image/color"
	"math"
)

// Simulation variables
var (
	dt = 100.0
	G  = 1.0
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
	// Runtime variables set for performance
	x, y, vx, vy, dx, dy, d, d3, theta, wmid, hmid float64
	ax, ay, ax1, ax2, ax3, ax4, ay1, ay2, ay3, ay4 float64
	cstars                                         float64
	maxdepth                                       = 12
	root                                           *Quad
)

// Star is a simple position and velocity coordinate
type Star struct {
	X, Y, vx, vy float64
}

// Quad is a rectangle that contains four smaller rectangles
// Values are total mass, center of gravity and depth
type Quad struct {
	tl, tr, dl, dr *Quad
	mass, cmx, cmy float64
	depth          int
}

// StarList is a global slice of stars
var StarList []*Star

func init() {
}

// StartValues set starting position and velocity
// Fixed starting position and velocity is random
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

func buildQuadTree(sl []*Star, wmin, wmax, hmin, hmax float64, depth int) *Quad {

	var subsl []*Star
	for i := range sl {
		if sl[i].X > wmin && sl[i].X < wmax && sl[i].Y > hmin && sl[i].Y < hmax {
			subsl = append(subsl, sl[i])
		}
	}
	cstars = float64(len(subsl))

	// Quad does not contain any objects
	if cstars == 0 {
		return nil
	}

	// Quad contains single object -> leaf
	if cstars == 1 {
		return &Quad{tl: nil, tr: nil, dl: nil, dr: nil,
			mass: 1.0, cmx: subsl[0].X, cmy: subsl[0].Y, depth: depth}
	}

	// Calculate center of mass for node
	var cmx, cmy float64
	for i := range subsl {
		cmx += subsl[i].X
		cmy += subsl[i].Y
	}
	cmx /= cstars
	cmy /= cstars

	// Quad has reached max depth -> leaf
	if depth >= maxdepth {
		return &Quad{tl: nil, tr: nil, dl: nil, dr: nil,
			mass: cstars, cmx: cmx, cmy: cmy, depth: depth}
	}

	// Quad needs split
	wmid = (wmin + wmax) / 2
	hmid = (hmin + hmax) / 2
	tl := buildQuadTree(subsl, wmin, wmid, hmin, hmid, depth+1)
	tr := buildQuadTree(subsl, wmid, wmax, hmin, hmid, depth+1)
	dl := buildQuadTree(subsl, wmin, wmid, hmid, hmax, depth+1)
	dr := buildQuadTree(subsl, wmid, wmax, hmid, hmax, depth+1)
	return &Quad{tl: tl, tr: tr, dl: dl, dr: dr,
		mass: cstars, cmx: cmx, cmy: cmy, depth: depth}
}

func calcAcc(star *Star, subRoot *Quad) (float64, float64) {

	// Empty Node
	if subRoot == nil {
		return 0, 0
	}

	// Calulate distances and theta (that decides when to approximate)
	dx = subRoot.cmx - star.X
	dy = subRoot.cmy - star.Y
	d = math.Sqrt(dx*dx + dy*dy)
	theta = (fW / math.Exp2(float64(subRoot.depth))) / d

	if theta < 0.5 {
		// Treat node a single object
		d3 = d * d * d
		ax = G * subRoot.mass * dx / d3
		ay = G * subRoot.mass * dy / d3
	} else {
		// Sum up the forces from all 4 Quads
		ax1, ay1 = calcAcc(star, subRoot.tl)
		ax2, ay2 = calcAcc(star, subRoot.tr)
		ax3, ay3 = calcAcc(star, subRoot.dl)
		ax4, ay4 = calcAcc(star, subRoot.dr)
		ax = ax1 + ax2 + ax3 + ax4
		ay = ay1 + ay2 + ay3 + ay4
	}
	return ax, ay
}

// TimestepBarnesHut updates position and velocity of all stars
// Velocity update is based on Barnes Hut approximation
func TimestepBarnesHut() error {

	// Update positions of all stars based on current velocity
	for i := range StarList {
		StarList[i].X = StarList[i].X + StarList[i].vx/dt
		StarList[i].Y = StarList[i].Y + StarList[i].vy/dt
	}

	// Update velocities of all stars based approximation gravity calculation
	// Create a quadtree with all stars inserted
	root = buildQuadTree(StarList, 0, fW, 0, fH, 0)
	for i := range StarList {
		// Update Velocities
		ax, ay := CalcAcc(StarList[i], root)
		StarList[i].vx = StarList[i].vx + ax/dt
		StarList[i].vy = StarList[i].vy + ay/dt
	}
	return nil

}

// TimestepExact updates position and velocity of all stars
// Velocity update is based on exact calculation
func TimestepExact() error {

	// Update positions of all stars based on current velocity
	for i := range StarList {
		StarList[i].X = StarList[i].X + StarList[i].vx/dt
		StarList[i].Y = StarList[i].Y + StarList[i].vy/dt
	}

	// Update velocities of all stars based on "exact" gravity calculation
	for i := range StarList {
		ax = 0.0
		ay = 0.0
		// Compare star to all other stars and add up acceleration
		for j := range StarList {
			if i == j {
				continue
			}
			dx = StarList[j].X - StarList[i].X
			dy = StarList[j].Y - StarList[i].Y
			d = math.Sqrt(dx*dx + dy*dy)
			// Skip when stars are closer than 1 pixel
			if d < 1.0 {
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
