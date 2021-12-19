package agents

type Wall struct {
	x, y int
}

func NewWall(x, y int) *Wall {
	return &Wall{
		x: x,
		y: y,
	}
}

func (w *Wall) X() int { return w.x }
func (w *Wall) Y() int { return w.y }
