package game

import (
	"github.com/cgardner/ebiten-test/lib/entity"
	"github.com/cgardner/ebiten-test/lib/interfaces"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Player  *entity.Player
	Systems []*interfaces.System
}

var debugPrint = ebitenutil.DebugPrint

func NewGame() *Game {
	g := &Game{}
	g.Player = entity.NewPlayer()
	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)
}

// Add a system to the game
func (g *Game) AddSystem(sys interfaces.System) {
	g.Systems = append(g.Systems, &sys)
}
