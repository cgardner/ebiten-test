package entity

import "github.com/cgardner/ebiten-test/lib/interfaces"

type Entity interface {
	Draw(screen interfaces.Drawable)
}
