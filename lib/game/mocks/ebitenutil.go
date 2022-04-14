package mocks

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/mock"
)

type EbitenUtilMock struct {
	mock.Mock
}

func (m *EbitenUtilMock) DebugPrint(s *ebiten.Image, text string) {
	m.Called(s, text)
}
