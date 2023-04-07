package game

import (
	"image/color"
	"snake/internal/snake"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *snake.Snake
}

func New() (*Game, error) {
	var err error 
	g := new(Game)
	g.player, err = snake.New()
	return g, err
}

func (g *Game) Update() error {
	g.player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0xFF, 0, 0xFF})

	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 480, 480
}
