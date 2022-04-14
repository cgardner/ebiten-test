package main

import (
	"github.com/cgardner/ebiten-test/lib/game"
	"github.com/hajimehoshi/ebiten/v2"
)


func main() {
  g:= game.NewGame()
  ebiten.SetWindowSize(640, 480)
  ebiten.SetWindowTitle("Hello, World!")
  err := ebiten.RunGame(g)
  if err != nil {
    panic(err)
  }
}
