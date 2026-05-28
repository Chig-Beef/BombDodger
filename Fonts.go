package main

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (g *Game) drawText(screen *ebiten.Image, x, y float64, str string, drawFont *text.GoTextFaceSource, size float64, alignment text.Align) {
	op := text.DrawOptions{}
	op.PrimaryAlign = alignment
	op.GeoM.Translate(x, y)
	text.Draw(
		screen,
		str,
		&text.GoTextFace{
			Source: drawFont,
			Size:   size,
		},
		&op,
	)
}

func (g *Game) loadFonts() {
	g.fonts = make(map[string]*text.GoTextFaceSource)

	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal(err)
	}

	g.fonts["default"] = s
}
