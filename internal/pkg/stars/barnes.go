package stars

import (
	"math"
	"sync"
)

// Simulation variables
const maxdepth = 12

var root *Quad

// Quad is a rectangle that contains four smaller rectangles
// Values are total mass, center of gravity and depth
type Quad struct {
	tl, tr, dl, dr *Quad
	mass, cmx, cmy float64
	depth          int
}

func buildQuadTree(sl []*Star, wmin, wmax, hmin, hmax float64, depth int) *Quad {

	var subsl []*Star
	for i := range sl {
		if sl[i].X > wmin && sl[i].X < wmax && sl[i].Y > hmin && sl[i].Y < hmax {
			subsl = append(subsl, sl[i])
		}
	}
	cstars := float64(len(subsl))

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
	wmid := (wmin + wmax) / 2
	hmid := (hmin + hmax) / 2
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
	Dx := subRoot.cmx - star.X
	Dy := subRoot.cmy - star.Y
	D := math.Sqrt(Dx*Dx + Dy*Dy)
	// Star is same or too close
	if D < 1.0 {
		return 0, 0
	}
	theta := (fW / math.Exp2(float64(subRoot.depth))) / D

	var Ax, Ay float64
	if subRoot.mass == 1.0 || theta < 0.5 {
		// Treat node a single object (or it is a single object)
		D3 := D * D * D
		Ax = G * subRoot.mass * Dx / D3
		Ay = G * subRoot.mass * Dy / D3
	} else {
		// Sum up the forces from all 4 Quads
		Ax1, Ay1 := calcAcc(star, subRoot.tl)
		Ax2, Ay2 := calcAcc(star, subRoot.tr)
		Ax3, Ay3 := calcAcc(star, subRoot.dl)
		Ax4, Ay4 := calcAcc(star, subRoot.dr)
		Ax = Ax1 + Ax2 + Ax3 + Ax4
		Ay = Ay1 + Ay2 + Ay3 + Ay4
	}
	return Ax, Ay
}

// TimestepBarnesHut updates position and velocity of all stars
// Velocity update is based on Barnes Hut approximation
func TimestepBarnesHut() {

	// Update positions of all stars based on current velocity
	for i := range StarList {
		StarList[i].X = StarList[i].X + StarList[i].vx
		StarList[i].Y = StarList[i].Y + StarList[i].vy
	}

	// Update velocities of all stars based approximation gravity calculation
	// Create a quadtree with all stars inserted
	root = buildQuadTree(StarList, 0, fW, 0, fH, 0)
	for i := range StarList {
		// Update Velocities
		ax, ay := calcAcc(StarList[i], root)
		StarList[i].vx = StarList[i].vx + ax
		StarList[i].vy = StarList[i].vy + ay
	}
}

// Helper function to be able to use go routines
func updateAcc(star *Star, root *Quad, wg *sync.WaitGroup) {
	defer wg.Done()
	ax, ay := calcAcc(star, root)
	star.vx = star.vx + ax
	star.vy = star.vy + ay
}

// TimestepBarnesHutGR updates position and velocity of all stars
// Velocity update is based on Barnes Hut approximation with go-routines
func TimestepBarnesHutGR() {

	// Update positions of all stars based on current velocity
	for i := range StarList {
		StarList[i].X = StarList[i].X + StarList[i].vx
		StarList[i].Y = StarList[i].Y + StarList[i].vy
	}

	// Update velocities of all stars based approximation gravity calculation
	// Create a quadtree with all stars inserted
	root = buildQuadTree(StarList, 0, fW, 0, fH, 0)
	var wg sync.WaitGroup
	for i := range StarList {
		// Update Velocities
		wg.Add(1)
		go updateAcc(StarList[i], root, &wg)
	}
	wg.Wait()
}
