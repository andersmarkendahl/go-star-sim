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

	var err error

	err = stars.TimestepBarnesHut()
	if err != nil {
		return err
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw all stars
	for i := range stars.StarList {
		screen.Set(int(stars.StarList[i].X), int(stars.StarList[i].Y), stars.White)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Stars: %d\nTPS: %0.2f", len(stars.StarList), ebiten.CurrentTPS()))

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
	flag.Parse()

	// Spawn all stars
	stars.StartValues(*nstars)

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
