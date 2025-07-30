package utils

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func WorldToScreenPosition(cameraX, cameraY, x, y, w, h, z float64) (float32, float32) {
	return float32(w/2 + (x-cameraX)*z), float32(h/2 + (y-cameraY)*z)
}

func WorldToScreenPositionX(cameraX, x, w, z float64) float32 {
	return float32(w/2 + (x-cameraX)*z)
}

func WorldToScreenPositionY(cameraY, y, h, z float64) float32 {
	return float32(h/2 + (y-cameraY)*z)
}

func ScreenToWorldPositionX(cameraX, screenX, w, z float64) float32 {
	return float32((screenX-w/2)/z + cameraX)
}

func ScreenToWorldPositionY(cameraY, screenY, h, z float64) float32 {
	return float32((screenY-h/2)/z + cameraY)
}

func RenderGridLines(screen *ebiten.Image, increments float32, strokeWidth float32, CameraX float64, CameraY float64, W float64, H float64, Z float64, strokeColor color.Color) {
	var x0 float32 = ScreenToWorldPositionX(CameraX, 0, W, Z)
	var y0 float32 = ScreenToWorldPositionY(CameraY, 0, H, Z)

	var x1 float32 = ScreenToWorldPositionX(CameraX, W, W, Z)
	var y1 float32 = ScreenToWorldPositionY(CameraY, H, H, Z)

	x0 = x0 - float32(math.Mod(float64(x0), float64(increments)))
	y0 = y0 - float32(math.Mod(float64(y0), float64(increments)))

	for x := x0; x <= x1; x += increments {
		// print("l")
		vector.StrokeLine(
			screen,
			WorldToScreenPositionX(CameraX, float64(x), W, Z),
			0,
			WorldToScreenPositionX(CameraX, float64(x), W, Z),
			float32(H),
			strokeWidth*float32(Z),
			strokeColor,
			true)
	}

	for y := y0; y <= y1; y += increments {
		// print("l")
		vector.StrokeLine(
			screen,
			0,
			WorldToScreenPositionY(CameraY, float64(y), H, Z),
			float32(W),
			WorldToScreenPositionY(CameraY, float64(y), H, Z),
			strokeWidth*float32(Z),
			strokeColor,
			true)
	}
	// vector.StrokeLine(screen, x0, y0, x1, y1, 10, color.Black, false)
	// fmt.Printf("(%f,%f) -> (%f,%f)\n", x0, y0, x1, y1)

}

func RenderAxisLines(screen *ebiten.Image, strokeWidth float32, CameraX float64, CameraY float64, W float64, H float64, Z float64, strokeColor color.Color) {
	x, y := WorldToScreenPosition(CameraX, CameraY, 0, 0, W, H, Z)

	//horizontal line
	vector.StrokeLine(
		screen,
		0,
		y,
		float32(W),
		y,
		strokeWidth*float32(Z),
		strokeColor,
		true)

	//vertical line
	vector.StrokeLine(
		screen,
		x,
		0,
		x,
		float32(H),
		strokeWidth*float32(Z),
		strokeColor,
		true)
}
