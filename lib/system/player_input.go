package system

import (
	"github.com/cgardner/ebiten-test/lib/component"

	"code.rocketnine.space/tslocum/gohan"
	"github.com/hajimehoshi/ebiten/v2"
)

type playerInputSystem struct {
	Velocity *component.Velocity
}

func NewPlayerInputSystem() *playerInputSystem {
	return &playerInputSystem{}
}

const MAX_SPEED = 15

func (s *playerInputSystem) Update(entity gohan.Entity) error {
	keyPressed := false
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		keyPressed = true
		s.Velocity.X -= 0.5
		if s.Velocity.X < -MAX_SPEED {
			s.Velocity.X = -MAX_SPEED
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		keyPressed = true
		s.Velocity.X += 0.5
		if s.Velocity.X > MAX_SPEED {
			s.Velocity.X = MAX_SPEED
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		keyPressed = true
		s.Velocity.Y -= 0.5
		if s.Velocity.Y < -MAX_SPEED {
			s.Velocity.Y = -MAX_SPEED
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		keyPressed = true
		s.Velocity.Y += 0.5
		if s.Velocity.Y > MAX_SPEED {
			s.Velocity.Y = MAX_SPEED
		}
	}
	// Set velocity to zero

	if !keyPressed {
		s.Velocity.X = 0
		s.Velocity.Y = 0
	}

	return nil
}

func (s *playerInputSystem) Draw(_ gohan.Entity, _ *ebiten.Image) error {
	return gohan.ErrUnregister
}
