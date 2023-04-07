package snake

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Snake struct {
	xPos, yPos   float64
	lastX, lastY float64
	body         []*body
	delta    DELTA
	BitSelf bool
	HitWall bool
}

func New() (*Snake, error) {
	s := new(Snake)

	s.xPos = 0
	s.yPos = 14
	s.lastX = s.xPos
	s.lastY = s.yPos
	s.delta = deltas[UP]
	s.BitSelf = false
	s.HitWall = false

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
		s.delta = deltas[UP]
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		s.delta = deltas[DOWN]
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		s.delta = deltas[LEFT]
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		s.delta = deltas[RIGHT]
	}

	s.xPos += s.delta[0]
	s.yPos += s.delta[1]

	if !(s.xPos >= (s.lastX+1) || s.xPos < s.lastX || s.yPos >= (s.lastY+1) || s.yPos < s.lastY) {
		return nil
	}

	s.body[0].update(s.lastX, s.lastY)

	if s.xPos < 0 || s.xPos > 15 || s.yPos < 0 || s.yPos > 15 {
		s.HitWall = true
		return nil
	}

	s.lastX = math.Floor(s.xPos)
	s.lastY = math.Floor(s.yPos)

	for _, b := range s.body {
		if s.TouchesPos(b.xPos, b.yPos) {
			s.BitSelf = true
		}
	}

	return nil
}

func (s *Snake) Grow() {
	tail := s.body[len(s.body) - 1]

	newBody := &body{xPos: tail.xPos - s.delta[0], yPos: tail.yPos - s.delta[1], next: nil}

	tail.next = newBody

	s.body = append(s.body, newBody)
}

func (s *Snake) TouchesPos(aX, aY float64) bool {
	return math.Floor(s.xPos) == math.Floor(aX) && math.Floor(s.yPos) == math.Floor(aY)
}