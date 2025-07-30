package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

////////////////////////////////////////////////
// Update Function
////////////////////////////////////////////////

func (game *Game) Update() error {
	//uses WASD keys to move the camera
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		game.CameraY -= 1 / game.CameraScale
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		game.CameraY += 1 / game.CameraScale
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		game.CameraX -= 1 / game.CameraScale
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		game.CameraX += 1 / game.CameraScale
	}

	_, yoff := ebiten.Wheel()
	game.CameraScale *= math.Pow(1.01, yoff)
	return nil
}
