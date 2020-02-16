package main

import (
	"flag"
	"fmt"
	"github.com/Aoana/go-star-sim/internal/pkg/stars"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

// Game is part of ebiten and defines the game
type Game struct{}

var timestep func()
var model *string

func init() {
}

// Layout is part of ebiten Game interface
// Defines the screen and is set to always run in full screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return stars.W, stars.H
}

// Update is part of ebiten Game interface
// Is called for every frame and executes one timestep
func (g *Game) Update(screen *ebiten.Image) error {

	timestep()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw all stars
	for i := range stars.StarList {
		screen.Set(int(stars.StarList[i].X), int(stars.StarList[i].Y), stars.White)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Stars: %d\nModel: %s\nTPS: %0.2f", len(stars.StarList), *model, ebiten.CurrentTPS()))

	return nil
}

func main() {

	game := &Game{}

	// Specify the window size.
	stars.W, stars.H = ebiten.ScreenSizeInFullscreen()
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Star System")

	// Set radius of star cluster
	nstars := flag.Int("nstars", 12, "Number of stars in cluster")
	model = flag.String("model", "Exact", "Gravity calculation model\n\"Exact\", \"BarnesHut\"")
	flag.Parse()

	switch *model {
	case "Exact":
		timestep = stars.TimestepExact
	case "BarnesHut":
		timestep = stars.TimestepBarnesHut
	default:
		log.Fatal("Unknown gravity model")
	}

	// Spawn all stars
	stars.StartValues(*nstars)

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
