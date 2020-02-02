package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
)

var w, h int
var white color.RGBA

func init() {
	w, h = ebiten.ScreenSizeInFullscreen()
	white = color.RGBA{
		byte(255),
		byte(255),
		byte(255),
		byte(0xff),
	}
}

func update(screen *ebiten.Image) error {

	screen.Set(500, 500, white)
	screen.Set(501, 501, white)
	screen.Set(502, 502, white)
	screen.Set(503, 503, white)

	if ebiten.IsDrawingSkipped() {
		return nil
	}

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
