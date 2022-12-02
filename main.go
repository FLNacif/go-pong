package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/FLNacif/go-pong/engine"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	State *engine.State
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (g *Game) Update() error {
	g.State.Update()
	g.State.PrintState()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.State.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	s := new(engine.State)
	s.InitializeState()

	ebiten.SetWindowSize(640, 480)

	if err := ebiten.RunGame(&Game{s}); err != nil {
		log.Fatal(err)
	}
}
