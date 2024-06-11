package goraphics

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var imageBackground *ebiten.Image

func init() {
	img, _, err := ebitenutil.NewImageFromFile("./assets/solids/blue.png")
	if err != nil {
		panic(err)
	}
	imageBackground = ebiten.NewImageFromImage(img)
}

type TitleScene struct {
	count int
}

func (ts *TitleScene) Update(state *GameState) error {
	ts.count++
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.SceneManager.GoTo(NewGameScene())
	}
	return nil
}

func (ts *TitleScene) Draw(r *ebiten.Image) {
	ts.drawTitleBackground(r, ts.count)
	
	message := "PRESS SPACE TO START"
	x := ScreenWidth / 2
	y := ScreenHeight - 48
	DrawTextWithShadow(r, message, x, y, 1, color.RGBA{0x80, 0, 0, 0xff}, text.AlignCenter, text.AlignStart)
}


func (ts *TitleScene) drawTitleBackground(r *ebiten.Image, c int) {
	w, h := imageBackground.Bounds().Dx(), imageBackground.Bounds().Dy()
	op := &ebiten.DrawImageOptions{}
	for i := 0; i < (ScreenWidth/w+1)*(ScreenHeight/h+2); i++ {
		op.GeoM.Reset()
		dx := -(c / 4) % w
		dy := (c / 4) % h
		dstX := (i % (ScreenWidth/w + 1)) * w + dx
		dstY := (i / (ScreenWidth/w + 1) - 1) * h + dy
		op.GeoM.Translate(float64(dstX), float64(dstY))
		r.DrawImage(imageBackground, op)
	}
}