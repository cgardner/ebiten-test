package game

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/cgardner/ebiten-test/lib/game/mocks"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewGame(t *testing.T) {
	g := NewGame()
	if g == nil {
		t.Error("NewGame() should not return nil")
	}
}

func TestUpdateReturnsNil(t *testing.T) {
	g := NewGame()
	if g.Update() != nil {
		t.Error("Update() should return nil")
	}
}

func TestLayoutReturnsStatic(t *testing.T) {
	type testCase struct {
		input    [2]int
		expected [2]int
	}
  expected := [2]int{320,240}
	testCases := []testCase{
		{
			[2]int{1, 2},
      expected,
		},
    {
      expected,
      expected,
    },
		{
			[2]int{rand.Int(), rand.Int()},
      expected,
		},
	}

	g := &Game{}
	for _, testCase := range testCases {
		description := fmt.Sprintf("Layout(%v) should return %v", testCase.input, testCase.expected)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			width, height := g.Layout(testCase.input[0], testCase.input[1])
			assert.Equal(t, testCase.expected[0], width)
			assert.Equal(t, testCase.expected[1], height)
		})
	}
}

func TestDrawReturnsNil(t *testing.T) {
  origDebugPrint := debugPrint
  defer func() { debugPrint = origDebugPrint }()
  image := ebiten.NewImage(320, 240)
  m := &mocks.EbitenUtilMock{}

  m.On("DebugPrint", mock.AnythingOfType("*ebiten.Image"), "Hello, World!").Once()

  debugPrint = m.DebugPrint

  g := NewGame()

  g.Draw(image)
  
  m.AssertExpectations(t)
}

