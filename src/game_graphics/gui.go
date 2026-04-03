package game_graphics

import (
	"fmt"

	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/Go-20255/team-project-go-getters/src/internal"
	//"board_gen"
)

// original palette
// maybe have mult pallettes in the furture? idk...
var (
	colorHidden   = color.RGBA{180, 180, 180, 255}
	colorRevealed = color.RGBA{220, 220, 220, 255}
	colorFlagged  = color.RGBA{255, 165, 0, 255}
	colorMine     = color.RGBA{255, 0, 0, 255}
	colorBorder   = color.RGBA{90, 90, 90, 255}

	// minesweeper num colors
	mineCountColors = [9]color.RGBA{
		{},                // 0 idk
		{0, 0, 255, 255},    // 1 blue
		{0, 128, 0, 255},    // 2 green
		{255, 0, 0, 255},    // 3 red
		{0, 0, 128, 255},    // 4 dark blue
		{128, 0, 0, 255},    // 5 maroon
		{0, 128, 128, 255},   // 6 teal
		{0, 0, 0, 255},     // 7 black
		{128, 128, 128, 255}, // 8 gray
	}
)


func NewGUI(bg *BoardGen, width, height int) *GUI {
	return &GUI{
		BoardGen: bg,
		Width:    width,
		Height:   height,
	}
}

//mouse input
func (g *GUI) Update() error {
	mx, my := ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		tx, ty := g.BoardGen.PixelToTile(mx, my)
		if g.BoardGen.InBounds(tx, ty) {
			g.BoardGen.PropagateReveal(tx, ty)
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		tx, ty := g.BoardGen.PixelToTile(mx, my)
		if g.BoardGen.InBounds(tx, ty) {
			g.BoardGen.PlaceFlag(tx, ty)
		}
	}

	return nil
}

//renders every tile on the board each frame
func (g *GUI) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{50, 50, 50, 255}) //bg

	ctrl := g.BoardGen.Controller
	for ty := 0; ty < ctrl.Height; ty++ {
		for tx := 0; tx < ctrl.Width; tx++ {
			g.drawTile(screen, tx, ty)
		}
	}

	if ctrl.IsGameOver() {
		ebitenutil.DebugPrint(screen, "GAME OVER, press R to restart")
	} else if ctrl.IsGameWon() {
		ebitenutil.DebugPrint(screen, "YOU WIN!, press R to restart")
	}
}

// Layout returns the logical screen dimensions
func (g *GUI) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.Width, g.Height
}

//helpers funcs

// drawTile renders a single tile at given pos
func (g *GUI) drawTile(screen *ebiten.Image, tx, ty int) {
	tile := g.BoardGen.Controller.GetTile(tx, ty)
	px, py := g.BoardGen.TileToPixel(tx, ty)
	size := float32(g.BoardGen.TileSize)

	//fill color
	var fill color.RGBA
	switch tile.State {
	case internal.TileHidden:
		fill = colorHidden
	case internal.TileFlagged:
		fill = colorFlagged
	case internal.TileRevealed:
		if tile.HasMine {
			fill = colorMine
		} else {
			fill = colorRevealed
		}
	}

	// Fill
	vector.DrawFilledRect(
		screen,
		float32(px)+1, float32(py)+1,
		size-2, size-2,
		fill, false,
	)

	// Border
	vector.StrokeRect(
		screen,
		float32(px), float32(py),
		size, size,
		1, colorBorder, false,
	)

	// Num overlay
	if tile.State == internal.TileRevealed && !tile.HasMine {
		g.drawMineCount(screen, tx, ty)
	}

	//flag symbl
	if tile.State == internal.TileFlagged {
		cx := px + g.BoardGen.TileSize/2
		cy := py + g.BoardGen.TileSize/2
		ebitenutil.DebugPrintAt(screen, "F", cx-3, cy-4)
	}
}

//  renders the adjacent-mine digit inside a revealed tile
func (g *GUI) drawMineCount(screen *ebiten.Image, tx, ty int) {
	tile := g.BoardGen.Controller.GetTile(tx, ty)
	if tile.AdjacentMines == 0 {
		return
	}

	px, py := g.BoardGen.TileToPixel(tx, ty)
	label := fmt.Sprintf("%d", tile.AdjacentMines)

	// Centre the single character kinda within the tile
	offsetX := g.BoardGen.TileSize/2 - 3
	offsetY := g.BoardGen.TileSize/2 - 4

	// ebitenutil.DebugPrintAt doesn't support custom colors, so using temp image tinted to the correct number color, then composite it.
	_ = mineCountColors[tile.AdjacentMines] // apparently this can be used to render custom colored text... Idk if I want to do this yet
	ebitenutil.DebugPrintAt(screen, label, px+offsetX, py+offsetY)
}
