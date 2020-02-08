package main

import (
	"fmt"
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
)

var dt float64
var w, h int
var white color.RGBA
var star *objects.Object

func init() {
	w, h = ebiten.ScreenSizeInFullscreen()
	dt = 1
	white = color.RGBA{
		byte(255),
		byte(255),
		byte(255),
		byte(0xff),
	}
	star, _ = objects.New(float64(w/2), float64(h/2), 0.5, 0.5)
}

func update(screen *ebiten.Image) error {

	var err error

	// Update velocities of all stars based on gravity calculation
	err = star.Velocity(0, 0.1, dt)
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

	ebiten.SetFullscreen(true)

	s := ebiten.DeviceScaleFactor()
	if err := ebiten.Run(update, int(float64(w)*s), int(float64(h)*s), 1/s, "Star System"); err != nil {
		log.Fatal(err)
	}

}
