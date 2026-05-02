package game_graphics

import (
	"fmt"
	"time"

	"image/color"
	"embed"
    "image"
    _ "image/png"
    "bytes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/Go-20255/team-project-go-getters/src/internal"
	//"board_gen"
)

//for images
//go:embed assets/*.png
var numberAssets embed.FS
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
		TopBorder: 60,
		SideBorder: 10,
		StartTime: time.Now(),
		NumImages: loadNumberImages(),
		FlagImage: loadSprite("assets/flag.png"),
        BombImage: loadSprite("assets/bomb.png"),
	}
}

// for loading bomb and flag
func loadSprite(path string) *ebiten.Image {
    data, err := numberAssets.ReadFile(path)
    if err != nil {
        panic(fmt.Sprintf("failed to read %s: %v", path, err))
    }
    img, _, err := image.Decode(bytes.NewReader(data))
    if err != nil {
        panic(fmt.Sprintf("failed to decode %s: %v", path, err))
    }
    return ebiten.NewImageFromImage(img)
}


//load images for minesweeper in GUI with  a map

func loadNumberImages() [9]*ebiten.Image {
    var imgs [9]*ebiten.Image
    for i := 1; i <= 8; i++ {
        path := fmt.Sprintf("assets/%d.png", i)
        data, err := numberAssets.ReadFile(path)
        if err != nil {
            panic(fmt.Sprintf("failed to read %s: %v", path, err))
        }
        img, _, err := image.Decode(bytes.NewReader(data))
        if err != nil {
            panic(fmt.Sprintf("failed to decode %s: %v", path, err))
        }
        imgs[i] = ebiten.NewImageFromImage(img)
    }
    return imgs
}
//mouse input
func (g *GUI) Update() error {
	mx, my := ebiten.CursorPosition()

	//reset button check
    btnX, btnY := g.Width-60, 10
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if mx >= btnX && mx <= btnX+50 && my >= btnY && my <= btnY+30 {
			g.BoardGen.Reset()
			g.StartTime = time.Now()
			return nil
    	}
	}

	//get rid of border offset
	mx -= g.SideBorder
    my -= g.TopBorder

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
	screen.Fill(colorRevealed) //bg

	//draw border/ timers and such
	g.drawHUD(screen)

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


func (g *GUI) drawHUD(screen *ebiten.Image) {

	//top panel
	vector.FillRect(screen, 0,0,
		float32(g.Width), float32(g.TopBorder),
		colorBorder, false,
	)

	//timer
	elapsed := int(time.Since(g.StartTime).Seconds())
	timerStr := fmt.Sprintf("Time: %03d", elapsed)
	ebitenutil.DebugPrintAt(screen, timerStr, 10, 20)

	// counter of remaining flags
	remaining := g.BoardGen.Controller.MineCount - g.BoardGen.FlagsPlaced()
	flagStr := fmt.Sprintf("Flags Remaining: %d", remaining)
	ebitenutil.DebugPrintAt(screen, flagStr, g.Width/2-50, 20)
	
	btnX, btnY := g.Width-60, 10
    vector.FillRect(screen,
        float32(btnX), float32(btnY), 50, 30,
        color.RGBA{180, 60, 60, 255}, false,
    )
    ebitenutil.DebugPrintAt(screen, "Reset", btnX+8, btnY+10)
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

    px += g.SideBorder
    py += g.TopBorder

    var fill color.RGBA
    switch tile.State {
    case internal.TileHidden:
        fill = colorHidden
    case internal.TileFlagged:
        fill = colorHidden 
        // if tile.HasMine {
        //     fill = colorRevealed 
        // }
    }

    vector.FillRect(screen, float32(px)+1, float32(py)+1, size-2, size-2, fill, false)
    vector.StrokeRect(screen, float32(px), float32(py), size, size, 1, colorBorder, false)

    //sprite overlay helper
    drawSprite := func(img *ebiten.Image) {
        if img == nil {
            return
        }
        imgW, imgH := img.Bounds().Dx(), img.Bounds().Dy()
        op := &ebiten.DrawImageOptions{}
        centerX := float64(px) + float64(g.BoardGen.TileSize-imgW)/2
        centerY := float64(py) + float64(g.BoardGen.TileSize-imgH)/2
        op.GeoM.Translate(centerX, centerY)
        screen.DrawImage(img, op)
    }

    switch tile.State {
    case internal.TileRevealed:
        if tile.HasMine {
            drawSprite(g.BombImage)
        } else {
            g.drawMineCount(screen, tx, ty)
        }
    case internal.TileFlagged:
        drawSprite(g.FlagImage)
    }
}

//  renders the adjacent-mine digit inside a revealed tile
func (g *GUI) drawMineCount(screen *ebiten.Image, tx, ty int) {
	tile := g.BoardGen.Controller.GetTile(tx, ty)
	if tile.AdjacentMines == 0 {
		return
	}

	img := g.NumImages[tile.AdjacentMines]
	//fmt.Printf("drawing %d at tile %d,%d, img nil: %v\n", tile.AdjacentMines, tx, ty, img == nil)
    if img == nil {
        return
    }
	

	px, py := g.BoardGen.TileToPixel(tx, ty)
	px += g.SideBorder
    py += g.TopBorder

	imgW, imgH := img.Bounds().Dx(), img.Bounds().Dy()

    op := &ebiten.DrawImageOptions{}
    // Center the image on the tile
    centerX := float64(px) + float64(g.BoardGen.TileSize-imgW)/2
    centerY := float64(py) + float64(g.BoardGen.TileSize-imgH)/2
    op.GeoM.Translate(centerX, centerY)

	screen.DrawImage(img, op)
}
