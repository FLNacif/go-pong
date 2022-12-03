package engine

import (
	"fmt"
	"image"
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

const (
	playerWidth   float64 = 4.0
	playerHeight  float64 = 20.0
	player1StartX float64 = 0.05
	player2StartX float64 = 0.95
	playerSpeed   float64 = 0.005
)

var (
	playerRED   color.RGBA = color.RGBA{255, 0, 0, 255}
	playerGREEN color.RGBA = color.RGBA{0, 255, 0, 255}
)

func (p *Player) Bounds() image.Rectangle {
	return image.Rectangle{}
}

func (p *Player) Move() {
	p.Y = p.Y + float64(p.speed*p.direction[1])
	if p.Y < 0 || p.Y > 1 {
		p.direction[1] = -p.direction[1]
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
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
		p.X = player1StartX
		p.playerColor = playerRED
	}
	if playerNumber == 2 {
		p.X = player2StartX
		p.playerColor = playerGREEN
	}
}

func (p *Player) Debug() {
	fmt.Printf("Player%v | %p | x: %.2f y: %.2f speed: %.2f direction:[%.2f,%.2f]\n", p.playerNumber, p, p.X, p.Y, p.speed, p.direction[0], p.direction[1])
}
