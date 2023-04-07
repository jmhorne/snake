package main

import (
	"log"
	"snake/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := game.New()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(480, 480)
	ebiten.SetWindowTitle("Snake!")
	// ebiten.SetTPS(2)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}