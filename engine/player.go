package engine

import (
	"fmt"
	"image/color"

	"github.com/FLNacif/go-pong/consts"
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
	playerSpeed   float64 = 0.01
)

var (
	playerRED   color.RGBA = color.RGBA{255, 0, 0, 255}
	playerGREEN color.RGBA = color.RGBA{0, 255, 0, 255}
)

func (p *Player) Bounds() *Bounds {
	return &Bounds{
		[2]float64{
			p.X*float64(consts.CanvasWidth) - (playerWidth / 2),
			p.Y*float64(consts.CanvasHeight) - (playerHeight / 2),
		},
		[2]float64{
			p.X*float64(consts.CanvasWidth) + (playerWidth / 2),
			p.Y*float64(consts.CanvasHeight) + (playerHeight / 2),
		},
	}
}

func (p *Player) Update() {
	p.Y = p.Y + float64(p.speed*p.direction[1])
	if p.Y > 1 {
		p.Y = 1
	}
	if p.Y < 0 {
		p.Y = 0
	}
}

func (p *Player) ChangeDirection(direction consts.MovingDirection) {
	if direction == consts.Up {
		p.direction[1] = -1
	}
	if direction == consts.Down {
		p.direction[1] = 1
	}
	if direction == consts.Stop {
		p.direction[1] = 0
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(
		screen,
		p.X*float64(consts.CanvasWidth)-(playerWidth/2),
		p.Y*float64(consts.CanvasHeight)-(playerHeight/2),
		playerWidth,
		playerHeight,
		p.playerColor,
	)
}

func (p *Player) Initialize(playerNumber int) {
	p.playerNumber = playerNumber
	p.Y = 0.5

	p.direction = [2]float64{float64(0), float64(0)}
	p.speed = playerSpeed
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
