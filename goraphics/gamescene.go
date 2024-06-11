package goraphics

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	imageGameBG *ebiten.Image	
	imageWindows = ebiten.NewImage(ScreenWidth, ScreenHeight)
	imageGameOver = ebiten.NewImage(ScreenWidth, ScreenHeight)
)

var (
	lightGray colorm.ColorM
)

func windowPosition() (x, y int) {
	return 20, 20
}

func windowWidthHeight() (w, h int) {
	return 100, 100
}

// init for images
func init() {
	img, _, err := ebitenutil.NewImageFromFile("./assets/solids/red.png")
	if err != nil {
		log.Fatal(err)
	}

	imageGameBG = ebiten.NewImageFromImage(img)


	// Windows: Field
	x, y := windowPosition()
	w, h := windowWidthHeight()
	drawWindow(imageWindows, x, y, w, h)
}

func drawWindow(r *ebiten.Image, x, y, width, height int) {
	vector.DrawFilledRect(r, float32(x), float32(y), float32(width), float32(height), color.RGBA{0, 0, 0, 0xc0}, false)
}

type GameScene struct {
	mouseX int
	mouseY int
	inputDebug bool
}

func NewGameScene() *GameScene {
	return &GameScene{
		mouseX: 0,
		mouseY: 0,
	}
}

func init() {
	var id colorm.ColorM

	var mono colorm.ColorM
	mono.ChangeHSV(0, 0, 1)

	for j := 0; j < colorm.Dim-1; j++ {
		for i := 0; i < colorm.Dim-1; i++ {
			lightGray.SetElement(i, j, mono.Element(i, j)*0.7+id.Element(i, j)*0.3)
		}
	}

	lightGray.Translate(0.3, 0.3, 0.3, 0)
}

func (gs *GameScene) drawBackground(r *ebiten.Image) {
	r.Fill(color.Black)

	w, h := imageGameBG.Bounds().Dx(), imageGameBG.Bounds().Dy()
	scaleW := ScreenWidth / float64(w)
	scaleH := ScreenHeight / float64(h)
	scale := scaleW
	if scale < scaleH {
		scale = scaleH
	}

	op := &colorm.DrawImageOptions{}
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(ScreenWidth/2, ScreenHeight/2)
	op.Filter = ebiten.FilterLinear
	colorm.DrawImage(r, imageGameBG, lightGray, op)
}

func (gs *GameScene) Update(state *GameState) error {
	gs.mouseX, gs.mouseY = state.Input.MouseXY()
	
	if state.Input.Debug() {
		gs.inputDebug = true
	} else {
		gs.inputDebug = false
	}

	return nil
}

func (gs *GameScene) Draw(r *ebiten.Image) {
	gs.drawBackground(r)
	r.DrawImage(imageWindows, nil)
	if  gs.inputDebug {
		DrawText(r, "[]", gs.mouseX, gs.mouseY, 1, color.RGBA{0x80, 0, 0, 0xff}, text.AlignCenter, text.AlignStart)
	}
}