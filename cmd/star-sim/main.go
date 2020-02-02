package main

import (
	"github.com/Aoana/ball-sim-go/pkg/objects"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
)

var w, h int
var white color.RGBA
var star *objects.Object

func init() {
	w, h = ebiten.ScreenSizeInFullscreen()
	white = color.RGBA{
		byte(255),
		byte(255),
		byte(255),
		byte(0xff),
	}
	star, _ = objects.New(500, 500, 10, 10)
}

func update(screen *ebiten.Image) error {

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	screen.Set(int(star.X[0]), int(star.X[1]), white)

	return nil
}

func main() {

	ebiten.SetFullscreen(true)

	w, h := ebiten.ScreenSizeInFullscreen()
	// On mobiles, ebiten.MonitorSize is not available so far.
	// Use arbitrary values.
	if w == 0 || h == 0 {
		w = 300
		h = 450
	}

	s := ebiten.DeviceScaleFactor()
	if err := ebiten.Run(update, int(float64(w)*s), int(float64(h)*s), 1/s, "Star System"); err != nil {
		log.Fatal(err)
	}

}
