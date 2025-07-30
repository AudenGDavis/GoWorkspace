package game

import (
	"Balls/utils"
)

//////////////////////////////////////////////////////////////////////////
// Game Struct
//////////////////////////////////////////////////////////////////////////

type Game struct {
	//camera x and y as float64
	CameraX      float64
	CameraY      float64
	CameraScale  float64
	screenWidth  int
	screenHeight int

	GameObjects []utils.GameObject
}

// initializes a new Game instance with default/empty values
func EmptyGame() *Game {
	return &Game{
		CameraX:     0,
		CameraY:     0,
		CameraScale: 80,
		GameObjects: []utils.GameObject{},
	}
}

// intializes a new Game instance with default camera values but 3 ball objects
func EmptyGameWithBalls() *Game {
	game := EmptyGame()
	game.GameObjects = append(game.GameObjects, &utils.Ball{X0: 10, Y0: 10, Radius: 5, Vx: 1, Vy: 1})
	game.GameObjects = append(game.GameObjects, &utils.Ball{X0: 20, Y0: 20, Radius: 5, Vx: -1, Vy: -1})
	game.GameObjects = append(game.GameObjects, &utils.Ball{X0: 30, Y0: 30, Radius: 5, Vx: 2, Vy: 2})
	return game
}
