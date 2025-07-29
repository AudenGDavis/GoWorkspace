package utils

//	Test
// var x float32
// var y float32

// x, y = utils.WorldToScreenPosition(996.0, 1170.0, 790.00, .0, 1920.0, 1080.0, 2.4)
// fmt.Printf("%f, %f", x, y)

func WorldToScreenPosition(cameraX float32, cameraY float32, x float32, y float32, w float32, h float32, z float32) (float32, float32) {
	return (w/2 + (x-cameraX)*z), (h/2 + (y-cameraY)*z)
}

func WorldToScreenPositionX(cameraX float32, x float32, w float32, z float32) float32 {
	return (w/2 + (x-cameraX)*z)
}

func WorldToScreenPositionY(cameraY float32, y float32, h float32, z float32) float32 {
	return (h/2 + (y-cameraY)*z)
}

func ScreenToWorldPositionX(cameraX, screenX, w, z float32) float32 {
	return (screenX-w/2)/z + cameraX
}

func ScreenToWorldPositionY(cameraY, screenY, h, z float32) float32 {
	return (screenY-h/2)/z + cameraY
}
