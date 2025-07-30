package main

import (
	"log"

	"Balls/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	var game game.Game = *game.EmptyGame()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
