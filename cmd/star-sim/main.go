package main

import (
	"fmt"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
)

type Game struct{}

var dt float64
var w, h int
var white color.RGBA
var star *objects.Object

func init() {
	dt = 1
	white = color.RGBA{
		byte(255),
		byte(255),
		byte(255),
		byte(0xff),
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return w, h
}

func (g *Game) Update(screen *ebiten.Image) error {

	var err error

	// Update velocities of all stars based on gravity calculation
	err = star.Velocity(0, 0.0, dt)
	if err != nil {
		return fmt.Errorf("Velocity update failed: %+v", star)

	}

	// Update position of all stars based on new velocities
	err = star.Position(dt)
	if err != nil {
		return fmt.Errorf("Position update failed: %+v", star)

	}

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Set(int(star.X[0]), int(star.X[1]), white)

	return nil
}

func main() {

	game := &Game{}

	// Specify the window size.
	w, h = ebiten.ScreenSizeInFullscreen()
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Star System")

	star, _ = objects.New(float64(w/2), float64(h/2), 0.5, 0.5)
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
