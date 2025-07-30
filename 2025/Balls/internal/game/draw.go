package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"Balls/utils"
)

////////////////////////////////////////////////
// Draw Function
////////////////////////////////////////////////

func (game *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{255, 255, 255, 255})

	//render axis and grid lines
	utils.RenderGridLines(
		screen,
		1,
		0.025,
		game.CameraX,
		game.CameraY,
		float64(game.screenWidth),
		float64(game.screenHeight),
		game.CameraScale,
		color.RGBA{255, 0, 0, 255}) //red
	utils.RenderGridLines(
		screen,
		10,
		0.05,
		game.CameraX,
		game.CameraY,
		float64(game.screenWidth),
		float64(game.screenHeight),
		game.CameraScale,
		color.RGBA{0, 0, 255, 255}) //blue
	utils.RenderAxisLines(
		screen,
		0.1,
		game.CameraX,
		game.CameraY,
		float64(game.screenWidth),
		float64(game.screenHeight),
		game.CameraScale,
		color.Black)

	//iterates through all game objects and draws them
	for _, gameObject := range game.GameObjects {
		ballObject, ok := gameObject.(*utils.Ball)
		if ok {
			fmt.Printf("Drawing Ball at (%f, %f)\n", utils.WorldToScreenPositionX(game.CameraX, ballObject.X0, float64(game.screenWidth), game.CameraScale), ballObject.Y0)
			vector.DrawFilledCircle(
				screen,
				utils.WorldToScreenPositionX(game.CameraX, ballObject.X0, float64(game.screenWidth), game.CameraScale),
				utils.WorldToScreenPositionY(game.CameraY, ballObject.Y0, float64(game.screenHeight), game.CameraScale),
				float32(ballObject.Radius*game.CameraScale),
				color.RGBA{255, 0, 0, 255},
				true,
			)

		}
	}
}
