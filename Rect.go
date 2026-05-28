package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Rect struct {
	x, y, w, h float64
}

func (rect *Rect) draw(screen *ebiten.Image, c color.Color) {
	vector.DrawFilledRect(screen, float32(rect.x), float32(rect.y), float32(rect.w), float32(rect.h), c, false)
}

func (rect *Rect) collide(o Rect) bool {
	if rect.x >= o.x + o.w {
		return false
	}

	if rect.y >= o.y + o.h {
		return false
	}

	if o.x >= rect.x + rect.w {
		return false
	}

	if o.y >= rect.y + rect.h {
		return false
	}

	return true
}
