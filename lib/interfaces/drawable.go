package interfaces

import "github.com/hajimehoshi/ebiten/v2"

type Drawable interface {
	DrawImage(screen *ebiten.Image, opts *ebiten.DrawImageOptions)
}
