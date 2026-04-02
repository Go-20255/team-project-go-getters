package client


import(

	"image/color"

	
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/colorm"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)


func drawButton(screen *ebiten.Image, b *Button, clr color.RGBA, hover color.RGBA) {
	if b.hovered {
		clr = hover
	}
	vector.FillRect(screen, float32(b.x), float32(b.y), float32(b.w), float32(b.h), clr, true)
	ebitenutil.DebugPrintAt(screen, b.label, b.x+4, b.y+8)

}