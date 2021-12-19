package world

import (
	"image/color"
	"maze/lib"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type Grid struct {
	lock          sync.RWMutex
	width, height int
	cells         []lib.Agent
}

func NewGrid(width, height int) *Grid {
	g := &Grid{
		width:  width,
		height: height,
	}
	g.cells = make([]lib.Agent, g.size())
	for i := 0; i < g.size(); i++ {
		g.cells[i] = nil
	}
	return g
}

func (g *Grid) DrawGrid() *fyne.Container {
	agents := make([]fyne.CanvasObject, g.size())
	for i := 0; i < g.size(); i++ {
		agents[i] = canvas.NewRectangle(color.White)
	}
	return container.New(layout.NewGridLayout(g.width), agents...)
}

func (g *Grid) size() int {
	return g.height * g.width
}
