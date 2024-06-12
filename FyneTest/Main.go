package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fyne Test")

	w.Resize(fyne.NewSize(500, 500))

	label := canvas.NewText("text", fyne.NewColor(0, 0, 0))

	label.Move(fyne.NewPos(100, 100))
	fmt.Println(label.Position())

	w.SetContent(label)

	w.ShowAndRun()

}
