package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"snake/internal/apple"
	"snake/internal/snake"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	player    *snake.Snake
	apple     *apple.Apple
	score     int
	running   bool
	quitCause string
}

func New() (*Game, error) {
	g := new(Game)
	err := g.Reset()
	return g, err
}

func (g *Game) Reset() error {
	var err error
	g.player, err = snake.New()
	g.apple = apple.New(float64(rand.Intn(15)), float64(rand.Intn(15)))
	g.score = 0
	g.running = true
	return err
}

func (g *Game) Update() error {
	if !g.running {
		var err error
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			err = g.Reset()
		}
		return err
	}

	g.player.Update()

	if g.player.TouchesPos(g.apple.X, g.apple.Y) {
		g.score++
		g.apple.X = float64(rand.Intn(15))
		g.apple.Y = float64(rand.Intn(15))
		g.player.Grow()
	}

	if g.player.BitSelf {
		g.running = false
		g.quitCause = "YOU BIT YOURSELF!"
	}
	if g.player.HitWall {
		g.running = false
		g.quitCause = "YOU HIT THE WALL"
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0xFF, 0, 0xFF})

	g.player.Draw(screen)
	g.apple.Draw(screen)

	if g.running {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.score))
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("%s High Score: %d", g.quitCause, g.score))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 480, 480
}
