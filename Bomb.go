package main

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const bombSize = 50
const bombSpeed = 5.0

type Bomb struct {
	rect Rect
	alive bool
}

func newBomb() Bomb {
	b := Bomb{}
	x := rand.Intn(simWidth-bombSize)
	b.rect = Rect{float64(x), -bombSize, bombSize, bombSize}
	b.alive = true
	return b
}

func (b *Bomb) update() bool {
	b.rect.y += bombSpeed
	b.alive = b.rect.y < simHeight
	return b.alive
}

func (b *Bomb) draw(screen *ebiten.Image) {
	b.rect.draw(screen, color.RGBA{196, 64, 64, 255})
}
