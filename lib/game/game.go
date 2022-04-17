package game

import (
	"github.com/cgardner/ebiten-test/lib/entity"
	"github.com/cgardner/ebiten-test/lib/system"
	"github.com/cgardner/ebiten-test/lib/world"

	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Player *gohan.Entity
	w, h   int
}

var debugPrint = ebitenutil.DebugPrint

func NewGame() *Game {
	g := &Game{}
	g.addSystems()
	world.Player = *entity.NewPlayer()
	gohan.Preallocate(10000)
	return g
}

func (g *Game) Update() error {
	return gohan.Update()
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	s := ebiten.DeviceScaleFactor()
	w, h := int(s*float64(outsideWidth)), int(s*float64(outsideHeight))
	if w != g.w || h != g.h {
		g.w, g.h = w, h
		world.ScreenW, world.ScreenH = float64(w), float64(h)
	}
	return g.w, g.h
}

func (g *Game) Draw(screen *ebiten.Image) {
	err := gohan.Draw(screen)
	if err != nil {
		panic(err)
	}
}

func (g *Game) addSystems() {
	gohan.AddSystem(system.NewMovementSystem())
	gohan.AddSystem(system.NewPlayerDrawSystem())
	gohan.AddSystem(system.NewPlayerInputSystem())
}
