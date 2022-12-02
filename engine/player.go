package engine

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	playerNumber int
	X            float64
	Y            float64
	speed        float64
	direction    [2]float64
	playerColor  color.RGBA
}

func (p *Player) Move() {
	p.Y = p.Y + float64(p.speed*p.direction[1])
	if p.Y < 0 || p.Y > 1 {
		p.direction[1] = -p.direction[1]
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	playerWidth := 4.0
	playerHeight := 20.0
	ebitenutil.DrawRect(
		screen,
		p.X*float64(screen.Bounds().Max.X)-(playerWidth/2),
		p.Y*float64(screen.Bounds().Max.Y)-(playerHeight/2),
		playerWidth,
		playerHeight,
		p.playerColor,
	)
}

func (p *Player) Initialize(playerNumber int) {
	p.playerNumber = playerNumber
	p.Y = 0.5
	direction := rand.Intn(2)
	if direction == 0 {
		direction = -1
	}
	p.direction = [2]float64{float64(0), float64(direction)}
	p.speed = 0.005
	if playerNumber == 1 {
		p.X = 0.05
		p.playerColor = color.RGBA{255, 0, 0, 255}
	}
	if playerNumber == 2 {
		p.X = 0.95
		p.playerColor = color.RGBA{0, 255, 0, 255}
	}
}

func (p *Player) Debug() {
	fmt.Printf("Player%v | x: %.2f y: %.2f speed: %.2f direction:[%.2f,%.2f]\n", p.playerNumber, p.X, p.Y, p.speed, p.direction[0], p.direction[1])
}
