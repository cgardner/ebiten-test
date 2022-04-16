package entity

import (
	"image/color"

	"code.rocketnine.space/tslocum/gohan"
	"github.com/cgardner/ebiten-test/lib/component"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	Position component.Position
}

var playerSprite = ebiten.NewImage(32, 32)

func init() {
	playerSprite.Fill(color.White)
}

func NewPlayer() *gohan.Entity {
	player := gohan.NewEntity()

	player.AddComponent(&component.Position{X: 0, Y: 0})
	player.AddComponent(&component.Sprite{Image: playerSprite})

	return &player
}
