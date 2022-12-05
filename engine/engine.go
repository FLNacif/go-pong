package engine

import (
	"fmt"
	"image/color"
	"log"

	"github.com/FLNacif/go-pong/consts"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Entity interface {
	Draw(screen *ebiten.Image)
	Update()
	Debug()
	Bounds() *Bounds
}

type Bounds struct {
	Min [2]float64
	Max [2]float64
}

type State struct {
	Player1  *Player
	Player2  *Player
	Ball     *Ball
	score    [2]int
	Entities []Entity
}

func (s *State) Update() {
	s.isKeyJustPressed()
	for _, entitie := range s.Entities {
		entitie.Update()
	}
	s.CheckHit()
	s.CheckGoal()
}

var (
	titleArcadeFont font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    8,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

func (s *State) Draw(screen *ebiten.Image) {

	text.Draw(screen, fmt.Sprintf("RED %d x %d GREEN", s.score[0], s.score[1]), titleArcadeFont, 10, 10, color.White)
	for _, entitie := range s.Entities {
		entitie.Draw(screen)
	}
}

func (s *State) AddScore(playerNumber int) {
	if playerNumber == 1 {
		s.score[0] = s.score[0] + 1
	}
	if playerNumber == 2 {
		s.score[1] = s.score[1] + 1
	}
}

func (s *State) CheckGoal() {
	ballBounds := s.Ball.Bounds()
	player1Bounds := s.Player1.Bounds()
	player2Bounds := s.Player2.Bounds()
	if ballBounds.Max[0] < player1Bounds.Min[0] {
		s.AddScore(2)
		s.Ball.Reset()
	}
	if ballBounds.Min[0] > player2Bounds.Max[0] {
		s.AddScore(1)
		s.Ball.Reset()
	}
}

func (s *State) CheckHit() {
	ballBounds := s.Ball.Bounds()
	player1Bounds := s.Player1.Bounds()
	player2Bounds := s.Player2.Bounds()

	if ballBounds.Overlaps(player1Bounds) {
		s.Ball.Hit((*s.Player1))
	}
	if ballBounds.Overlaps(player2Bounds) {
		s.Ball.Hit((*s.Player2))
	}
}

func (s *State) isKeyJustPressed() {

	if inpututil.KeyPressDuration(ebiten.KeyArrowUp) > 0 {
		s.Player2.ChangeDirection(consts.Up)
	}
	if inpututil.KeyPressDuration(ebiten.KeyDown) > 0 {
		s.Player2.ChangeDirection(consts.Down)
	}
	if inpututil.KeyPressDuration(ebiten.KeyArrowUp) == 0 &&
		inpututil.KeyPressDuration(ebiten.KeyDown) == 0 {
		s.Player2.ChangeDirection(consts.Stop)
	}

	if inpututil.KeyPressDuration(ebiten.KeyW) > 0 {
		s.Player1.ChangeDirection(consts.Up)
	}
	if inpututil.KeyPressDuration(ebiten.KeyS) > 0 {
		s.Player1.ChangeDirection(consts.Down)
	}
	if inpututil.KeyPressDuration(ebiten.KeyW) == 0 &&
		inpututil.KeyPressDuration(ebiten.KeyS) == 0 {
		s.Player1.ChangeDirection(consts.Stop)
	}
}

func (s *State) InitializeState() {
	s.Player1 = new(Player)
	s.Player1.Initialize(1)
	s.Entities = append(s.Entities, s.Player1)

	s.Player2 = new(Player)
	s.Player2.Initialize(2)
	s.Entities = append(s.Entities, s.Player2)

	s.Ball = new(Ball)
	s.Ball.Initialize()
	s.Entities = append(s.Entities, s.Ball)
}

func (s *State) PrintState() {
	fmt.Println("\n################")
	for _, entitie := range s.Entities {
		entitie.Debug()
	}
	fmt.Println("################")
}

func (ball *Bounds) Overlaps(player *Bounds) bool {
	return ball.Min[0] < player.Max[0] && player.Min[0] < ball.Max[0] &&
		ball.Min[1] < player.Max[1] && player.Min[1] < ball.Max[1]
}
