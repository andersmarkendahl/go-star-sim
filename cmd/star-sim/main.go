package main

import (
	"github.com/Aoana/go-star-sim/internal/pkg/stars"
	"github.com/hajimehoshi/ebiten"
	"log"
)

type Game struct{}

func init() {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return stars.W, stars.H
}

func (g *Game) Update(screen *ebiten.Image) error {

	var err error

	err = stars.TimestepStars()
	if err != nil {
		return err
	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw all stars
	for i := range stars.StarList {
		screen.Set(int(stars.StarList[i].X[0]), int(stars.StarList[i].X[1]), stars.White)
	}

	return nil
}

func main() {

	game := &Game{}

	// Specify the window size.
	stars.W, stars.H = ebiten.ScreenSizeInFullscreen()
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Star System")

	// Spawn all stars
	stars.StartValues()

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
