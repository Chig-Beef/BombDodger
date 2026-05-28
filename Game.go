package main

import (
	"image/color"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const bombTimer = 180
const bottomBar = 100
const sideBar = 200
const populationSize = 100
const startX = (screenWidth-sideBar) / 2 - agentSize / 2
const startY = screenHeight - bottomBar - agentSize
const mutationRate = 0.1
const simWidth = screenWidth-sideBar
const simHeight = screenHeight-bottomBar
const booster = 20

type Game struct {
	fonts map[string]*text.GoTextFaceSource

	fastMode bool
	ticker int

	bombFreq int
	bombFreqConst int

	generation int
	bestTime int
	newBestTime int
	bombs []Bomb
	agents []Agent

	improvement *ebiten.Image
	fits []int
	drawn int
}

func (g *Game) updateWorld() {
	g.ticker++
	if g.ticker%3600 == 0 {
		// Faster bombs!
		if g.bombFreqConst > 0 {
			g.bombFreqConst--
		}
	}

	g.bombFreq--
	if g.bombFreq == 0 {
		g.bombFreq = g.bombFreqConst
		g.createBomb()
	}

	// Update bombs
	aliveCount := 0
	for i := 0; i < len(g.bombs); i++ {
		if g.bombs[i].update() {
			aliveCount++
		}
	}

	newBombs := make([]Bomb, aliveCount)
	j := 0
	for i := 0; i < len(g.bombs); i++ {
		if g.bombs[i].alive {
			newBombs[j] = g.bombs[i]
			j++
		}
	}
	g.bombs = newBombs

	dead := true
	for i := 0; i < len(g.agents); i++ {
		if !g.agents[i].alive {
			continue
		}

		dead = false
		g.agents[i].update(g.bombs)
	}

	if dead {
		g.nextGeneration()
	}
}

func (g *Game) Update() error {
	// Activate fastMode
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		if g.fastMode {
			ebiten.SetTPS(60)
		} else {
			ebiten.SetTPS(20_000)
		}
		g.fastMode = !g.fastMode
	}

	if g.fastMode {
		for i := 0; i < booster; i++ {
			g.updateWorld()
		}
	} else {
		g.updateWorld()
	}

	return nil
}

func (g *Game) createBomb() {
	bomb := newBomb()
	g.bombs = append(g.bombs, bomb)
}

func (g *Game) drawImprovement() {
	// g.improvement.Fill(color.RGBA{0, 255, 0, 255})
	for i := g.drawn; i < screenWidth; i++ {
		t := g.fits[i]
		g.drawn = i

		if t == 0 {
			break
		}

		// Compared to 5 hours stright
		h := float32(float32(t) * bottomBar / (60.0*60.0*60.0*5.0))

		x := float32(i+5)
		y := float32(bottomBar)-h

		vector.StrokeLine(g.improvement, x, y, x, bottomBar, 1, color.RGBA{0, 255, 0, 255}, true)
	}
}

func (g *Game) nextGeneration() {
	g.generation++
	g.ticker = 0
	g.bombs = []Bomb{}
	g.bombFreqConst = bombTimer
	g.bombFreq = bombTimer

	g.newBestTime = g.findBestScore(g.agents)
	if g.newBestTime > g.bestTime {
		g.bestTime = g.newBestTime
	}

	if g.generation < screenWidth {
		g.fits[g.generation-1] = g.newBestTime
	}

	totalFitness := 0
	for i := 0; i < populationSize; i++ {
		totalFitness += g.agents[i].timeAlive
	}

	newPopulation := make([]Agent, populationSize)

	for i := 0; i < populationSize; i++ {
		a := newAgent(startX, startY)

		parentA := g.randomParent(int(float64(totalFitness)*rand.Float64()))
		parentB := g.randomParent(int(float64(totalFitness)*rand.Float64()))

		a.brain.Merge(&parentA.brain, &parentB.brain, mutationRate)

		newPopulation[i] = a
	}

	g.agents = newPopulation
}

func (g *Game) initPopulation() {
	g.agents = make([]Agent, populationSize)
	for i := 0; i < populationSize; i++ {
		a := newAgent(startX, startY)
		g.agents[i] = a
	}
}

func (g *Game) findBestScore(population []Agent) int {
	best := 0
	for i := 0; i < len(population); i++ {
		if population[i].timeAlive > best {
			best = population[i].timeAlive
		}
	}
	return best
}

func (g *Game) randomParent(fitScale int) Agent {
	for i := 0; i < len(g.agents); i++ {
		fitScale -= g.agents[i].timeAlive
		if fitScale <= 0 {
			return g.agents[i]
		}
	}
	return g.agents[len(g.agents)-1]
}

func (g *Game) Draw(screen *ebiten.Image) {
	for i := 0; i < len(g.bombs); i++ {
		g.bombs[i].draw(screen)
	}

	count := 0
	for i := 0; i < len(g.agents); i++ {
		if g.agents[i].alive {
			g.agents[i].draw(screen)
			count++
		}
	}

	(&Rect{0, screenHeight-bottomBar, screenWidth-sideBar, bottomBar}).draw(screen, color.RGBA{128, 128, 128, 255})
	(&Rect{screenWidth-sideBar, 0, sideBar, screenHeight}).draw(screen, color.RGBA{128, 128, 128, 255})

	g.drawImprovement()
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, simHeight)
	screen.DrawImage(g.improvement, &op)

	textSize := 15.0
	f := g.fonts["default"]

	g.drawText(screen, 100, simHeight+10, "TPS: " + strconv.FormatFloat(ebiten.ActualTPS(), 'f', 4, 64), f, textSize, text.AlignCenter)
	g.drawText(screen, 100, simHeight+30, "Gen: " + strconv.Itoa(g.generation), f, textSize, text.AlignCenter)
	g.drawText(screen, 100, simHeight+50, "Count: " + strconv.Itoa(count), f, textSize, text.AlignCenter)

	g.drawText(screen, 400, simHeight+10, "Last Best Time (seconds): " + strconv.Itoa(g.newBestTime/60), f, textSize, text.AlignCenter)
	g.drawText(screen, 400, simHeight+30, "Last Best Time (minutes): " + strconv.Itoa(g.newBestTime/3600), f, textSize, text.AlignCenter)
	g.drawText(screen, 400, simHeight+50, "All Time Best (seconds): " + strconv.Itoa(g.bestTime/60), f, textSize, text.AlignCenter)
	g.drawText(screen, 400, simHeight+70, "All Time Best (minutes): " + strconv.Itoa(g.bestTime/3600), f, textSize, text.AlignCenter)
}

func (g *Game) Layout(int, int) (int, int) {
	return screenWidth, screenHeight
}

func createGame() *Game {
	g := Game{}

	g.bombFreqConst = bombTimer
	g.bombFreq = bombTimer

	g.initPopulation()

	g.loadFonts()
	g.improvement = ebiten.NewImage(screenWidth, bottomBar)
	g.fits = make([]int, screenWidth)

	return &g
}
