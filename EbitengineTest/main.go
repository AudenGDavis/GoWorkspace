package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1680
	screenHeight = 1050
)

type Game struct {
}

var (
	pieces []ebiten.Image
)

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Clear the screen

	var randInt int = rand.Intn(10)

	if randInt == 0 {
		screen.Fill(color.White)
	} else {
		screen.Fill(color.Black)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	image1 := ebiten.NewImage(100, 100)
	image2 := ebiten.NewImage(100, 100)
	pieces = append(pieces, image1, image2)

	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("Drawing Shapes with Ebiten")

	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
