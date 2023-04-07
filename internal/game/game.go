package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"snake/internal/apple"
	"snake/internal/snake"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player *snake.Snake
	apple  *apple.Apple
	score  int
}

func New() (*Game, error) {
	var err error
	g := new(Game)
	g.player, err = snake.New()
	g.apple = apple.New(14, 14)
	g.score = 0
	return g, err
}

func (g *Game) Update() error {
	g.player.Update()

	if g.player.TouchesApple(g.apple.X, g.apple.Y) {
		g.score++
		g.apple.X = float64(rand.Intn(15))
		g.apple.Y = float64(rand.Intn(15))
		g.player.Grow()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0xFF, 0, 0xFF})

	g.player.Draw(screen)
	g.apple.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.score))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 480, 480
}
