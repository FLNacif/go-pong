package engine

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ball struct {
	X         float64
	Y         float64
	speed     float64
	ballColor color.RGBA
	direction [2]float64
}

func (b *Ball) Draw(screen *ebiten.Image) {
	ballWidth := 3.0
	ballHeight := 3.0

	ebitenutil.DrawRect(
		screen,
		b.X*float64(screen.Bounds().Max.X)-(ballWidth/2),
		b.Y*float64(screen.Bounds().Max.Y)-(ballHeight/2),
		ballWidth,
		ballHeight,
		b.ballColor,
	)
}
func (b *Ball) Move() {
	b.X = b.X + float64(b.speed*b.direction[0])
	b.Y = b.Y + float64(b.speed*b.direction[1])
	if b.Y < 0 || b.Y > 1 {
		b.direction[1] = -b.direction[1]
	}
}

func (b *Ball) Reset() {
	b.Initialize()
}

func (b *Ball) Initialize() {
	b.X = 0.5
	b.Y = 0.5
	b.ballColor = color.RGBA{255, 255, 255, 255}
	b.speed = 0.002
	b.direction = [2]float64{rand.Float64()*10 - 5, rand.Float64()*10 - 5}
}

func (b *Ball) Debug() {
	fmt.Printf("Ball    | x: %.2f y: %.2f speed: %.2f direction:[%.2f,%.2f]\n", b.X, b.Y, b.speed, b.direction[0], b.direction[1])
}
