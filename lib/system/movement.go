package system

import (
	"code.rocketnine.space/tslocum/gohan"
	"github.com/cgardner/ebiten-test/lib/component"
	"github.com/cgardner/ebiten-test/lib/world"
	"github.com/hajimehoshi/ebiten/v2"
)

type MovementSystem struct {
	Position *component.Position
	Velocity *component.Velocity
}

const PLAYER_WIDTH = 32 

func NewMovementSystem() *MovementSystem {
	return &MovementSystem{}
}

func (s *MovementSystem) Update(entity gohan.Entity) error {
	// Check for collision.
	if s.Position.X+s.Velocity.X < 0 {
		s.Position.X = 0
		s.Velocity.X = 0
	} else if s.Position.X+s.Velocity.X > world.ScreenW-PLAYER_WIDTH {
		s.Position.X = world.ScreenW - PLAYER_WIDTH
		s.Velocity.X = 0
	}
	if s.Position.Y+s.Velocity.Y < 0 {
		s.Position.Y = 0
		s.Velocity.Y = 0
	} else if s.Position.Y+s.Velocity.Y > world.ScreenH-PLAYER_WIDTH {
		s.Position.Y = world.ScreenH - PLAYER_WIDTH
		s.Velocity.Y = 0
	}

	s.Position.X, s.Position.Y = s.Position.X+s.Velocity.X, s.Position.Y+s.Velocity.Y

	return nil
}

func (_ *MovementSystem) Draw(_ gohan.Entity, _ *ebiten.Image) error {
	return gohan.ErrUnregister
}
