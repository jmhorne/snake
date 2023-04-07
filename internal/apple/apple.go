package apple

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Apple struct {
	X, Y float64
}

func New(X, Y float64) *Apple {
	a := new(Apple)
	a.X = X
	a.Y = Y
	return a
}

func (a *Apple) Draw(screen *ebiten.Image) {
	x := float32(a.X * 32) + 16
	y := float32(a.Y * 32) + 16
	color := color.RGBA{0, 0, 0xFF, 0xFF}
	vector.DrawFilledCircle(screen, x, y, 16, color, true)
}