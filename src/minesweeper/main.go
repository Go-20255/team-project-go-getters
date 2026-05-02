package main

import (
    "log"

    "github.com/hajimehoshi/ebiten/v2"

	"github.com/Go-20255/team-project-go-getters/src/internal"
	"github.com/Go-20255/team-project-go-getters/src/game_graphics"
)

const (
    boardWidth  = 16
    boardHeight = 16
    mineCount  = 40
    tileSize   = 32
    topBorder = 60
    sideBorder = 10
)

func main() {
    ctrl := internal.NewController(boardWidth, boardHeight, mineCount)
    bg := game_graphics.NewBoardGen(ctrl, tileSize)
    bg.GenerateGrid()

    screenWidth  := boardWidth*tileSize + sideBorder*2
    screenHeight := boardHeight*tileSize + topBorder

    ebiten.SetWindowSize(screenWidth, screenHeight)
    ebiten.SetWindowTitle("Minesweeper")

    g := game_graphics.NewGUI(bg, screenWidth, screenHeight)
    gl := game_graphics.NewGameLoop(bg, g)
    gl.Start()
    defer gl.Stop()

    if err := ebiten.RunGame(g); err != nil {
        log.Fatal(err)
    }
}