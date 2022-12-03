package engine

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Entity interface {
	Draw(screen *ebiten.Image)
	Move()
	Debug()
}

type State struct {
	Player1  *Player
	Player2  *Player
	Ball     *Ball
	score    [2]int
	Entities []Entity
}

func (s *State) Update() {
	for _, entitie := range s.Entities {
		entitie.Move()
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
	if s.Ball.X < s.Player1.X {
		s.AddScore(2)
		s.Ball.Reset()

	}
	if s.Ball.X > s.Player2.X {
		s.AddScore(1)
		s.Ball.Reset()

	}
}

func (s *State) CheckHit() {

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
