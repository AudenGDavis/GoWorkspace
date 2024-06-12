package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	box     *ebiten.Image
	gopher  *ebiten.Image
	scaler  *ebiten.DrawImageOptions
	playerX float64
	playerY float64
)

type Game struct{}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		playerY = -4
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		playerY = 4
	} else {
		playerY = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		playerX = -4
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		playerX = 4
	} else {
		playerX = 0
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	scaler.GeoM.Translate(playerX, playerY)
	screen.DrawImage(box, scaler)
	box.DrawImage(gopher, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1000, 1000
}

func main() {
	var err error
	scaler = &ebiten.DrawImageOptions{}
	gopher, _, err = ebitenutil.NewImageFromFile("gopher.png")

	if err != nil {
		log.Fatal(err)
	}

	box = ebiten.NewImage(240, 240)
	box.Fill(color.White)

	game := &Game{}
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("Your game's title")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
