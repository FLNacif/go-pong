package engine

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"

	"github.com/FLNacif/go-pong/math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ball struct {
	X         float64
	Y         float64
	speed     float64
	ballColor color.RGBA
	direction math.Vector
}

const (
	ballWidth  = 3.0
	ballHeight = 3.0
	ballSpeed  = 0.008
	ballStartX = 0.5
	ballStartY = 0.5
)

var (
	ballColor = color.RGBA{255, 255, 255, 255}
)

func (b *Ball) Bounds() image.Rectangle {
	return image.Rectangle{}
}

func (b *Ball) Draw(screen *ebiten.Image) {
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
	b.X = ballStartX
	b.Y = ballStartY
	b.ballColor = ballColor
	b.speed = ballSpeed
	b.direction = math.Vector{rand.Float64()*10 - 5, rand.Float64()*10 - 5}
	b.direction.Normalize()
}

func (b *Ball) Debug() {
	fmt.Printf("Ball    | %p | x: %.2f y: %.2f speed: %.2f direction:[%.2f,%.2f]\n", b, b.X, b.Y, b.speed, b.direction[0], b.direction[1])
}
