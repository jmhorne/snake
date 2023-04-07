package snake

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type body struct {
	xPos, yPos float64
	next       *body
}

func (b *body) update(dX, dY float64) {
	if b.next != nil {
		b.next.update(b.xPos, b.yPos)
	}

	b.xPos = dX
	b.yPos = dY
}

func (b *body) draw(screen *ebiten.Image) {
	x := float32((math.Ceil(b.xPos) * 32) + 16)
	y := float32((math.Ceil(b.yPos) * 32) + 16)
	vector.DrawFilledCircle(screen, x, y, bodyRadius, bodyColor, true)
}
