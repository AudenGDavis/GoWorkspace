package utils

type Wall struct {
	X0 float64
	Y0 float64

	X1 float64
	Y1 float64
}

func (wall *Wall) Dx() float64 {
	return wall.X1 - wall.X0
}

func (wall *Wall) Dy() float64 {
	return wall.Y1 - wall.Y0
}

type Ball struct {
	X0 float64
	Y0 float64

	Radius float64
	Type   BallType
}
