package utils

////////////////////////////////////////////////////////////////
// Wall Struct (implements Gameobject Interface)
////////////////////////////////////////////////////////////////

type Wall struct {
	X0 float64
	Y0 float64

	X1 float64
	Y1 float64

	//Physics fields
	Vx float64
	Vy float64

	//type of wall
	Type WallType
}

func (wall *Wall) Dx() float64 {
	return wall.X1 - wall.X0
}

func (wall *Wall) Dy() float64 {
	return wall.Y1 - wall.Y0
}

func (wall *Wall) X() float64 { //Gameobject method
	return wall.X0
}

func (wall *Wall) Y() float64 { //Gameobject method
	return wall.Y0
}

////////////////////////////////////////////////////////////////
// Ball Struct (implements Gameobject Interface)
////////////////////////////////////////////////////////////////

type Ball struct {
	X0 float64
	Y0 float64

	Radius float64

	//Physics fields
	Vx float64
	Vy float64

	//Type of ball
	Type BallType
}

func (ball *Ball) X() float64 { //Gameobject method
	return ball.X0
}

func (ball *Ball) Y() float64 { //Gameobject method
	return ball.Y0
}
