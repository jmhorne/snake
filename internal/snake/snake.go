package snake

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	headColor          = color.RGBA{0xFF, 0x25, 0xa1, 0xFF}
	bodyColor          = color.RGBA{0xFF, 0, 0, 0xFF}
	bodyRadius float32 = 16
)

type Snake struct {
	xPos, yPos   float64
	dX, dY, d    float64
	lastX, lastY float64
	body         []*body
}

func New() (*Snake, error) {
	s := new(Snake)

	s.xPos = 0
	s.yPos = 13
	s.lastX = s.xPos
	s.lastY = s.yPos
	s.d = 0.1
	s.dX = 0
	s.dY = -s.d

	s.body = make([]*body, 0)

	s.body = append(s.body, &body{xPos: s.xPos, yPos: s.yPos + 1})
	return s, nil
}

func (s *Snake) Draw(screen *ebiten.Image) {
	// draw head
	x := float32((math.Floor(s.xPos) * 32) + 16)
	y := float32((math.Floor(s.yPos) * 32) + 16)
	vector.DrawFilledCircle(screen, x, y, bodyRadius, headColor, true)

	// draw body
	for _, b := range s.body {
		b.draw(screen)
	}
}

func (s *Snake) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		s.dX = 0
		s.dY = -s.d
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		s.dX = 0
		s.dY = s.d
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		s.dX = -s.d
		s.dY = 0
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		s.dX = s.d
		s.dY = 0
	}

	s.xPos += s.dX
	s.yPos += s.dY

	if s.xPos >= (s.lastX + 1) || s.xPos < s.lastX || s.yPos >= (s.lastY + 1) || s.yPos < s.lastY {
		for _, b := range s.body {
			b.update(s.lastX, s.lastY)
		}

		s.lastX = math.Floor(s.xPos)
		s.lastY = math.Floor(s.yPos)
	}

	return nil
}

func (s *Snake) Grow() {
	
}