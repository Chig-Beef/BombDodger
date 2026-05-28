package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 500
)

func main() {
	g := createGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Bomb Dodger")
	ebiten.SetTPS(60)

	if err := ebiten.RunGame(g); err != nil {
		fmt.Println(err)
	}
}
