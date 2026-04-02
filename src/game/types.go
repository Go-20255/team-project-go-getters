package game

import (
	//"fmt"
	//"image"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/colorm"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "github.com/hajimehoshi/ebiten/v2/vector"
)


// USE FOR UNIVERSAL TYPES AND STRUCTS

type pos struct {
	x int
	y int
}

type Game struct {
	cursor  pos
	count   int

	canvasImage *ebiten.Image
}
