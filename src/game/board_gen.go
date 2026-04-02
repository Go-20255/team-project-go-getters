package game

import (
	//"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/colorm"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "github.com/hajimehoshi/ebiten/v2/vector"
)


//LOCAL TYPES





// GLOBAL VARS

const (
	screenWidth  = 600
	screenHeight = 600
)



func init() {
}

func NewGame() *Game{
	g := &Game{
		canvasImage: ebiten.NewImage(screenWidth, screenHeight),
	}
	
	g.canvasImage.Fill(color.White)
	return g
}


func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.canvasImage, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func GenBoard(){
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Drawing App")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}