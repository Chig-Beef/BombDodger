package main

import (
	"BombDodger/brain"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const agentSize = 50
const agentSpeed = 8.0

const IN_BOMB_ABOVE = 0
const IN_BOMB_LEFT = 1
const IN_BOMB_RIGHT = 2
const IN_WALL_LEFT = 3
const IN_WALL_RIGHT = 4
const IN_CONSTANT = 5
const NUM_INPUTS = 6

const OUT_MOVE_LEFT = 0
const OUT_MOVE_RIGHT = 1
const NUM_OUTPUTS = 2

const LAYERS = 3
const DEPTH = 6

type Agent struct {
	brain brain.Brain
	rect Rect
	timeAlive int
	alive bool
}

func newAgent(x, y float64) Agent {
	a := Agent{}

	a.brain = brain.NewBrain(NUM_INPUTS, LAYERS, DEPTH, NUM_OUTPUTS)
	a.brain.Randomize()

	a.rect = Rect{x, y, agentSize, agentSize}

	a.timeAlive = 0
	a.alive = true

	return a
}

func (a *Agent) update(bombs []Bomb) {
	a.timeAlive++

	for i := 0; i < len(bombs); i++ {
		if a.rect.collide(bombs[i].rect) {
			a.alive = false
			return
		}
	}

	a.brain.In(float32(a.checkBombAbove(bombs)), IN_BOMB_ABOVE)
	a.brain.In(float32(a.checkBombLeft(bombs)), IN_BOMB_LEFT)
	a.brain.In(float32(a.checkBombRight(bombs)), IN_BOMB_RIGHT)
	a.brain.In(float32(a.checkWallLeft()), IN_WALL_LEFT)
	a.brain.In(float32(a.checkWallRight()), IN_WALL_RIGHT)
	a.brain.In(1.0, IN_CONSTANT)

	a.brain.Push()

	boolOutput, _ := a.brain.Dump()

	if boolOutput[OUT_MOVE_LEFT] {
		a.moveLeft()
	}

	if boolOutput[OUT_MOVE_RIGHT] {
		a.moveRight()
	}

	a.brain.Set()
}

func (a *Agent) draw(screen *ebiten.Image) {
	a.rect.draw(screen, color.RGBA{64, 128, 196, 255})
}

// INPUTS

func (a *Agent) checkBombAbove(bombs []Bomb) float64 {
	if len(bombs) == 0 {
		return 0.0
	}

	// Too high to care
	if bombs[0].rect.y < 150 {
		return 0.0
	}

	// Too far left (with gap)
	if bombs[0].rect.x + bombs[0].rect.w < a.rect.x-5 {
		return 0.0
	}

	// Too far right (with gap)
	if bombs[0].rect.x > a.rect.x+a.rect.w+5 {
		return 0.0
	}

	return 1.0
}

func (a *Agent) checkBombLeft(bombs []Bomb) float64 {
	if len(bombs) == 0 {
		return 0.0
	}

	if bombs[0].rect.x+bombs[0].rect.w < a.rect.x - a.rect.w - 5 {
		return 0.0
	}

	return 1.0
}

func (a *Agent) checkBombRight(bombs []Bomb) float64 {
	if len(bombs) == 0 {
		return 0.0
	}

	if bombs[0].rect.x > a.rect.x + a.rect.w*2 + 5 {
		return 0.0
	}

	return 1.0
}

func (a *Agent) checkWallLeft() float64 {
	if a.rect.x < a.rect.w*3 {
		return 1.0
	}

	return 0.0
}

func (a *Agent) checkWallRight() float64 {
	if a.rect.x+a.rect.w > simWidth-a.rect.w*3 {
		return 1.0
	}

	return 0.0
}

// OUTPUTS

func (a *Agent) moveLeft() {
	a.rect.x -= agentSpeed
	if a.rect.x < 0 {
		a.rect.x = 0
	}
}

func (a *Agent) moveRight() {
	a.rect.x += agentSpeed
	if a.rect.x > simWidth-agentSize {
		a.rect.x = simWidth-agentSize
	}
}
