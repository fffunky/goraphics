package goraphics

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct {
	keyInput   KeyInput
	mouseInput MouseInput
	debug bool
}

func (i *Input) MouseX() int {
	return i.mouseInput.mouseX
}

func (i *Input) MouseY() int {
	return i.mouseInput.mouseY
}

func (i *Input) MouseXY() (x, y int) {
	return i.mouseInput.mouseX, i.mouseInput.mouseY
}

func (i *Input) Update() {
	i.keyInput.Update()
	i.mouseInput.Update()
	
	if !i.debug && inpututil.IsKeyJustPressed(ebiten.Key0) {
		i.debug = true
	} else if i.debug && inpututil.IsKeyJustPressed(ebiten.Key0) {
		i.debug = false
	}

	if i.debug {
		i.keyInput.PrintKeyPresses()
		if i.mouseInput.StateChanged() {
			i.mouseInput.PrintMouseInfo()
		}
	}
}

func (i *Input) Debug() bool {
	return i.debug
}

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

// returns a [-1, 1] value for each axis
func (d Direction) Vector() (x, y int) {
	switch d {
	case UP:
		return 0, -1
	case RIGHT:
		return 1, 0
	case DOWN:
		return 0, 1
	case LEFT:
		return -1, 0
	}
	panic("not reach")
}

// returns the correct direction given a a vector=(dx, dy int)
func vectorToDirection(dx, dy int) (Direction, bool) {
	if abs(dx) < 4 && abs(dy) < 4 {
		return 0, false
	}
	if abs(dx) < abs(dy) {
		if dy < 0 {
			return UP, true
		}
		return DOWN, true
	}
	if dx < 0 {
		return LEFT, true
	}
	return RIGHT, true
}

// KEY INPUT

type KeyInput struct {
	pressedKeys     []ebiten.Key
	justPressedKeys []ebiten.Key
}

func (ki *KeyInput) Update() {
	ki.pressedKeys = inpututil.AppendPressedKeys(ki.pressedKeys[:0])
	ki.justPressedKeys = inpututil.AppendJustPressedKeys(ki.justPressedKeys[:0])
}

func (ki *KeyInput) PressedKeys() []ebiten.Key {
	return ki.pressedKeys
}

func (ki *KeyInput) PrintKeyPresses() {
	for _, k := range ki.justPressedKeys {
		fmt.Printf("keypress: %s\n", k.String())
	}
}

// MOUSE INPUT

type mouseState int

const (
	mouseStateNone mouseState = iota
	mouseStatePressing
	mouseStateSettled
)

type MouseInput struct {
	mouseState         mouseState
	lastMouseState     mouseState
	mouseX             int
	mouseY             int
	mouseInitPosX      int
	mouseInitPosY      int
	mouseDragDirection Direction
}

func (mi *MouseInput) Update() {
	mi.mouseX, mi.mouseY = ebiten.CursorPosition()
	mi.lastMouseState = mi.mouseState

	switch mi.mouseState {
	case mouseStateNone:
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			mi.mouseInitPosX = mi.mouseX
			mi.mouseInitPosY = mi.mouseY
			mi.mouseState = mouseStatePressing
		}
	case mouseStatePressing:
		if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			dx := mi.mouseX - mi.mouseInitPosX
			dy := mi.mouseY - mi.mouseInitPosY
			d, ok := vectorToDirection(dx, dy)
			if !ok {
				mi.mouseState = mouseStateNone
				break
			}
			mi.mouseDragDirection = d
			mi.mouseState = mouseStateSettled
		}
	case mouseStateSettled:
		mi.mouseState = mouseStateNone
	}
	
}

func (mi *MouseInput) PrintMouseInfo() {
	fmt.Printf("mouse: (x, y) = (%d, %d), state: %s, last drag direction: %s\n",
		mi.mouseX, mi.mouseY, mi.mouseState.String(), mi.mouseDragDirection.String())
}

func (mi *MouseInput) StateChanged() bool {
	return mi.mouseState != mi.lastMouseState
}

func (mi *MouseInput) Pressed() bool {
	return mi.mouseState == mouseStatePressing
}

func (mi *MouseInput) Clicked() bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func (mi *MouseInput) Released() bool {
	return inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (m mouseState) String() string {
	switch m {
	case mouseStateNone:
		return "none"
	case mouseStatePressing:
		return "pressing"
	case mouseStateSettled:
		return "settled"
	}
	panic("not reach")
}

func (d Direction) String() string {
	switch d {
	case UP:
		return "up"
	case RIGHT:
		return "right"
	case DOWN:
		return "down"
	case LEFT:
		return "left"
	}
	panic("not reach")
}
