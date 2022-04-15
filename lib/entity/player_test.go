package entity

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockimage struct {
	mock.Mock
}

func (m *mockimage) DrawImage(screen *ebiten.Image, opts *ebiten.DrawImageOptions) {
	m.Called(screen, opts)
}

func TestNewPlayer(t *testing.T) {
	p := NewPlayer()
	assert.IsType(t, &Player{}, p)
}

func TestPlayerDraw(t *testing.T) {
	s := &mockimage{}
	s.On("DrawImage", mock.Anything, mock.Anything).Return().Once()
	p := NewPlayer()

	p.Draw(s)

	s.AssertCalled(t, "DrawImage", playerSprite, mock.Anything)
}
