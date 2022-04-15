package entity

import (
	"image/color"

	"github.com/cgardner/ebiten-test/lib/component"
	"github.com/cgardner/ebiten-test/lib/interfaces"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Position component.Position
}

var playerSprite = ebiten.NewImage(32, 32)

func init() {
	playerSprite.Fill(color.White)
}

func NewPlayer() *Player {
	return &Player{
		Position: component.Position{
			X: float64(0),
			Y: float64(0),
		},
	}
}

// Draw the player
func (p *Player) Draw(screen interfaces.Drawable) {
	pos := p.Position
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(pos.X, pos.Y)
	screen.DrawImage(playerSprite, nil)
}
