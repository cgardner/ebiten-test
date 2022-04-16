package system

import (
	"code.rocketnine.space/tslocum/gohan"
	"github.com/cgardner/ebiten-test/lib/component"
	"github.com/hajimehoshi/ebiten/v2"
)

type playerDrawSystem struct {
	Position *component.Position
	Sprite   *component.Sprite

	op *ebiten.DrawImageOptions
}

func NewPlayerDrawSystem() *playerDrawSystem {
	return &playerDrawSystem{
		op: &ebiten.DrawImageOptions{},
	}
}

func (s *playerDrawSystem) Update(_ gohan.Entity) error {
	return gohan.ErrUnregister
}

func (s *playerDrawSystem) Draw(entity gohan.Entity, screen *ebiten.Image) error {
	s.op.GeoM.Reset()
	s.op.GeoM.Translate(s.Position.X, s.Position.Y)
	screen.DrawImage(s.Sprite.Image, s.op)
	return nil
}
