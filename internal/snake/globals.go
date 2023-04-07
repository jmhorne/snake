package snake

import "image/color"

type DELTA []float64

var (
	headColor          = color.RGBA{0xFF, 0x25, 0xa1, 0xFF}
	bodyColor          = color.RGBA{0xFF, 0, 0, 0xFF}
	bodyRadius float32 = 16
	delta              = 0.1
	deltas             = map[DIRECTION]DELTA{
		LEFT:  {-delta, 0},
		RIGHT: {delta, 0},
		UP:    {0, -delta},
		DOWN:  {0, delta},
	}
)

type DIRECTION int

const (
	LEFT DIRECTION = iota
	RIGHT
	UP
	DOWN
)
