package main

import (
	"maze/world"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Rect struct {
	Test *fyne.Container
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Maze")

	grid := world.NewGrid(30, 30)
	myWindow.SetContent(grid.DrawGrid())

	myWindow.Resize(fyne.NewSize(600, 600))
	myWindow.SetFixedSize(true)
	myWindow.SetPadded(false)
	myWindow.ShowAndRun()
}
