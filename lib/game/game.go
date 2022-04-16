package game

import (
	"github.com/cgardner/ebiten-test/lib/entity"
	"github.com/cgardner/ebiten-test/lib/system"

	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Player *gohan.Entity
}

var debugPrint = ebitenutil.DebugPrint

func NewGame() *Game {
	g := &Game{}
	g.addSystems()
	g.Player = entity.NewPlayer()
	return g
}

func (g *Game) Update() error {
	gohan.Update()
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 320, 240
}

func (g *Game) Draw(screen *ebiten.Image) {
	gohan.Draw(screen)
}

func (g *Game) addSystems() {
	gohan.AddSystem(system.NewPlayerDrawSystem())
}
